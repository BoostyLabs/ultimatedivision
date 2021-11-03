// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nfts

// NFT entity describes nft token format erc-721.
type NFT struct {
	Attributes  []Attribut `json:"attributes"`
	Description string     `json:"description"`
	ExternalURL string     `json:"external_url"`
	Image       string     `json:"image"`
	Name        string     `json:"name"`
}

// Attribut entity describes attributes for nft token.
type Attribut struct {
	TraitType   string      `json:"trait_type"`
	Value       interface{} `json:"value"`
	MaxValue    interface{} `json:"max_value"`
	DisplayType DisplayType `json:"display_type,omitempty"`
}

// DisplayType defines the list of possible attribut types.
type DisplayType string

const (
	// DisplayTypeNumber indicates that the attribut is number.
	DisplayTypeNumber DisplayType = "number"
	// DisplayTypeBoostPercentage indicates that the attribut is boost percentage.
	DisplayTypeBoostPercentage DisplayType = "boost_percentage"
	// DisplayTypeBoostNumber indicates that the attribut is boost number.
	DisplayTypeBoostNumber DisplayType = "boost_number"
)

// MaxValueGameParameter indicates that max value game parameter is 100.
const MaxValueGameParameter = 100
