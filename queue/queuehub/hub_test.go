// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queuehub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueueHub(t *testing.T) {
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

		request := Request{
			Action: "search",
			Value:  true,
		}
		err = c.WriteJSON(request)
		require.NoError(t, err)

		var message Message
		err = c.ReadJSON(&message)
		require.NoError(t, err)
		assert.Equal(t, 200, message.Status)

		var message2 Message
		err = c.ReadJSON(&message2)
		require.NoError(t, err)
		assert.Equal(t, 200, message2.Status)

		request2 := Request{
			Action: ActionPlay,
			Value:  true,
		}
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

		request := Request{
			Action: "search",
			Value:  true,
		}
		err = c.WriteJSON(request)
		require.NoError(t, err)

		var message Message
		err = c.ReadJSON(&message)
		require.NoError(t, err)
		assert.Equal(t, 200, message.Status)

		var message2 Message
		err = c.ReadJSON(&message2)
		require.NoError(t, err)
		assert.Equal(t, 200, message2.Status)

		request2 := Request{
			Action: ActionPlay,
			Value:  true,
		}
		err = c.WriteJSON(request2)
		require.NoError(t, err)
	})
}
