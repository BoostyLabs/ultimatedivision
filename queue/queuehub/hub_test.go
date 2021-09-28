// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queuehub

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision/users"
)

// UserDB entity describes user from DB.
type UserDB struct {
	Email    string `json:"email"`
	NickName string `json:"nickName"`
}

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
	var userDB1, userDB2 UserDB

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
		} else {
			cookie1 = resp.Cookies()[0].Value
		}
	})

	t.Run("get profile user1", func(t *testing.T) {
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
		client := &http.Client{
			Jar: jar,
		}

		resp, err := client.Get("http://localhost:8088/api/v0/profile")
		require.NoError(t, err)
		defer func() {
			err = resp.Body.Close()
			require.NoError(t, err)
		}()

		if resp.Status != "200 OK" {
			var res map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&res)
			require.NoError(t, err)
		} else {
			err = json.NewDecoder(resp.Body).Decode(&userDB1)
			require.NoError(t, err)
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

		var user users.User
		err = c.ReadJSON(&user)
		require.NoError(t, err)
		assert.Equal(t, userDB2.Email, user.Email)
		assert.Equal(t, userDB2.NickName, user.NickName)

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
		} else {
			cookie2 = resp.Cookies()[0].Value
		}
	})

	t.Run("get profile user2", func(t *testing.T) {
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
		client := &http.Client{
			Jar: jar,
		}

		resp, err := client.Get("http://localhost:8088/api/v0/profile")
		require.NoError(t, err)
		defer func() {
			err = resp.Body.Close()
			require.NoError(t, err)
		}()

		if resp.Status != "200 OK" {
			var res map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&res)
			require.NoError(t, err)
		} else {
			err = json.NewDecoder(resp.Body).Decode(&userDB2)
			require.NoError(t, err)
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

		var user users.User
		err = c.ReadJSON(&user)
		require.NoError(t, err)
		assert.Equal(t, userDB1.Email, user.Email)
		assert.Equal(t, userDB1.NickName, user.NickName)

		request2 := Request{
			Action: ActionPlay,
			Value:  true,
		}
		err = c.WriteJSON(request2)
		require.NoError(t, err)
	})
}
