// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package waitlist

import (
	"bufio"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/internal/remotefilestorage/storj"
	contract "ultimatedivision/pkg/contractcasper"
	"ultimatedivision/pkg/eventparsing"
	"ultimatedivision/pkg/imageprocessing"
	"ultimatedivision/users"
)

// ErrWaitlist indicated that there was an error in service.
var ErrWaitlist = errs.Class("waitlist service error")

// Service is handling waitList related logic.
//
// architecture: Service
type Service struct {
	config   Config
	waitList DB
	casper   contract.Casper
	cards    *cards.Service
	avatars  *avatars.Service
	users    *users.Service
	nfts     *nfts.Service
	events   *http.Client

	gctx context.Context
}

// NewService is a constructor for waitlist service.
func NewService(config Config, waitList DB, cards *cards.Service, avatars *avatars.Service, users *users.Service, nfts *nfts.Service) *Service {
	eventsClient := &http.Client{
		Transport: &http.Transport{
			DisableCompression: true,
		},
	}
	ctx, cancel := context.WithCancel(context.Background())
	onSigInt(func() {
		// starting graceful exit on context cancellation.
		cancel()
	})
	return &Service{
		config:   config,
		waitList: waitList,
		cards:    cards,
		avatars:  avatars,
		users:    users,
		nfts:     nfts,
		events:   eventsClient,
		gctx:     ctx,
	}
}

// Create creates nft for wait list.
func (service *Service) Create(ctx context.Context, createNFT CreateNFT) (Transaction, error) {
	var transaction Transaction

	user, err := service.users.Get(ctx, createNFT.UserID)
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	if len(createNFT.WalletAddress.String()) == 0 {
		createNFT.WalletAddress = user.Wallet
	}

	if len(createNFT.CasperWallet) == 0 {
		createNFT.CasperWallet = user.CasperWallet
	}

	card, err := service.cards.Get(ctx, createNFT.CardID)
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	if createNFT.Value.Cmp(big.NewInt(0)) <= 0 {
		if card.UserID != createNFT.UserID {
			return transaction, ErrWaitlist.New("this card does not belongs to user")
		}
	}

	if item, err := service.GetByCardID(ctx, createNFT.CardID); item.Password != "" && err == nil {
		switch item.WalletType {
		case users.WalletTypeVelas:
			transaction = Transaction{
				Password:          item.Password,
				NFTCreateContract: NFTCreateContract(service.config.NFTCreateVelasContract),
				TokenID:           item.TokenID,
				Value:             item.Value,
				WalletType:        item.WalletType,
			}
		case users.WalletTypeCasper:
			transaction = Transaction{
				Password:                item.Password,
				NFTCreateCasperContract: service.config.NFTCreateCasperContract,
				TokenID:                 item.TokenID,
				Value:                   item.Value,
				WalletType:              item.WalletType,
				RPCNodeAddress:          service.config.RPCNodeAddress,
			}
		default:
			transaction = Transaction{
				Password:          item.Password,
				NFTCreateContract: service.config.NFTCreateContract,
				TokenID:           item.TokenID,
				Value:             item.Value,
				WalletType:        item.WalletType,
			}
		}

		return transaction, nil
	}

	image, err := service.avatars.GetImage(ctx, createNFT.CardID)
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	client, err := storj.NewClient(service.config.FileStorage)
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	// TODO: add transaction and mb lock db.
	lastTokenID, err := service.GetLastTokenID(ctx)
	if err != nil {
		if !ErrNoItem.Has(err) {
			return transaction, ErrWaitlist.Wrap(err)
		}
	}

	nextTokenID := lastTokenID + 1

	if err = client.Upload(ctx, service.config.Bucket, fmt.Sprintf("%d.%s", nextTokenID, imageprocessing.TypeFilePNG), image); err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	nft := service.nfts.Generate(ctx, card, fmt.Sprintf(service.config.URLToAvatar, nextTokenID))
	fileMetadata, err := json.MarshalIndent(nft, "", " ")
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	if err = client.Upload(ctx, service.config.Bucket, fmt.Sprintf("%d.%s", nextTokenID, imageprocessing.TypeFileJSON), fileMetadata); err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	if user.WalletType != users.WalletTypeCasper {
		if err = service.users.UpdateWalletAddress(ctx, createNFT.WalletAddress, createNFT.UserID, user.WalletType); err != nil {
			if !users.ErrWalletAddressAlreadyInUse.Has(err) {
				return transaction, ErrWaitlist.Wrap(err)
			}
		}
	}

	item := Item{
		TokenID:          uuid.New(),
		CardID:           createNFT.CardID,
		Wallet:           createNFT.WalletAddress,
		WalletType:       user.WalletType,
		CasperWallet:     createNFT.CasperWallet,
		CasperWalletHash: user.CasperWalletHash,
		Value:            createNFT.Value,
	}

	if err = service.waitList.Create(ctx, item); err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	for range time.NewTicker(time.Millisecond * service.config.WaitListCheckSignature).C {
		if item, err := service.GetByCardID(ctx, createNFT.CardID); item.Password != "" && err == nil {
			switch item.WalletType {
			case users.WalletTypeVelas:
				transaction = Transaction{
					Password:          item.Password,
					NFTCreateContract: NFTCreateContract(service.config.NFTCreateVelasContract),
					TokenID:           item.TokenID,
					Value:             item.Value,
					WalletType:        item.WalletType,
				}
			case users.WalletTypeCasper:
				transaction = Transaction{
					Password:                item.Password,
					NFTCreateCasperContract: service.config.NFTCreateCasperContract,
					TokenID:                 item.TokenID,
					Value:                   item.Value,
					WalletType:              item.WalletType,
					RPCNodeAddress:          service.config.RPCNodeAddress,
				}
			default:
				transaction = Transaction{
					Password:          item.Password,
					NFTCreateContract: service.config.NFTCreateContract,
					TokenID:           item.TokenID,
					Value:             item.Value,
					WalletType:        item.WalletType,
				}
			}
			break
		}
	}

	return transaction, err
}

