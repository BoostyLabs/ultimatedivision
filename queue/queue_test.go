// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/queue"
	"ultimatedivision/users"
)

func TestQueue(t *testing.T) {
	user1 := users.User{
		ID:           uuid.New(),
		Email:        "tarkovskynik@gmail.com",
		PasswordHash: []byte{0},
		NickName:     "Nik",
		FirstName:    "Nikita",
		LastName:     "Tarkovskyi",
		LastLogin:    time.Now(),
		Status:       0,
		CreatedAt:    time.Now(),
	}

	user2 := users.User{
		ID:           uuid.New(),
		Email:        "3560876@gmail.com",
		PasswordHash: []byte{1},
		NickName:     "qwerty",
		FirstName:    "Stas",
		LastName:     "Isakov",
		LastLogin:    time.Now(),
		Status:       1,
		CreatedAt:    time.Now(),
	}

	queueClient1 := queue.Client{UserID: user1.ID, Conn: nil}
	queueClient2 := queue.Client{UserID: user2.ID, Conn: nil}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryQueue := db.Queue()
		repositoryUsers := db.Users()
		userID := uuid.New()

		t.Run("get sql no rows", func(t *testing.T) {
			_, err := repositoryQueue.Get(userID)
			require.Error(t, err)
			assert.Equal(t, true, queue.ErrNoClient.Has(err))
		})

		t.Run("get", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)

			repositoryQueue.Create(queueClient1)

			queueFromDB, err := repositoryQueue.Get(user1.ID)
			require.NoError(t, err)
			compareQueues(t, queueClient1, queueFromDB)
		})

		t.Run("list", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user2)
			require.NoError(t, err)

			repositoryQueue.Create(queueClient2)

			queueList := repositoryQueue.List()
			assert.Equal(t, len(queueList), 2)
			compareQueues(t, queueClient1, queueList[0])
			compareQueues(t, queueClient2, queueList[1])
		})

		t.Run("delete", func(t *testing.T) {
			repositoryQueue.Delete(queueClient1.UserID)

			queueList := repositoryQueue.List()
			assert.Equal(t, len(queueList), 1)
			compareQueues(t, queueClient2, queueList[0])
		})
	})

	userLogin1 := map[string]string{
		"email":    "dmytroyakimuk@gmail.com",
		"password": "1212-Dima",
	}
	userLogin2 := map[string]string{
		"email":    "dmytroyakimuk@gmail.com2",
		"password": "1212-Dima2",
	}

	var cookie1, cookie2 string

	t.Run("login user1", func(t *testing.T) {
		jsonData, err := json.Marshal(userLogin1)
		require.NoError(t, err)

		resp, err := http.Post("http://localhost:8088/api/v0/auth/login", "application/json", bytes.NewBuffer(jsonData))
		require.NoError(t, err)
		defer func() {
			err = resp.Body.Close()
			require.NoError(t, err)
		}()

		if resp.Status != "200 OK" {
			var res map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&res)
			require.NoError(t, err)
			require.NoError(t, fmt.Errorf(fmt.Sprint(res)))
		} else {
			cookie1 = resp.Cookies()[0].Value
		}
	})

	go t.Run("test queue1", func(t *testing.T) {
		urlObj, err := url.Parse("http://localhost:8088/")
		require.NoError(t, err)

		cookie := &http.Cookie{
			Name:     "ultimatedivision_console",
			Value:    cookie1,
			Path:     "/",
			Domain:   "localhost",
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Secure:   false,
		}

		jar, err := cookiejar.New(&cookiejar.Options{})
		if err != nil {
			require.NoError(t, err)
		}
		jar.SetCookies(urlObj, []*http.Cookie{cookie})
		var dialer = &websocket.Dialer{
			HandshakeTimeout: 4500000 * time.Second,
			Jar:              jar,
		}

		c, r, err := dialer.Dial("ws://localhost:8088/api/v0/queue", nil)
		require.NoError(t, err)
		defer func() {
			err = c.Close()
			err = r.Body.Close()
			require.NoError(t, err)
		}()

		request := queue.Request{Action: queue.ActionStartSearch}
		err = c.WriteJSON(request)
		require.NoError(t, err)

		var message queue.Response
		err = c.ReadJSON(&message)
		log.Println(message)
		require.NoError(t, err)
		assert.Equal(t, 200, message.Status)

		var message2 queue.Response
		err = c.ReadJSON(&message2)
		log.Println(message2)
		require.NoError(t, err)
		assert.Equal(t, 200, message2.Status)

		request2 := queue.Request{Action: queue.ActionConfirm}
		err = c.WriteJSON(request2)
		require.NoError(t, err)
	})

	t.Run("login user2", func(t *testing.T) {
		jsonData, err := json.Marshal(userLogin2)
		require.NoError(t, err)

		resp, err := http.Post("http://localhost:8088/api/v0/auth/login", "application/json", bytes.NewBuffer(jsonData))
		require.NoError(t, err)
		defer func() {
			err = resp.Body.Close()
			require.NoError(t, err)
		}()

		if resp.Status != "200 OK" {
			var res map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&res)
			require.NoError(t, err)
			require.NoError(t, fmt.Errorf(fmt.Sprint(res)))
		} else {
			cookie2 = resp.Cookies()[0].Value
		}
	})

	t.Run("test queue2", func(t *testing.T) {
		urlObj, err := url.Parse("http://localhost:8088/")
		require.NoError(t, err)

		cookie := &http.Cookie{
			Name:     "ultimatedivision_console",
			Value:    cookie2,
			Path:     "/",
			Domain:   "localhost",
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Secure:   false,
		}

		jar, err := cookiejar.New(&cookiejar.Options{})
		if err != nil {
			require.NoError(t, err)
		}
		jar.SetCookies(urlObj, []*http.Cookie{cookie})
		var dialer = &websocket.Dialer{
			HandshakeTimeout: 4500000 * time.Second,
			Jar:              jar,
		}

		c, r, err := dialer.Dial("ws://localhost:8088/api/v0/queue", nil)
		require.NoError(t, err)
		defer func() {
			err = c.Close()
			err = r.Body.Close()
			require.NoError(t, err)
		}()

		request := queue.Request{Action: queue.ActionStartSearch}
		err = c.WriteJSON(request)
		require.NoError(t, err)

		var message queue.Response
		err = c.ReadJSON(&message)
		log.Println(message)
		require.NoError(t, err)
		assert.Equal(t, 200, message.Status)

		var message2 queue.Response
		err = c.ReadJSON(&message2)
		log.Println(message2)
		require.NoError(t, err)
		assert.Equal(t, 200, message2.Status)

		request2 := queue.Request{Action: queue.ActionConfirm}
		err = c.WriteJSON(request2)
		require.NoError(t, err)
	})
}

func compareQueues(t *testing.T, queue1, queue2 queue.Client) {
	assert.Equal(t, queue1.UserID, queue2.UserID)
}
