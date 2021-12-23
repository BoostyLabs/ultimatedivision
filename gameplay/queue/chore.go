// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"context"
	"math/big"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/cryptoutils"
	sync2 "ultimatedivision/pkg/sync"
	"ultimatedivision/seasons"
	"ultimatedivision/udts/currencywaitlist"
	"ultimatedivision/users"
)

var (
	// ChoreError represents place chore error type.
	ChoreError = errs.Class("expiration place chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore
type Chore struct {
	config           Config
	log              logger.Logger
	service          *Service
	Loop             *sync2.Cycle
	matches          *matches.Service
	seasons          *seasons.Service
	clubs            *clubs.Service
	currencywaitlist *currencywaitlist.Service
	users            *users.Service
}

// NewChore instantiates Chore.
func NewChore(config Config, log logger.Logger, service *Service, matches *matches.Service, seasons *seasons.Service, clubs *clubs.Service, currencywaitlist *currencywaitlist.Service, users *users.Service) *Chore {
	return &Chore{
		config:           config,
		log:              log,
		service:          service,
		Loop:             sync2.NewCycle(config.PlaceRenewalInterval),
		matches:          matches,
		seasons:          seasons,
		clubs:            clubs,
		currencywaitlist: currencywaitlist,
		users:            users,
	}
}

// Run starts the chore for re-check the expiration time of the token.
func (chore *Chore) Run(ctx context.Context) (err error) {
	firstRequestChan := make(chan Request)
	secondRequestChan := make(chan Request)

	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		clients := chore.service.ListNotPlayingUsers()

		if len(clients) >= 2 {
			for k := range clients {
				isEvenNumber := (k%2 != 1)
				if isEvenNumber {
					continue
				}

				go func(clients []Client, k int) {
					firstClient := clients[k-1]
					secondClient := clients[k]

					if err = chore.service.UpdateIsPlaying(firstClient.UserID, true); err != nil {
						chore.log.Error("could not update is play", ChoreError.Wrap(err))
					}
					if err = chore.service.UpdateIsPlaying(secondClient.UserID, true); err != nil {
						chore.log.Error("could not update is play", ChoreError.Wrap(err))
					}

					if err := firstClient.WriteJSON(http.StatusOK, "you confirm play?"); err != nil {
						chore.log.Error("could not write json", ChoreError.Wrap(err))
					}
					if err := secondClient.WriteJSON(http.StatusOK, "you confirm play?"); err != nil {
						chore.log.Error("could not write json", ChoreError.Wrap(err))
					}

					go func() {
						request, err := firstClient.ReadJSON()
						if err != nil {
							chore.log.Error("could not read json", ChoreError.Wrap(err))
						}
						firstRequestChan <- request
					}()

					go func() {
						request, err := secondClient.ReadJSON()
						if err != nil {
							chore.log.Error("could not read json", ChoreError.Wrap(err))
						}
						secondRequestChan <- request
					}()

					var firstRequest, secondRequest Request
					for {
						select {
						case firstRequest = <-firstRequestChan:
							if (firstRequest != Request{}) {
								if firstRequest.Action != ActionConfirm && firstRequest.Action != ActionReject {
									if err := firstClient.WriteJSON(http.StatusBadRequest, "wrong action"); err != nil {
										chore.log.Error("could not write json", ChoreError.Wrap(err))
									}

									if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
										chore.log.Error("could not update is play", ChoreError.Wrap(err))
									}
									if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
										chore.log.Error("could not update is play", ChoreError.Wrap(err))
									}
									return
								}
							}
						case secondRequest = <-secondRequestChan:
							if (secondRequest != Request{}) {
								if secondRequest.Action != ActionConfirm && secondRequest.Action != ActionReject {
									if err := secondClient.WriteJSON(http.StatusBadRequest, "wrong action"); err != nil {
										chore.log.Error("could not write json", ChoreError.Wrap(err))
									}

									if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
										chore.log.Error("could not update is play", ChoreError.Wrap(err))
									}
									if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
										chore.log.Error("could not update is play", ChoreError.Wrap(err))
									}
									return
								}
							}
						}

						if (firstRequest == Request{} && secondRequest == Request{}) {
							continue
						}

						if firstRequest.Action == ActionReject || secondRequest.Action == ActionReject {
							if err := firstClient.WriteJSON(http.StatusOK, "you are still in search!"); err != nil {
								chore.log.Error("could not write json", ChoreError.Wrap(err))
							}
							if err := secondClient.WriteJSON(http.StatusOK, "you are still in search!"); err != nil {
								chore.log.Error("could not write json", ChoreError.Wrap(err))
							}

							if err = chore.service.Finish(firstClient.UserID); err != nil {
								chore.log.Error("could not delete client from queue", ChoreError.Wrap(err))
							}
							if err = chore.service.Finish(secondClient.UserID); err != nil {
								chore.log.Error("could not delete client from queue", ChoreError.Wrap(err))
							}
							return
						}

						if (firstRequest == Request{} || secondRequest == Request{}) {
							continue
						}

						if err = chore.Play(ctx, firstClient, secondClient); err != nil {
							if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
							}
							if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
							}
							chore.log.Error("could not play game", ChoreError.Wrap(err))
						}
						return
					}
				}(clients, k)
			}
		}
		return ChoreError.Wrap(err)
	})
}

