// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nfts

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/pkg/cryptoutils"
	"ultimatedivision/users"
)

// ErrNFTs indicated that there was an error in service.
var ErrNFTs = errs.Class("NFTs service error")

// Service is handling NFTs related logic.
//
// architecture: Service
type Service struct {
	storage Storage
	cards   *cards.Service
	avatars *avatars.Service
	users   *users.Service
}

// NewService is a constructor for NFTs service.
func NewService(storage Storage, cards *cards.Service, avatars *avatars.Service, users *users.Service) *Service {
	return &Service{
		storage: storage,
		cards:   cards,
		avatars: avatars,
		users:   users,
	}
}

func (service *Service) Create(ctx context.Context, cardID uuid.UUID, wallet cryptoutils.Address, userID uuid.UUID) error {
	card, err := service.cards.Get(ctx, cardID)
	if err != nil {
		return ErrNFTs.Wrap(err)
	}

	avatar, err := service.avatars.Get(ctx, cardID)
	if err != nil {
		return ErrNFTs.Wrap(err)
	}

	avatarURL, err := service.avatars.Save(ctx, avatar)
	if err != nil {
		return ErrNFTs.Wrap(err)
	}

	nft, err := service.Generate(ctx, card, avatarURL)
	if err != nil {
		return ErrNFTs.Wrap(err)
	}

	if err = service.Save(ctx, nft); err != nil {
		return ErrNFTs.Wrap(err)
	}

	return service.users.UpdateWalletAddress(ctx, wallet, userID)
}

// Generate generates values for nft token.
func (service *Service) Generate(ctx context.Context, card cards.Card, avatarURL string) (NFT, error) {
	var attributes []Attribute

	attributes = append(attributes, Attribute{TraitType: "Id", Value: card.ID.String()})
	attributes = append(attributes, Attribute{TraitType: "Quality", Value: card.Quality})
	attributes = append(attributes, Attribute{TraitType: "Height", Value: fmt.Sprintf("%f", card.Height)})
	attributes = append(attributes, Attribute{TraitType: "Weight", Value: fmt.Sprintf("%f", card.Weight)})
	attributes = append(attributes, Attribute{TraitType: "Dominant Foot", Value: card.DominantFoot})

	// Game parameters
	attributes = append(attributes, Attribute{TraitType: "Tactics", Value: card.Tactics, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Positioning", Value: card.Positioning, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Composure", Value: card.Composure, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Aggression", Value: card.Aggression, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Vision", Value: card.Vision, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Awareness", Value: card.Awareness, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Crosses", Value: card.Crosses, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Physique", Value: card.Physique, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Acceleration", Value: card.Acceleration, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Running Speed", Value: card.RunningSpeed, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Reaction Speed", Value: card.ReactionSpeed, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Agility", Value: card.Agility, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Stamina", Value: card.Stamina, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Strength", Value: card.Strength, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Jumping", Value: card.Jumping, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Balance", Value: card.Balance, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Technique", Value: card.Technique, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Dribbling", Value: card.Dribbling, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Ball Control", Value: card.BallControl, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Weak Foot", Value: card.WeakFoot, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Skill Moves", Value: card.SkillMoves, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Finesse", Value: card.Finesse, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Curve", Value: card.Curve, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Volleys", Value: card.Volleys, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Short Passing", Value: card.ShortPassing, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Long Passing", Value: card.LongPassing, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Forward Pass", Value: card.ForwardPass, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Offence", Value: card.Offence, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Finishing Ability", Value: card.FinishingAbility, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Shot Power", Value: card.ShotPower, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Accuracy", Value: card.Accuracy, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Distance", Value: card.Distance, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Penalty", Value: card.Penalty, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Free Kicks", Value: card.FreeKicks, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Corners", Value: card.Corners, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Heading Accuracy", Value: card.HeadingAccuracy, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Defence", Value: card.Defence, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Offside Trap", Value: card.OffsideTrap, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Sliding", Value: card.Sliding, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Tackles", Value: card.Tackles, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Ball Focus", Value: card.BallFocus, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Interceptions", Value: card.Interceptions, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Vigilance", Value: card.Vigilance, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Goalkeeping", Value: card.Goalkeeping, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Reflexes", Value: card.Reflexes, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Diving", Value: card.Diving, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Handling", Value: card.Handling, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Sweeping", Value: card.Sweeping, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, Attribute{TraitType: "Throwing", Value: card.Throwing, MaxValue: MaxValueGameParameter})

	nft := NFT{
		Attributes:  attributes,
		Description: "",
		ExternalURL: "",
		Image:       avatarURL,
		Name:        card.PlayerName,
	}
	return nft, nil
}

// Save saves nft in the storage.
func (service *Service) Save(ctx context.Context, nft NFT) error {
	return ErrNFTs.Wrap(service.storage.Save(ctx, nft))
}
