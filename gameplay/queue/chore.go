// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"context"
	"net/http"
	"ultimatedivision/cards"

	"github.com/BoostyLabs/evmsignature"
	"github.com/BoostyLabs/thelooper"
	"github.com/ethereum/go-ethereum/common"
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
	cards            *cards.Service
	currencywaitlist *currencywaitlist.Service
	users            *users.Service
}

// NewChore instantiates Chore.
func NewChore(config Config, log logger.Logger, service *Service, matches *matches.Service, seasons *seasons.Service, clubs *clubs.Service, cards *cards.Service, currencywaitlist *currencywaitlist.Service, users *users.Service) *Chore {
	return &Chore{
		config:           config,
		log:              log,
		service:          service,
		Loop:             thelooper.NewLoop(config.PlaceRenewalInterval),
		matches:          matches,
		seasons:          seasons,
		clubs:            clubs,
		cards:            cards,
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

						response, err := chore.Play(ctx, firstClient, secondClient)
						if err != nil {
							if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
							}
							if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
							}
							chore.log.Error("could not play game", ChoreError.Wrap(err))
						}

						if err := firstClient.WriteJSON(http.StatusOK, response); err != nil {
							chore.log.Error("could not write json", ChoreError.Wrap(err))
						}
						if err := secondClient.WriteJSON(http.StatusOK, response); err != nil {
							chore.log.Error("could not write json", ChoreError.Wrap(err))
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
func (chore *Chore) Play(ctx context.Context, firstClient, secondClient Client) (GameplayResponse, error) {
	squadCardsFirstClient, err := chore.service.clubs.ListSquadCards(ctx, firstClient.SquadID)
	if err != nil {
		return GameplayResponse{}, ChoreError.Wrap(err)
	}
	if len(squadCardsFirstClient) != clubs.SquadSize {
		if err := firstClient.WriteJSON(http.StatusInternalServerError, "squad is not full"); err != nil {
			return GameplayResponse{}, ChoreError.Wrap(err)
		}
	}

	squadCardsSecondClient, err := chore.service.clubs.ListSquadCards(ctx, secondClient.SquadID)
	if err != nil {
		return GameplayResponse{}, ChoreError.Wrap(err)
	}
	if len(squadCardsSecondClient) != clubs.SquadSize {
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "squad is not full"); err != nil {
			return GameplayResponse{}, ChoreError.Wrap(err)
		}
	}

	firstClientSquad, err := chore.clubs.GetSquad(ctx, firstClient.SquadID)
	if err != nil {
		return GameplayResponse{}, ChoreError.Wrap(err)
	}

	firstClientClub, err := chore.clubs.Get(ctx, firstClientSquad.ClubID)
	if err != nil {
		return GameplayResponse{}, ChoreError.Wrap(err)
	}

	secondClientSquad, err := chore.clubs.GetSquad(ctx, firstClient.SquadID)
	if err != nil {
		return GameplayResponse{}, ChoreError.Wrap(err)
	}

	secondClientClub, err := chore.clubs.Get(ctx, firstClientSquad.ClubID)
	if err != nil {
		return GameplayResponse{}, ChoreError.Wrap(err)
	}

	season, err := chore.seasons.GetSeasonByDivisionID(ctx, firstClientClub.DivisionID)
	if err != nil {
		if err := firstClient.WriteJSON(http.StatusInternalServerError, "could not season id"); err != nil {
			return GameplayResponse{}, ChoreError.Wrap(err)
		}
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not season id"); err != nil {
			return GameplayResponse{}, ChoreError.Wrap(err)
		}
		return GameplayResponse{}, ChoreError.Wrap(err)
	}

	matchID, err := chore.matches.Create(ctx, firstClient.SquadID, secondClient.SquadID, firstClient.UserID, secondClient.UserID, season.ID)
	if err != nil {
		if err := firstClient.WriteJSON(http.StatusInternalServerError, "match error"); err != nil {
			return GameplayResponse{}, ChoreError.Wrap(err)
		}
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "match error"); err != nil {
			return GameplayResponse{}, ChoreError.Wrap(err)
		}
		return GameplayResponse{}, ChoreError.Wrap(err)
	}

	var response GameplayResponse

	response.MatchID = matchID
	response.FirstClub.Club = firstClientClub
	response.FirstClub.Squad = firstClientSquad

	firstClientSquadCards, err := chore.clubs.ListSquadCards(ctx, firstClientSquad.ID)
	if err != nil {
		return GameplayResponse{}, ChoreError.Wrap(err)
	}

	for i := 0; i < len(firstClientSquadCards); i++ {
		firstClientSquadCards[i].Card, err = chore.cards.Get(ctx, firstClientSquadCards[i].Card.ID)
		if err != nil {
			return GameplayResponse{}, ChoreError.Wrap(err)
		}
	}
	response.FirstClub.SquadCards = firstClientSquadCards

	response.SecondClub.Club = secondClientClub
	response.SecondClub.Squad = secondClientSquad

	secondClientSquadCards, err := chore.clubs.ListSquadCards(ctx, firstClientSquad.ID)
	if err != nil {
		return GameplayResponse{}, ChoreError.Wrap(err)
	}

	for i := 0; i < len(secondClientSquadCards); i++ {
		secondClientSquadCards[i].Card, err = chore.cards.Get(ctx, secondClientSquadCards[i].Card.ID)
		if err != nil {
			return GameplayResponse{}, ChoreError.Wrap(err)
		}
	}
	response.FirstClub.SquadCards = secondClientSquadCards

	return response, nil
}

// FinishWithWinResult sends win result and finishes the connection.
func (chore *Chore) FinishWithWinResult(ctx context.Context, winResult WinResult) {
	user, err := chore.users.Get(ctx, winResult.Client.UserID)
	if err != nil {
		chore.log.Error("could not get user", ChoreError.Wrap(err))
		return
	}

	winResult.GameResult.Question = "do you allow us to take your address?"
	winResult.GameResult.Transaction.Value = evmsignature.WeiBigToEthereumBig(winResult.Value).String()
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
		if err := request.WalletAddress.IsValidAddress(); err != nil {
			if err := winResult.Client.WriteJSON(http.StatusBadRequest, "invalid address of user's wallet"); err != nil {
				chore.log.Error("could not write json", ChoreError.Wrap(err))
				return
			}
		}

		if err = chore.users.UpdateWalletAddress(ctx, common.HexToAddress(string(request.WalletAddress)), winResult.Client.UserID, users.WalletTypeETH); err != nil {
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
