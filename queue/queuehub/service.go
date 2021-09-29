// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queuehub

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/zeebo/errs"
)

// ErrHub indicated that there was an error in service.
var ErrHub = errs.Class("hub service error")

// ErrWriteHub indicated that there was an error sending response.
var ErrWriteHub = errs.Class("write response in hub error")

// NewHub is a constructor for hub entity.
func NewHub() *Hub {
	return &Hub{
		Clients: make(map[uuid.UUID]*websocket.Conn),
	}
}

// NewMessage is a constructor for message entity.
func NewMessage(status int, text string) *Message {
	return &Message{
		Status: status,
		Text:   text,
	}
}

// ReadSearch reads user's request for search.
func (h *Hub) ReadSearch(client Client) (bool, error) {
	var (
		request Request
		err     error
		message *Message
	)

	if err = client.Conn.ReadJSON(&request); err != nil {
		message = NewMessage(http.StatusBadRequest, err.Error())
		if err = h.SendMessage(client, *message); err != nil {
			return false, ErrWriteHub.Wrap(err)
		}
		return false, ErrHub.Wrap(err)
	}

	if request.Action != ActionSearch {
		message = NewMessage(http.StatusBadRequest, "action not search!")
		if err = h.SendMessage(client, *message); err != nil {
			return false, ErrWriteHub.Wrap(err)
		}
		return false, ErrHub.New("action not search!")
	}

	if request.Value {
		if ok := h.isAdded(client); ok {
			if err = h.SendMessage(client, *NewMessage(http.StatusInternalServerError, "you have already been added!")); err != nil {
				return false, ErrWriteHub.Wrap(err)
			}
			return false, nil
		}
		if err = h.addClient(client); err != nil {
			return false, ErrHub.Wrap(err)
		}
		return true, nil
	}

	if ok := h.isAdded(client); !ok {
		if err = h.SendMessage(client, *NewMessage(http.StatusInternalServerError, "you don't have been added!")); err != nil {
			return false, ErrWriteHub.Wrap(err)
		}
		return false, nil
	}

	if err = h.RemoveClient(client); err != nil {
		return false, ErrHub.Wrap(err)
	}

	defer func() {
		err = errs.Combine(err, ErrHub.Wrap(client.Conn.Close()))
	}()
	if err != nil {
		return false, err
	}

	return false, nil

}

// ReadPlay reads user's request for play.
func (h *Hub) ReadPlay(userID uuid.UUID) (bool, error) {
	var (
		request Request
		err     error
		message *Message
		client  Client
	)

	client = Client{
		UserID: userID,
		Conn:   h.Clients[userID],
	}

	if err = client.Conn.ReadJSON(&request); err != nil {
		message = NewMessage(http.StatusBadRequest, err.Error())
		if err = h.SendMessage(client, *message); err != nil {
			return false, ErrWriteHub.Wrap(err)
		}
		return false, ErrHub.Wrap(err)
	}

	if request.Action == ActionPlay {
		return request.Value, nil
	}

	message = NewMessage(http.StatusBadRequest, "action not play!")
	if err = h.SendMessage(client, *message); err != nil {
		return false, ErrWriteHub.Wrap(err)
	}
	return false, ErrHub.New("action not play!")
}

// addClient addes user to hub.
func (h *Hub) addClient(client Client) error {
	h.Clients[client.UserID] = client.Conn
	return h.SendMessage(client, *NewMessage(http.StatusOK, "you added!"))
}

// RemoveClient removes user from hub.
func (h *Hub) RemoveClient(client Client) error {
	delete(h.Clients, client.UserID)
	return h.SendMessage(client, *NewMessage(http.StatusOK, "you leaved!"))
}

// SendMessage sends message to user.
func (h *Hub) SendMessage(client Client, message Message) error {
	return client.Conn.WriteJSON(message)
}

// isAdded checks if the user has been added.
func (h *Hub) isAdded(client Client) bool {
	_, ok := h.Clients[client.UserID]
	return ok
}
