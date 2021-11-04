// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package avatarcards

import (
	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/cards/nfts"
)

// CardWithLinkToAvatar describes card entity with link to avatar.
type CardWithLinkToAvatar struct {
	cards.Card
	OriginalURL string `json:"originalUrl"`
}

// Config defines values needed to generate card with avatar.
type Config struct {
	CardConfig           cards.Config              `json:"cardConfig"`
	PercentageQualities  cards.PercentageQualities `json:"percentageQualities"`
	AvatarConfig         avatars.Config            `json:"avatarConfig"`
	NFTConfig            nfts.Config               `json:"nftConfig"`
	PathToOutputJSONFile string                    `json:"pathToOutputJsonFile"`
	PathToNamesDataset   string                    `json:"pathToNamesDataset"`
}
