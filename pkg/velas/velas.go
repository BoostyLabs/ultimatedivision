// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package velas

// APISResponse for velas response.
type APISResponse struct {
	State               string `json:"state"`
	Stage               string `json:"stage"`
	AccessToken         string `json:"access_token"`
	ExpiresAt           int64  `json:"expires_at"`
	TokenType           string `json:"token_type"`
	AuthorizedChallenge string `json:"authorized_challenge"`
	AccessTokenPayload  struct {
		Iss                        string   `json:"iss"`
		Aud                        string   `json:"aud"`
		Sub                        string   `json:"sub"`
		Ses                        string   `json:"ses"`
		Scopes                     []string `json:"scopes"`
		AuthTime                   int64    `json:"auth_time"`
		TransactionsSponsorAPIHost string   `json:"transactions_sponsor_api_host"`
		TransactionsSponsorPubKey  string   `json:"transactions_sponsor_pub_key"`
		Exp                        int64    `json:"exp"`
		Iat                        int64    `json:"iat"`
		Type                       string   `json:"type"`
	} `json:"access_token_payload"`
}

// VAClientFields for velas va client fields from config.
type VAClientFields struct {
	ClientID                   string `json:"client_id"`
	RedirectURI                string `json:"redirectUri"`
	AccountProviderHost        string `json:"accountProviderHost"`
	NetworkAPIHost             string `json:"networkApiHost"`
	TransactionsSponsorAPIHost string `json:"transactionsSponsorApiHost"`
	TransactionsSponsorPubKey  string `json:"transactionsSponsorPubKey"`
}

// Config defines configuration for velas va client.
type Config struct {
	ClientID                   string `json:"clientId"`
	RedirectURI                string `json:"redirectUri"`
	AccountProviderHost        string `json:"accountProviderHost"`
	NetworkAPIHost             string `json:"networkApiHost"`
	TransactionsSponsorAPIHost string `json:"transactionsSponsorApiHost"`
	TransactionsSponsorPubKey  string `json:"transactionsSponsorPubKey"`
}