// Play method contains all the logic for playing matches.
func (chore *Chore) Play(ctx context.Context, firstClient, secondClient Client) error {
	squadCardsFirstClient, err := chore.service.clubs.ListSquadCards(ctx, firstClient.SquadID)
	if err != nil {
		return ChoreError.Wrap(err)
	}
	if len(squadCardsFirstClient) != clubs.SquadSize {
		if err := firstClient.WriteJSON(http.StatusInternalServerError, "squad is not full"); err != nil {
			return ChoreError.Wrap(err)
		}
	}

	squadCardsSecondClient, err := chore.service.clubs.ListSquadCards(ctx, secondClient.SquadID)
	if err != nil {
		return ChoreError.Wrap(err)
	}
	if len(squadCardsSecondClient) != clubs.SquadSize {
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "squad is not full"); err != nil {
			return ChoreError.Wrap(err)
		}
	}

	firstClientSquad, err := chore.clubs.GetSquad(ctx, firstClient.SquadID)
	if err != nil {
		return ChoreError.Wrap(err)
	}

	firstClientClub, err := chore.clubs.Get(ctx, firstClientSquad.ClubID)
	if err != nil {
		return ChoreError.Wrap(err)
	}

	season, err := chore.seasons.GetSeasonByDivisionID(ctx, firstClientClub.DivisionID)
	if err != nil {
		if err := firstClient.WriteJSON(http.StatusInternalServerError, "could not season id"); err != nil {
			return ChoreError.Wrap(err)
		}
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not season id"); err != nil {
			return ChoreError.Wrap(err)
		}
		return ChoreError.Wrap(err)
	}

	matchesID, err := chore.matches.Create(ctx, firstClient.SquadID, secondClient.SquadID, firstClient.UserID, secondClient.UserID, season.ID)
	if err != nil {
		if err := firstClient.WriteJSON(http.StatusInternalServerError, "match error"); err != nil {
			return ChoreError.Wrap(err)
		}
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "match error"); err != nil {
			return ChoreError.Wrap(err)
		}
		return ChoreError.Wrap(err)
	}

	gameResult, err := chore.matches.GetGameResult(ctx, matchesID)
	if err != nil {
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not get result of match"); err != nil {
			return ChoreError.Wrap(err)
		}
		return ChoreError.Wrap(err)
	}

	var firstClientResult matches.GameResult
	var secondClientResult matches.GameResult

	firstClientResult.MatchResults = make([]matches.MatchResult, len(gameResult.MatchResults))
	_ = copy(firstClientResult.MatchResults, gameResult.MatchResults)
	secondClientResult.MatchResults = make([]matches.MatchResult, len(gameResult.MatchResults))
	_ = copy(secondClientResult.MatchResults, gameResult.MatchResults)

	switch {
	case firstClient.UserID == gameResult.MatchResults[0].UserID:
		secondClientResult.MatchResults = matches.Swap(gameResult.MatchResults)
	case secondClient.UserID == gameResult.MatchResults[0].UserID:
		firstClientResult.MatchResults = matches.Swap(gameResult.MatchResults)
	}

	var value = new(big.Int)
	value.SetString(chore.config.WinValue, 10)

	switch {
	case firstClientResult.MatchResults[0].QuantityGoals > secondClientResult.MatchResults[0].QuantityGoals:
		winResult := WinResult{
			Client:     firstClient,
			GameResult: firstClientResult,
			Value:      value,
		}

		go chore.FinishWithWinResult(ctx, winResult)
		go chore.Finish(secondClient, secondClientResult)
	case firstClientResult.MatchResults[0].QuantityGoals < secondClientResult.MatchResults[0].QuantityGoals:
		winResult := WinResult{
			Client:     secondClient,
			GameResult: secondClientResult,
			Value:      value,
		}

		go chore.FinishWithWinResult(ctx, winResult)
		go chore.Finish(firstClient, firstClientResult)
	default:
		var value = new(big.Int)
		value.SetString(chore.config.DrawValue, 10)

		winResult := WinResult{
			Client:     firstClient,
			GameResult: firstClientResult,
			Value:      value,
		}
		go chore.FinishWithWinResult(ctx, winResult)

		winResult = WinResult{
			Client:     secondClient,
			GameResult: secondClientResult,
			Value:      value,
		}
		go chore.FinishWithWinResult(ctx, winResult)
	}

	return nil
}

