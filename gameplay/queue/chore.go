// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"context"
	"net/http"

	"github.com/BoostyLabs/thelooper"
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
	ChoreError = errs.Class("queue chore error")
)

// Chore checks user which don't play and initiates game for them.
//
// architecture: Chore
type Chore struct {
	config Config
	log    logger.Logger

	service *Service
	Loop    *thelooper.Loop

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

// Run starts chore that checks user which don't play and initiates game for them.
func (chore *Chore) Run(ctx context.Context) (err error) {
	firstRequestChan := make(chan Request)
	secondRequestChan := make(chan Request)

	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		clients := chore.service.ListNotPlayingUsers()

		if len(clients) >= 2 {
			for k := range clients {
				isEvenNumber := k%2 != 1
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

						if err = chore.service.Play(ctx, firstClient, secondClient); err != nil {
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

// Close closes the chore for re-check the expiration time of the token.
func (chore *Chore) Close() {
	chore.Loop.Close()
}
