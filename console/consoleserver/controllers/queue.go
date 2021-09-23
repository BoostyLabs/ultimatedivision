// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/auth"
	"ultimatedivision/internal/logger"
	"ultimatedivision/queue"
	"ultimatedivision/queue/queuehub"
)

var (
	// ErrQueue is an internal error type for queue controller.
	ErrQueue = errs.Class("queue controller error")
)

// Queue is a mvc controller that handles all queue related views.
type Queue struct {
	log   logger.Logger
	queue *queue.Service
	hub   *queuehub.Hub
}

// NewQueue is a constructor for queue controller.
func NewQueue(log logger.Logger, queue *queue.Service, hub *queuehub.Hub) *Queue {
	queueController := &Queue{
		log:   log,
		queue: queue,
		hub:   hub,
	}

	return queueController
}

// Create is an endpoint that creates place.
func (controller *Queue) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		controller.log.Error("could not connect to websocket", ErrQueue.Wrap(err))
		return
	}

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.log.Error("could not unauthorized user", ErrQueue.Wrap(err))
		return
	}

	client := queuehub.Client{
		UserID: claims.UserID,
		Conn:   conn,
	}
	isSearch, err := controller.hub.ReadSearch(client)
	if err != nil {
		controller.log.Error("could not read searches for client", ErrQueue.Wrap(err))
		return
	}

	if !isSearch {
		if err = controller.queue.Finish(ctx, claims.UserID); err != nil {
			controller.log.Error("could not create place", ErrQueue.Wrap(err))
			return
		}
	} else {
		place := queue.Place{
			UserID: claims.UserID,
			Status: queue.StatusSearches,
		}
		if err = controller.queue.Create(ctx, place); err != nil {
			controller.log.Error("could not create place", ErrQueue.Wrap(err))
			return
		}
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Queue) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrQueue.Wrap(err))
	}
}
