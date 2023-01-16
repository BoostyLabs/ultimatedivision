// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nfts

import (
	"bufio"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/BoostyLabs/evmsignature"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	contract "ultimatedivision/pkg/contractcasper"
	eventparsing "ultimatedivision/pkg/evetparsing"
	"ultimatedivision/pkg/nft"
)

// ErrNFTs indicated that there was an error in service.
var ErrNFTs = errs.Class("NFTs service error")

// Service is handling NFTs related logic.
//
// architecture: Service.
type Service struct {
	gctx   context.Context
	config Config
	casper contract.Casper
	nfts   DB
	events *http.Client
}

// NewService is a constructor for NFTs service.
func NewService(gctx context.Context, config Config, nfts DB) *Service {
	eventsClient := &http.Client{
		Transport: &http.Transport{
			DisableCompression: true,
		},
	}

	return &Service{
		gctx:   gctx,
		config: config,
		nfts:   nfts,
		events: eventsClient,
	}
}

const (
	// WriteCLValueKey defines that transform key is WriteCLValue. This key stores the type and data of the transforming event.
	WriteCLValueKey string = "WriteCLValue"
	// BytesKey defines that WriteCLValue key is bytes. This key stores data of the transforming event.
	BytesKey string = "bytes"
)

// Create creates nft in the database.
func (service *Service) Create(ctx context.Context, nft NFT) error {
	return ErrNFTs.Wrap(service.nfts.Create(ctx, nft))
}

