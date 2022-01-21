// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"context"
	"errors"
	"math/big"
	"net/http"

	"github.com/BoostyLabs/evmsignature"
	"github.com/BoostyLabs/thelooper"
	"github.com/gorilla/websocket"
	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/internal/logger"
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
	Loop             *thelooper.Loop
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
		Loop:             thelooper.NewLoop(config.PlaceRenewalInterval),
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
	defer func() {
		close(firstRequestChan)
		close(secondRequestChan)
	}()

	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		notPlayingUsers := chore.service.ListNotPlayingUsers()
		notPlayingUsers = isLenOdd(notPlayingUsers)

		if len(notPlayingUsers) >= 2 {
			pairsOfClients := divideClients(notPlayingUsers)
			for _, pair := range pairsOfClients {
				go func(pair []Client) {
					firstClient := pair[0]
					secondClient := pair[1]
					if err = chore.service.UpdateIsPlaying(firstClient.UserID, true); err != nil {
						chore.log.Error("could not update is play", ChoreError.Wrap(err))
						return
					}
					if err = chore.service.UpdateIsPlaying(secondClient.UserID, true); err != nil {
						chore.log.Error("could not update is play", ChoreError.Wrap(err))
						return
					}
					if err := firstClient.WriteJSON(http.StatusOK, "do you confirm play?"); err != nil {
						chore.log.Error("could not write json", ChoreError.Wrap(err))
						return
					}
					if err := secondClient.WriteJSON(http.StatusOK, "do you confirm play?"); err != nil {
						chore.log.Error("could not write json", ChoreError.Wrap(err))
						return
					}

					go func() {
						request, err := firstClient.ReadJSON()

						if errors.Is(err, websocket.ErrCloseSent) {
							if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
								return
							}
							if err := secondClient.WriteJSON(http.StatusOK, "you are still in search"); err != nil {
								chore.log.Error("could not write json", ChoreError.Wrap(err))
								return
							}
						}
						firstRequestChan <- request
					}()

					go func() {
						request, err := secondClient.ReadJSON()

						if errors.Is(err, websocket.ErrCloseSent) {
							if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
								return
							}
							if err := firstClient.WriteJSON(http.StatusOK, "you are still in search"); err != nil {
								chore.log.Error("could not write json", ChoreError.Wrap(err))
								return
							}
						}
						secondRequestChan <- request
					}()

					var firstRequest, secondRequest Request
					for {
						notPlayingUsers = chore.service.ListNotPlayingUsers()
						if isClientInSlice(firstClient, notPlayingUsers) {

							if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
								return
							}

							if err := secondClient.WriteJSON(http.StatusOK, "you are still in search"); err != nil {
								chore.log.Error("could not write json", ChoreError.Wrap(err))
								return
							}
							return
						}

						if isClientInSlice(secondClient, notPlayingUsers) {
							if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
								return
							}

							if err := firstClient.WriteJSON(http.StatusOK, "you are still in search"); err != nil {
								chore.log.Error("could not write json", ChoreError.Wrap(err))
								return
							}
							return
						}

						select {
						case firstRequest = <-firstRequestChan:
							notPlayingUsers = chore.service.ListNotPlayingUsers()
							if isClientInSlice(firstClient, notPlayingUsers) || isClientInSlice(secondClient, notPlayingUsers) {
								return
							}
							if (firstRequest == Request{}) {
								continue
							}
							if !firstRequest.Action.isValid() {
								if err := firstClient.WriteJSON(http.StatusBadRequest, "wrong action"); err != nil {
									chore.log.Error("could not write json", ChoreError.Wrap(err))
									return
								}
								if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
									chore.log.Error("could not update is play", ChoreError.Wrap(err))
									return
								}
								if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
									chore.log.Error("could not update is play", ChoreError.Wrap(err))
									return
								}
								return
							}
						case secondRequest = <-secondRequestChan:
							notPlayingUsers = chore.service.ListNotPlayingUsers()
							if isClientInSlice(firstClient, notPlayingUsers) || isClientInSlice(secondClient, notPlayingUsers) {
								return
							}
							if (secondRequest == Request{}) {
								continue
							}

							if !secondRequest.Action.isValid() {
								if err := secondClient.WriteJSON(http.StatusBadRequest, "wrong action"); err != nil {
									chore.log.Error("could not write json", ChoreError.Wrap(err))
									return
								}
								if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
									chore.log.Error("could not update is play", ChoreError.Wrap(err))
									return
								}
								if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
									chore.log.Error("could not update is play", ChoreError.Wrap(err))
									return
								}
								return
							}
						}

						notPlayingUsers = chore.service.ListNotPlayingUsers()
						if isClientInSlice(firstClient, notPlayingUsers) {

							if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
								return
							}

							if err := secondClient.WriteJSON(http.StatusOK, "you are still in search"); err != nil {
								chore.log.Error("could not write json", ChoreError.Wrap(err))
								return
							}
							return
						}

						if isClientInSlice(secondClient, notPlayingUsers) {
							if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
								return
							}

							if err := firstClient.WriteJSON(http.StatusOK, "you are still in search"); err != nil {
								chore.log.Error("could not write json", ChoreError.Wrap(err))
								return
							}
							return
						}

						if firstRequest.Action == ActionReject || secondRequest.Action == ActionReject {
							if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
								return
							}
							if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
								return
							}
							if err := firstClient.WriteJSON(http.StatusOK, "you are still in search"); err != nil {
								chore.log.Error("could not write json", ChoreError.Wrap(err))
								return
							}
							if err := secondClient.WriteJSON(http.StatusOK, "you are still in search"); err != nil {
								chore.log.Error("could not write json", ChoreError.Wrap(err))
								return
							}
							return
						}

						notPlayingUsers = chore.service.ListNotPlayingUsers()
						if isClientInSlice(firstClient, notPlayingUsers) || isClientInSlice(secondClient, notPlayingUsers) {
							return
						}

						if (firstRequest == Request{} || secondRequest == Request{}) {
							continue
						}

						if firstRequest.Action == ActionConfirm && secondRequest.Action == ActionConfirm {
							if err = chore.Play(ctx, firstClient, secondClient); err != nil {
								if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
									chore.log.Error("could not update is play", ChoreError.Wrap(err))
									return
								}
								if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
									chore.log.Error("could not update is play", ChoreError.Wrap(err))
									return
								}
								if err = chore.service.Finish(firstClient.UserID); err != nil {
									chore.log.Error("could not delete client from queue", ChoreError.Wrap(err))
									return
								}
								if err = chore.service.Finish(secondClient.UserID); err != nil {
									chore.log.Error("could not delete client from queue", ChoreError.Wrap(err))
									return
								}
								chore.log.Error("could not play game", ChoreError.Wrap(err))
								return
							}
						}
						return
					}
				}(pair)
			}
		}
		return ChoreError.Wrap(err)
	})
}

