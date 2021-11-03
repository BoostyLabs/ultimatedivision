// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nfts

// NFT entity describes nft token format erc-721.
type NFT struct {
	Attributes  []Attribute `json:"attributes"`
	Description string      `json:"description"`
	ExternalURL string      `json:"external_url"`
	Image       string      `json:"image"`
	Name        string      `json:"name"`
}

// Attribute entity describes attributes for nft token.
type Attribute struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
	MaxValue  interface{} `json:"max_value"`
}

// MaxValueGameParameter indicates that max value game parameter is 100.
const MaxValueGameParameter = 100