// Generate generates values for nft token.
func (service *Service) Generate(ctx context.Context, card cards.Card, avatarURL string) nft.NFT {
	var attributes []nft.Attribute

	attributes = append(attributes, nft.Attribute{TraitType: "Id", Value: card.ID.String()})
	attributes = append(attributes, nft.Attribute{TraitType: "Quality", Value: card.Quality})
	attributes = append(attributes, nft.Attribute{TraitType: "Height", Value: fmt.Sprintf("%.2f", card.Height)})
	attributes = append(attributes, nft.Attribute{TraitType: "Weight", Value: fmt.Sprintf("%.2f", card.Weight)})
	attributes = append(attributes, nft.Attribute{TraitType: "Dominant Foot", Value: card.DominantFoot})

	// Game parameters.
	attributes = append(attributes, nft.Attribute{TraitType: "Tactics", Value: card.Tactics, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Positioning", Value: card.Positioning, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Composure", Value: card.Composure, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Aggression", Value: card.Aggression, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Vision", Value: card.Vision, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Awareness", Value: card.Awareness, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Crosses", Value: card.Crosses, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Physique", Value: card.Physique, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Acceleration", Value: card.Acceleration, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Running Speed", Value: card.RunningSpeed, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Reaction Speed", Value: card.ReactionSpeed, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Agility", Value: card.Agility, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Stamina", Value: card.Stamina, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Strength", Value: card.Strength, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Jumping", Value: card.Jumping, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Balance", Value: card.Balance, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Technique", Value: card.Technique, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Dribbling", Value: card.Dribbling, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Ball Control", Value: card.BallControl, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Weak Foot", Value: card.WeakFoot, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Skill Moves", Value: card.SkillMoves, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Finesse", Value: card.Finesse, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Curve", Value: card.Curve, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Volleys", Value: card.Volleys, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Short Passing", Value: card.ShortPassing, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Long Passing", Value: card.LongPassing, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Forward Pass", Value: card.ForwardPass, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Offence", Value: card.Offence, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Finishing Ability", Value: card.FinishingAbility, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Shot Power", Value: card.ShotPower, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Accuracy", Value: card.Accuracy, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Distance", Value: card.Distance, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Penalty", Value: card.Penalty, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Free Kicks", Value: card.FreeKicks, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Corners", Value: card.Corners, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Heading Accuracy", Value: card.HeadingAccuracy, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Defence", Value: card.Defence, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Offside Trap", Value: card.OffsideTrap, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Sliding", Value: card.Sliding, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Tackles", Value: card.Tackles, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Ball Focus", Value: card.BallFocus, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Interceptions", Value: card.Interceptions, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Vigilance", Value: card.Vigilance, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Goalkeeping", Value: card.Goalkeeping, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Reflexes", Value: card.Reflexes, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Diving", Value: card.Diving, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Handling", Value: card.Handling, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Sweeping", Value: card.Sweeping, MaxValue: MaxValueGameParameter})
	attributes = append(attributes, nft.Attribute{TraitType: "Throwing", Value: card.Throwing, MaxValue: MaxValueGameParameter})

	nft := nft.NFT{
		Attributes:  attributes,
		Description: service.config.Description,
		ExternalURL: fmt.Sprintf(service.config.ExternalURL, card.ID.String()),
		Image:       avatarURL,
		Name:        card.PlayerName,
	}

	return nft
}

// Get returns nft by token id and chain from database.
func (service *Service) Get(ctx context.Context, tokenID int64, chain evmsignature.Chain) (NFT, error) {
	nft, err := service.nfts.Get(ctx, tokenID, chain)
	return nft, ErrNFTs.Wrap(err)
}

// List returns nfts from database.
func (service *Service) List(ctx context.Context) ([]NFT, error) {
	nfts, err := service.nfts.List(ctx)
	return nfts, ErrNFTs.Wrap(err)
}

// Update updates users wallet address for nft token in the database.
func (service *Service) Update(ctx context.Context, nft NFT) error {
	return ErrNFTs.Wrap(service.nfts.Update(ctx, nft))
}

// Delete deletes nft token in the database.
func (service *Service) Delete(ctx context.Context, cardID uuid.UUID) error {
	return ErrNFTs.Wrap(service.nfts.Delete(ctx, cardID))
}

// SubscribeEvents is real time events streaming from blockchain to events subscribers.
func (service *Service) SubscribeEvents(ctx context.Context) (EventVariant, error) {
	var body io.Reader
	req, err := http.NewRequest(http.MethodGet, service.config.EventNodeAddress, body)
	if err != nil {
		return EventVariant{}, ErrNFTs.Wrap(err)
	}

	resp, err := service.events.Do(req)
	if err != nil {
		defer func() {
			err = errs.Combine(err, resp.Body.Close())
		}()
		return EventVariant{}, ErrNFTs.Wrap(err)
	}

	for {
		select {
		case <-service.gctx.Done():
			return EventVariant{}, nil
		case <-ctx.Done():
			return EventVariant{}, nil
		default:
		}

		reader := bufio.NewReader(resp.Body)
		rawBody, err := reader.ReadBytes('\n')
		if err != nil {
			return EventVariant{}, ErrNFTs.Wrap(err)
		}

		rawBody = []byte(strings.Replace(string(rawBody), "data:", "", 1))

		var event contract.Event
		_ = json.Unmarshal(rawBody, &event)

		transforms := event.DeployProcessed.ExecutionResult.Success.Effect.Transforms
		if len(transforms) == 0 {
			continue
		}

		for _, transform := range transforms {
			select {
			case <-service.gctx.Done():
				return EventVariant{}, nil
			case <-ctx.Done():
				return EventVariant{}, nil
			default:
			}

			eventFunds, err := service.parseEventFromTransform(event, transform)
			if err != nil {
				return eventFunds, ErrNFTs.Wrap(err)
			}
		}
	}
}

func (service *Service) parseEventFromTransform(event contract.Event, transform contract.Transform) (EventVariant, error) {
	transformMap, ok := transform.Transform.(map[string]interface{})
	if !ok {
		return EventVariant{}, ErrNFTs.New("couldn't parse map to transform")
	}

	writeCLValue, ok := transformMap[WriteCLValueKey].(map[string]interface{})
	if !ok {
		return EventVariant{}, ErrNFTs.New("couldn't parse map to transform map")
	}

	bytes, ok := writeCLValue[BytesKey].(string)
	if !ok {
		return EventVariant{}, ErrNFTs.New("couldn't parse string to bytes key")
	}

	eventData := eventparsing.EventData{
		Bytes: bytes,
	}

	eventType, err := eventData.GetEventType()
	if err != nil {
		return EventVariant{}, ErrNFTs.Wrap(err)
	}

	tokenContractAddress, err := hex.DecodeString(eventData.GetTokenContractAddress())
	if err != nil {
		return EventVariant{}, ErrNFTs.Wrap(err)
	}

	chainName, err := eventData.GetChainName()
	if err != nil {
		return EventVariant{}, ErrNFTs.Wrap(err)
	}

	chainAddress, err := eventData.GetChainAddress()
	if err != nil {
		return EventVariant{}, ErrNFTs.Wrap(err)
	}

	amount, err := eventData.GetAmount()
	if err != nil {
		return EventVariant{}, ErrNFTs.Wrap(err)
	}
	amountStr := strconv.Itoa(amount)

	userWalletAddress, err := hex.DecodeString(eventData.GetUserWalletAddress())
	if err != nil {
		return EventVariant{}, ErrNFTs.Wrap(err)
	}

	hash, err := hex.DecodeString(event.DeployProcessed.DeployHash)
	if err != nil {
		return EventVariant{}, ErrNFTs.Wrap(err)
	}

	sender, err := hex.DecodeString(event.DeployProcessed.Account)
	if err != nil {
		return EventVariant{}, ErrNFTs.Wrap(err)
	}

	blockNumber, err := service.casper.GetBlockNumberByHash(event.DeployProcessed.BlockHash)
	if err != nil {
		return EventVariant{}, ErrNFTs.Wrap(err)
	}

	transactionInfo := TransactionInfo{
		Hash:        hash,
		BlockNumber: uint64(blockNumber),
		Sender:      sender,
	}

	var eventFunds EventVariant
	switch eventType {
	case EventTypeIn.Int():
		eventFunds = EventVariant{
			Type: EventType(eventType),
			EventFundsIn: EventFundsIn{
				From: userWalletAddress,
				To: Address{
					NetworkName: chainName,
					Address:     chainAddress,
				},
				Amount: amountStr,
				Token:  tokenContractAddress,
				Tx:     transactionInfo,
			},
		}
	case EventTypeOut.Int():
		eventFunds = EventVariant{
			Type: EventType(eventType),
			EventFundsOut: EventFundsOut{
				From: Address{
					NetworkName: chainName,
					Address:     chainAddress,
				},
				To:     userWalletAddress,
				Amount: amountStr,
				Token:  tokenContractAddress,
				Tx:     transactionInfo,
			},
		}
	default:
		return EventVariant{}, ErrNFTs.New("invalid event type")
	}

	tokenIn := hex.EncodeToString(eventFunds.EventFundsIn.Token)
	eventFunds.EventFundsIn.Token, err = hex.DecodeString(eventparsing.TagHash.String() + tokenIn)
	if err != nil {
		return EventVariant{}, ErrNFTs.Wrap(err)
	}

	from := hex.EncodeToString(eventFunds.EventFundsIn.From)
	eventFunds.EventFundsIn.From, err = hex.DecodeString(eventparsing.TagAccount.String() + from)
	if err != nil {
		return EventVariant{}, ErrNFTs.Wrap(err)
	}

	tokenOut := hex.EncodeToString(eventFunds.EventFundsOut.Token)
	eventFunds.EventFundsOut.Token, err = hex.DecodeString(eventparsing.TagHash.String() + tokenOut)
	if err != nil {
		return EventVariant{}, ErrNFTs.Wrap(err)
	}

	to := hex.EncodeToString(eventFunds.EventFundsOut.To)
	eventFunds.EventFundsOut.To, err = hex.DecodeString(eventparsing.TagAccount.String() + to)
	if err != nil {
		return EventVariant{}, ErrNFTs.Wrap(err)
	}

	return eventFunds, nil
}
