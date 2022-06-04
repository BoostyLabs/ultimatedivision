// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"math/big"
	"net/http"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/zeebo/errs"

	"ultimatedivision/gameplay/gameengine"
	"ultimatedivision/gameplay/matches"
)

var (
	// ErrNoClient indicated that client does not exist.
	ErrNoClient = errs.Class("client does not exist")
	// ErrRead indicates a read error.
	ErrRead = errs.Class("error read from websocket")
	// ErrWrite indicates a write error.
	ErrWrite = errs.Class("error write to websocket")
)

// DB is exposing access to queue database.
//
// architecture: DB
type DB interface {
	// Create adds client to queue in the database.
	Create(client Client)
	// Get returns client from the database.
	Get(userID uuid.UUID) (Client, error)
	// List returns clients from the database.
	List() []Client
	// ListNotPlayingUsers returns clients who don't play game from database.
	ListNotPlayingUsers() []Client
	// UpdateIsPlaying updates is playing status of client in database.
	UpdateIsPlaying(userID uuid.UUID, IsPlaying bool) error
	// Delete deletes client record in database.
	Delete(userID uuid.UUID) error
}

// Client entity describes the value of connect with the client.
type Client struct {
	UserID     uuid.UUID
	Connection *websocket.Conn
	SquadID    uuid.UUID
	IsPlaying  bool
}

// Request entity describes values sent by client.
type Request struct {
	Action        Action               `json:"action"`
	SquadID       uuid.UUID            `json:"squadId"`
	WalletAddress evmsignature.Address `json:"walletAddress"`
	Nonce         int64                `json:"nonce"`
}

// Action defines list of possible clients action.
type Action string

const (
	// ActionStartSearch indicates that the client starts the search.
	ActionStartSearch Action = "startSearch"
	// ActionFinishSearch indicates that the client finishes the search.
	ActionFinishSearch Action = "finishSearch"
	// ActionConfirm indicates that the client confirms the game.
	ActionConfirm Action = "confirm"
	// ActionReject indicates that the client rejects the game.
	ActionReject Action = "reject"
	// ActionAllowAddress indicates that the client allows to add address of wallet.
	ActionAllowAddress Action = "allowAddress"
	// ActionForbidAddress indicates that the client is forbidden to add wallet address.
	ActionForbidAddress Action = "forbidAddress"
)

// Response entity describes values sent to user.
type Response struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

// Config defines configuration for queue.
type Config struct {
	PlaceRenewalInterval       time.Duration                `json:"placeRenewalInterval"`
	WinValue                   string                       `json:"winValue"`
	DrawValue                  string                       `json:"drawValue"`
	UDTContract                evmsignature.Contract        `json:"udtContract"`
	GameConfig                 gameengine.GameConfig        `json:"gameConfig"`
	CoordinateConfig           gameengine.CoordinatesConfig `json:"coordinateConfig"`
	MatchActionRenewalInterval time.Duration                `json:"matchActionRenewalInterval"`
	RoundDuration              time.Duration                `json:"roundDuration"`
	NumberOfRounds             int                          `json:"numberOfRounds"`
}

// ReadJSON reads request sent by client.
func (client *Client) ReadJSON() (Request, error) {
	var request Request
	if err := client.Connection.ReadJSON(&request); err != nil {
		if err = client.WriteJSON(http.StatusBadRequest, err.Error()); err != nil {
			return request, ErrWrite.Wrap(ErrQueue.Wrap(err))
		}
		return request, ErrRead.Wrap(ErrQueue.Wrap(err))
	}
	return request, nil
}

// TODO: move to api layer.

// ReadActionJSON reads action request sent by client.
func (client *Client) ReadActionJSON() ([]gameengine.MakeAction, error) {
	var request []gameengine.MakeAction

	err := client.Connection.SetReadDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		return []gameengine.MakeAction{}, ErrRead.Wrap(ErrQueue.Wrap(err))
	}

	if err = client.Connection.ReadJSON(&request); err != nil {
		if !websocket.IsCloseError(err) || !websocket.IsUnexpectedCloseError(err) {
			if err = client.WriteJSON(http.StatusBadRequest, err.Error()); err != nil {
				return request, ErrRead.Wrap(ErrQueue.Wrap(err))
			}
		}
	}

	return request, nil
}

// WriteJSON writes response to client.
func (client *Client) WriteJSON(status int, message interface{}) error {
	return client.Connection.WriteJSON(Response{
		Status:  status,
		Message: message,
	})
}

// WinResult entity describes values which send to user after win game.
type WinResult struct {
	Client     Client             `json:"client"`
	GameResult matches.GameResult `json:"gameResult"`
	Value      *big.Int           `json:"value"`
}
