// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nfts

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
)

// ErrNFTs indicated that there was an error in service.
var ErrNFTs = errs.Class("NFTs service error")

// Service is handling NFTs related logic.
//
// architecture: Service
type Service struct {
	cards *cards.Service
}

// NewService is a constructor for NFTs service.
func NewService(cards *cards.Service) *Service {
	return &Service{
		cards: cards,
	}
}

// GenerateNFT generates values for nft token.
func (service *Service) GenerateNFT(ctx context.Context, cardID uuid.UUID) (NFT, error) {
	card, err := service.cards.Get(ctx, cardID)
	if err != nil {
		return NFT{}, ErrNFTs.Wrap(err)
	}
	var attributes []Attribute
	attributes = append(attributes, Attribute{TraitType: "Id", Value: card.ID.String()})
	attributes = append(attributes, Attribute{TraitType: "Quality", Value: card.Quality})
	attributes = append(attributes, Attribute{TraitType: "Height", Value: fmt.Sprintf("%f", card.Height)})
	attributes = append(attributes, Attribute{TraitType: "Weight", Value: fmt.Sprintf("%f", card.Weight)})
	attributes = append(attributes, Attribute{TraitType: "Dominant Foot", Value: card.DominantFoot})

	// Game parameters
	attributes = append(attributes, Attribute{TraitType: "Tactics", Value: card.Tactics, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Positioning", Value: card.Positioning, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Composure", Value: card.Composure, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Aggression", Value: card.Aggression, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Vision", Value: card.Vision, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Awareness", Value: card.Awareness, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Crosses", Value: card.Crosses, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Physique", Value: card.Physique, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Acceleration", Value: card.Acceleration, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Running Speed", Value: card.RunningSpeed, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Reaction Speed", Value: card.ReactionSpeed, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Agility", Value: card.Agility, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Stamina", Value: card.Stamina, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Strength", Value: card.Strength, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Jumping", Value: card.Jumping, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Balance", Value: card.Balance, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Technique", Value: card.Technique, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Dribbling", Value: card.Dribbling, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Ball Control", Value: card.BallControl, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Weak Foot", Value: card.WeakFoot, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Skill Moves", Value: card.SkillMoves, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Finesse", Value: card.Finesse, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Curve", Value: card.Curve, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Volleys", Value: card.Volleys, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Short Passing", Value: card.ShortPassing, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Long Passing", Value: card.LongPassing, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Forward Pass", Value: card.ForwardPass, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Offence", Value: card.Offence, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Finishing Ability", Value: card.FinishingAbility, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Shot Power", Value: card.ShotPower, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Accuracy", Value: card.Accuracy, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Distance", Value: card.Distance, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Penalty", Value: card.Penalty, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Free Kicks", Value: card.FreeKicks, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Corners", Value: card.Corners, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Heading Accuracy", Value: card.HeadingAccuracy, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Defence", Value: card.Defence, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Offside Trap", Value: card.OffsideTrap, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Sliding", Value: card.Sliding, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Tackles", Value: card.Tackles, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Ball Focus", Value: card.BallFocus, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Interceptions", Value: card.Interceptions, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Vigilance", Value: card.Vigilance, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Goalkeeping", Value: card.Goalkeeping, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Reflexes", Value: card.Reflexes, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Diving", Value: card.Diving, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Handling", Value: card.Handling, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Sweeping", Value: card.Sweeping, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})
	attributes = append(attributes, Attribute{TraitType: "Throwing", Value: card.Throwing, MaxValue: MaxValueGameParameter, DisplayType: DisplayTypeBoostNumber})

	nft := NFT{
		Attributes:  attributes,
		Description: "",
		ExternalURL: "",
		Image:       "",
		Name:        card.PlayerName,
	}
	return nft, nil
}