// FinishWithWinResult sends win result and finishes the connection.
func (chore *Chore) FinishWithWinResult(ctx context.Context, winResult WinResult) {
	user, err := chore.users.Get(ctx, winResult.Client.UserID)
	if err != nil {
		chore.log.Error("could not get user", ChoreError.Wrap(err))
		return
	}

	if user.Wallet != "" {
		if winResult.GameResult.Transaction, err = chore.currencywaitlist.Create(ctx, user.ID, *winResult.Value); err != nil {
			chore.log.Error("could not create item of currencywaitlist", ChoreError.Wrap(err))
			return
		}
	} else {
		winResult.GameResult.Question = "you allow us to take your address?"
		winResult.GameResult.Transaction.Value = cryptoutils.WeiToEthereum(winResult.Value).String()
		if err := winResult.Client.WriteJSON(http.StatusOK, winResult.GameResult); err != nil {
			chore.log.Error("could not write json", ChoreError.Wrap(err))
			return
		}

		request, err := winResult.Client.ReadJSON()
		if err != nil {
			chore.log.Error("could not read json", ChoreError.Wrap(err))
			return
		}

		if request.Action != ActionForbidAddress && request.Action != ActionAllowAddress {
			if err := winResult.Client.WriteJSON(http.StatusBadRequest, "wrong action"); err != nil {
				chore.log.Error("could not write json", ChoreError.Wrap(err))
				return
			}
		}

		if request.Action == ActionAllowAddress {
			if !request.WalletAddress.IsValidAddress() {
				if err := winResult.Client.WriteJSON(http.StatusBadRequest, "invalid address of user's wallet"); err != nil {
					chore.log.Error("could not write json", ChoreError.Wrap(err))
					return
				}
			}

			if err = chore.users.UpdateWalletAddress(ctx, request.WalletAddress, winResult.Client.UserID); err != nil {
				if !users.ErrWalletAddressAlreadyInUse.Has(err) {
					chore.log.Error("could not update user's wallet address", ChoreError.Wrap(err))
					return
				}
			}

			if winResult.GameResult.Transaction, err = chore.currencywaitlist.Create(ctx, user.ID, *winResult.Value); err != nil {
				chore.log.Error("could not create item of currencywaitlist", ChoreError.Wrap(err))
				return
			}
		}

		chore.Finish(winResult.Client, winResult.GameResult)
	}
}

// Finish sends result and finishes the connection.
func (chore *Chore) Finish(client Client, gameResult matches.GameResult) {
	var err error

	if err = client.WriteJSON(http.StatusOK, gameResult); err != nil {
		chore.log.Error("could not write json", ChoreError.Wrap(err))
		return
	}

	if err = chore.service.Finish(client.UserID); err != nil {
		chore.log.Error("could not finish match", ChoreError.Wrap(err))
		return
	}
	defer func() {
		if err = client.Connection.Close(); err != nil {
			chore.log.Error("could not close websocket", ChoreError.Wrap(err))
		}
	}()
}

// Close closes the chore chore for re-check the expiration time of the token.
func (chore *Chore) Close() {
	chore.Loop.Close()
}