// GetByTokenID returns nft for wait list by token id.
func (service *Service) GetByTokenID(ctx context.Context, tokenNumber int64) (Item, error) {
	nft, err := service.waitList.GetByTokenID(ctx, tokenNumber)
	return nft, ErrWaitlist.Wrap(err)
}

// GetByCardID returns nft for wait list by card id.
func (service *Service) GetByCardID(ctx context.Context, cardID uuid.UUID) (Item, error) {
	nft, err := service.waitList.GetByCardID(ctx, cardID)
	return nft, ErrWaitlist.Wrap(err)
}

// GetLastTokenID returns id of latest nft for wait list.
func (service *Service) GetLastTokenID(ctx context.Context) (int64, error) {
	lastID, err := service.waitList.GetLastTokenID(ctx)
	return lastID, ErrWaitlist.Wrap(err)
}

// List returns all nft for wait list.
func (service *Service) List(ctx context.Context) ([]Item, error) {
	allNFT, err := service.waitList.List(ctx)
	return allNFT, ErrWaitlist.Wrap(err)
}

// ListWithoutPassword returns nft for wait list without password.
func (service *Service) ListWithoutPassword(ctx context.Context) ([]Item, error) {
	nftWithoutPassword, err := service.waitList.ListWithoutPassword(ctx)
	return nftWithoutPassword, ErrWaitlist.Wrap(err)
}

// Update updates signature to nft token.
func (service *Service) Update(ctx context.Context, tokenID uuid.UUID, password evmsignature.Signature) error {
	return ErrWaitlist.Wrap(service.waitList.Update(ctx, tokenID, password))
}

// Delete deletes nft for wait list.
func (service *Service) Delete(ctx context.Context, tokenIDs []int64) error {
	return ErrWaitlist.Wrap(service.waitList.Delete(ctx, tokenIDs))
}