func isClientInSlice(element Client, clients []Client) bool {
	for _, client := range clients {
		if client.Connection == client.Connection && client.IsPlaying == client.IsPlaying &&
			client.UserID == client.UserID && client.SquadID == client.SquadID {
			return true
		}
	}
	return false
}

// isLenOdd checks is length of not playing users odd.
// if odd - delete last client from slice.
func isLenOdd(notPlayingUsers []Client) []Client {
	isOddNumber := len(notPlayingUsers)%2 == 1
	if isOddNumber && len(notPlayingUsers) >= 2 {
		notPlayingUsers = notPlayingUsers[:len(notPlayingUsers)-1]
	}
	return notPlayingUsers
}

// DivideClients divides all clients into couples.
func divideClients(clients []Client) [][]Client {
	var dividedClients [][]Client
	for i := 0; i < len(clients); i += 2 {
		element := make([]Client, 2, 2)
		element[0] = clients[i]
		element[1] = clients[i+1]
		dividedClients = append(dividedClients, element)
	}
	return dividedClients
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

	winResult.GameResult.Question = "do you allow us to take your address?"
	winResult.GameResult.Transaction.Value = evmsignature.WeiToEthereum(winResult.Value).String()
	winResult.GameResult.Transaction.UDTContract.Address = chore.config.UDTContract.Address
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

		if winResult.GameResult.Transaction, err = chore.currencywaitlist.Create(ctx, user.ID, *winResult.Value, request.Nonce); err != nil {
			chore.log.Error("could not create item of currencywaitlist", ChoreError.Wrap(err))
			return
		}
	}
	chore.Finish(winResult.Client, winResult.GameResult)
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

// Close closes the chore for re-check the expiration time of the token.
func (chore *Chore) Close() {
	chore.Loop.Close()
}
