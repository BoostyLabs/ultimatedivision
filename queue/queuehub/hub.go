// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queuehub

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Hub describes hub entity.
type Hub struct {
	Clients map[uuid.UUID]*websocket.Conn
}

// Client describes client entity.
type Client struct {
	UserID uuid.UUID
	Conn   *websocket.Conn
	IsPlay bool
}

// Request entity describes values sent by user.
type Request struct {
	Action Action `json:"action"`
	Value  bool   `json:"value"`
}

// Action defines list of possible clients action.
type Action string

const (
	// ActionSearch indicates that user is ready to search for the game.
	ActionSearch Action = "search"
	// ActionPlay indicates that user is ready to play game.
	ActionPlay Action = "play"
)

// Message entity describes values sent to user.
type Message struct {
	Status int    `json:"status"`
	Text   string `json:"text"`
}