// SubscribeEvents is real time events streaming from blockchain to events subscribers.
func (service *Service) SubscribeEvents(ctx context.Context) (EventVariant, error) {
	var body io.Reader
	req, err := http.NewRequest(http.MethodGet, service.config.EventNodeAddress, body)
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
	}

	resp, err := service.events.Do(req)
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
	}
	fmt.Println("resp --> ", resp)
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
			return EventVariant{}, ErrWaitlist.Wrap(err)
		}

		rawBody = []byte(strings.Replace(string(rawBody), "data:", "", 1))

		var event contract.Event
		_ = json.Unmarshal(rawBody, &event)
		fmt.Println("event --> ", event)
		fmt.Println("DeployHash --> ", event.DeployProcessed.DeployHash)
		fmt.Println("Transforms --> ", event.DeployProcessed.ExecutionResult.Success.Effect.Transforms)
		fmt.Println("event --> ", event)
		transforms := event.DeployProcessed.ExecutionResult.Success.Effect.Transforms
		//if len(transforms) == 0 {
		//	return EventVariant{}, ErrWaitlist.Wrap(err)
		//}
		for _, transform := range transforms {
			eventFunds, err := service.parseEventFromTransform(event, transform)
			fmt.Println("eventFunds-->", eventFunds)
			if err != nil {
				return eventFunds, ErrWaitlist.Wrap(err)
			}
		}
	}
	//
	//fmt.Println("EventNodeAddress--> ", service.config.EventNodeAddress)
	//req, err := http.NewRequest(http.MethodGet, service.config.EventNodeAddress, body)
	//if err != nil {
	//	return EventVariant{}, ErrWaitlist.Wrap(err)
	//}
	//fmt.Println("req--->", req)
	//
	//resp, err := service.events.Do(req)
	//if err != nil {
	//	defer func() {
	//		err = errs.Combine(err, resp.Body.Close())
	//	}()
	//	fmt.Println("2222222222222222222 --> ERRROR")
	//	return EventVariant{}, ErrWaitlist.Wrap(err)
	//}
	//fmt.Println("resp --> ", resp)
	//reader := bufio.NewReader(resp.Body)
	//fmt.Println("reader --> ", reader)
	//rawBody, err := reader.ReadBytes('\n')
	//fmt.Println("rawBody --> ", rawBody)
	//if err != nil {
	//	return EventVariant{}, ErrWaitlist.Wrap(err)
	//}
	//
	//rawBody = []byte(strings.Replace(string(rawBody), "data:", "", 1))
	//
	//var event contract.Event
	//_ = json.Unmarshal(rawBody, &event)
	//
	//fmt.Println("event --> ", event)
	//
	//transforms := event.DeployProcessed.ExecutionResult.Success.Effect.Transforms
	//if len(transforms) == 0 {
	//	return EventVariant{}, ErrWaitlist.Wrap(err)
	//}
	//for _, transform := range transforms {
	//	eventFunds, err := service.parseEventFromTransform(event, transform)
	//	if err != nil {
	//		return eventFunds, ErrWaitlist.Wrap(err)
	//	}
	//}
	return EventVariant{}, ErrWaitlist.Wrap(err)
}
func (service *Service) parseEventFromTransform(event contract.Event, transform contract.Transform) (EventVariant, error) {
	transformMap, ok := transform.Transform.(map[string]interface{})
	if !ok {
		return EventVariant{}, ErrWaitlist.New("couldn't parse map to transform")
	}

	writeCLValue, ok := transformMap[WriteCLValueKey].(map[string]interface{})
	if !ok {
		return EventVariant{}, ErrWaitlist.New("couldn't parse map to transform map")
	}

	bytes, ok := writeCLValue[BytesKey].(string)
	if !ok {
		return EventVariant{}, ErrWaitlist.New("couldn't parse string to bytes key")
	}

	eventData := eventparsing.EventData{
		Bytes: bytes,
	}

	eventType, err := eventData.GetEventType()
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
	}

	fmt.Println("eventType--> ", eventType)

	tokenContractAddress, err := hex.DecodeString(eventData.GetTokenContractAddress())
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
	}

	fmt.Println("tokenContractAddress--> ", tokenContractAddress)

	chainName, err := eventData.GetChainName()
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
	}
	fmt.Println("chainName--> ", chainName)
	chainAddress, err := eventData.GetChainAddress()
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
	}

	amount, err := eventData.GetAmount()
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
	}
	amountStr := strconv.Itoa(amount)

	userWalletAddress, err := hex.DecodeString(eventData.GetUserWalletAddress())
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
	}

	fmt.Println("userWalletAddress--> ", userWalletAddress)

	hash, err := hex.DecodeString(event.DeployProcessed.DeployHash)
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
	}

	sender, err := hex.DecodeString(event.DeployProcessed.Account)
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
	}

	blockNumber, err := service.casper.GetBlockNumberByHash(event.DeployProcessed.BlockHash)
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
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
		return EventVariant{}, ErrWaitlist.New("invalid event type")
	}

	tokenIn := hex.EncodeToString(eventFunds.EventFundsIn.Token)
	eventFunds.EventFundsIn.Token, err = hex.DecodeString(eventparsing.TagHash.String() + tokenIn)
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
	}

	from := hex.EncodeToString(eventFunds.EventFundsIn.From)
	eventFunds.EventFundsIn.From, err = hex.DecodeString(eventparsing.TagAccount.String() + from)
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
	}

	tokenOut := hex.EncodeToString(eventFunds.EventFundsOut.Token)
	eventFunds.EventFundsOut.Token, err = hex.DecodeString(eventparsing.TagHash.String() + tokenOut)
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
	}

	to := hex.EncodeToString(eventFunds.EventFundsOut.To)
	eventFunds.EventFundsOut.To, err = hex.DecodeString(eventparsing.TagAccount.String() + to)
	if err != nil {
		return EventVariant{}, ErrWaitlist.Wrap(err)
	}

	return eventFunds, nil
}

// onSigInt fires in SIGINT or SIGTERM event (usually CTRL+C).
func onSigInt(onSigInt func()) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-done
		onSigInt()
	}()
}
