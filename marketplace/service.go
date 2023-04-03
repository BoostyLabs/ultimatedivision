// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package marketplace

import (
	"context"
	"encoding/hex"
	"github.com/casper-ecosystem/casper-golang-sdk/keypair"
	"github.com/casper-ecosystem/casper-golang-sdk/sdk"
	"github.com/casper-ecosystem/casper-golang-sdk/serialization"
	"github.com/casper-ecosystem/casper-golang-sdk/types"
	"math/big"
	"strconv"
	"strings"
	"time"
	contract "ultimatedivision/pkg/contractcasper"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/pkg/pagination"
	"ultimatedivision/users"
)

// ErrMarketplace indicated that there was an error in service.
var ErrMarketplace = errs.Class("marketplace service error")

// Service is handling marketplace related logic.
//
// architecture: Service.
type Service struct {
	config      Config
	marketplace DB
	users       *users.Service
	cards       *cards.Service
	nfts        *nfts.Service
	contract    *contract.Casper
}

// NewService is a constructor for marketplace service.
func NewService(config Config, marketplace DB, users *users.Service, cards *cards.Service, nfts *nfts.Service) *Service {
	return &Service{
		config:      config,
		marketplace: marketplace,
		users:       users,
		cards:       cards,
		nfts:        nfts,
	}
}

// CreateLot add lot in DB.
func (service *Service) CreateLot(ctx context.Context, createLot CreateLot) error {
	// TODO: add transaction.
	card, err := service.cards.Get(ctx, createLot.CardID)
	if err == nil {
		if card.UserID != createLot.UserID {
			return ErrMarketplace.New("it is not the user's card")
		}

		if card.Status == cards.StatusSale {
			return ErrMarketplace.New("the card is already on sale")
		}

		if err := service.cards.UpdateStatus(ctx, createLot.CardID, cards.StatusSale); err != nil {
			return ErrMarketplace.Wrap(err)
		}

		createLot.Type = TypeCard
	}
	// TODO: check other items.

	if createLot.Type == "" {
		return ErrMarketplace.New("not found item by id")
	}

	if _, err := service.users.Get(ctx, createLot.UserID); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	if createLot.MaxPrice.BitLen() != 0 && createLot.MaxPrice.Cmp(&createLot.StartPrice) == -1 {
		return ErrMarketplace.New("max price less start price")
	}

	if createLot.Period < MinPeriod && createLot.Period < MaxPeriod {
		return ErrMarketplace.New("period exceed the range from 1 to 120 hours")
	}

	lot := Lot{
		CardID:     card.ID,
		Type:       createLot.Type,
		UserID:     createLot.UserID,
		Status:     StatusActive,
		StartPrice: createLot.StartPrice,
		MaxPrice:   createLot.MaxPrice,
		StartTime:  time.Now().UTC(),
		EndTime:    time.Now().UTC().Add(time.Duration(createLot.Period) * time.Hour),
		Period:     createLot.Period,
	}

	return ErrMarketplace.Wrap(service.marketplace.CreateLot(ctx, lot))
}

// GetLotByID returns lot by id from DB.
func (service *Service) GetLotByID(ctx context.Context, id uuid.UUID) (Lot, error) {
	lot, err := service.marketplace.GetLotByID(ctx, id)
	return lot, ErrMarketplace.Wrap(err)
}

// GetNFTByCardID returns nft by card id from DB.
func (service *Service) GetNFTByCardID(ctx context.Context, id uuid.UUID) (nfts.NFT, error) {
	nft, err := service.nfts.GetNFTByCardID(ctx, id)
	return nft, ErrMarketplace.Wrap(err)
}

// ListActiveLots returns active lots from DB.
func (service *Service) ListActiveLots(ctx context.Context, cursor pagination.Cursor) (Page, error) {
	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}
	lotsPage, err := service.marketplace.ListActiveLots(ctx, cursor)
	return lotsPage, ErrMarketplace.Wrap(err)
}

// ListExpiredLots returns all expired lots form the database.
func (service *Service) ListExpiredLots(ctx context.Context) ([]Lot, error) {
	lots, err := service.marketplace.ListExpiredLot(ctx)
	return lots, ErrMarketplace.Wrap(err)
}

// ListActiveLotsWithFilters returns active lots from DB, taking the necessary filters.
func (service *Service) ListActiveLotsWithFilters(ctx context.Context, filters []cards.Filters, cursor pagination.Cursor) (Page, error) {
	var lotsPage Page
	for _, v := range filters {
		err := v.Validate()
		if err != nil {
			return lotsPage, ErrMarketplace.Wrap(err)
		}
	}

	cardIDs, err := service.cards.ListCardIDsWithFiltersWhereActiveLot(ctx, filters)
	if err != nil {
		return lotsPage, ErrMarketplace.Wrap(err)
	}

	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}
	lotsPage, err = service.marketplace.ListActiveLotsByCardID(ctx, cardIDs, cursor)
	return lotsPage, ErrMarketplace.Wrap(err)
}

// ListActiveLotsByPlayerName returns active lots from DB by player name card.
func (service *Service) ListActiveLotsByPlayerName(ctx context.Context, filter cards.Filters, cursor pagination.Cursor) (Page, error) {
	var lotsPage Page
	strings.ToValidUTF8(filter.Value, "")

	// TODO: add best check.
	_, err := strconv.Atoi(filter.Value)
	if err == nil {
		return lotsPage, ErrMarketplace.Wrap(cards.ErrInvalidFilter.New("%s %s", filter.Value, err))
	}

	cardIDs, err := service.cards.ListCardIDsByPlayerNameWhereActiveLot(ctx, filter)
	if err != nil {
		return lotsPage, ErrMarketplace.Wrap(err)
	}

	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}
	lotsPage, err = service.marketplace.ListActiveLotsByCardID(ctx, cardIDs, cursor)
	return lotsPage, ErrMarketplace.Wrap(err)
}

// ListExpiredLot returns not active lots from DB.
func (service *Service) ListExpiredLot(ctx context.Context) ([]Lot, error) {
	lots, err := service.marketplace.ListExpiredLot(ctx)
	return lots, ErrMarketplace.Wrap(err)
}

// PlaceBetLot checks the amount of money and makes a bet.
func (service *Service) PlaceBetLot(ctx context.Context, betLot BetLot) error {
	if _, err := service.users.Get(ctx, betLot.UserID); err != nil {
		return ErrMarketplace.Wrap(err)
	}
	// TODO: check if the user has the required amount of money.

	lot, err := service.GetLotByID(ctx, betLot.CardID)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}
	if lot.Status == StatusSold || lot.Status == StatusSoldBuynow {
		return ErrMarketplace.New("the lot is already on sale")
	}
	if lot.Status == StatusExpired {
		return ErrMarketplace.New("the lot is already on expired")
	}

	if betLot.BetAmount.Cmp(&lot.StartPrice) == -1 || betLot.BetAmount.Cmp(&lot.CurrentPrice) == -1 || betLot.BetAmount.Cmp(&lot.CurrentPrice) == 0 {
		return ErrMarketplace.New("not enough money")
	}

	/** TODO: the transaction may be required for all operations,
	  so that an error in the middle does not lead to an unwanted result in the database. **/

	// TODO: update status to `hold` for new user's money.
	// TODO: unhold old user's money if exist.

	if err := service.UpdateShopperIDLot(ctx, betLot.CardID, betLot.UserID); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	if (betLot.BetAmount.Cmp(&lot.MaxPrice) == 1 || betLot.BetAmount.Cmp(&lot.MaxPrice) == 0) && lot.MaxPrice.BitLen() != 0 {
		if err = service.UpdateCurrentPriceLot(ctx, betLot.CardID, lot.MaxPrice); err != nil {
			return ErrMarketplace.Wrap(err)
		}

		winLot := WinLot{
			CardID:    betLot.CardID,
			Type:      TypeCard,
			UserID:    lot.UserID,
			ShopperID: betLot.UserID,
			Status:    StatusSoldBuynow,
			Amount:    lot.MaxPrice,
		}

		if err = service.WinLot(ctx, winLot); err != nil {
			return ErrMarketplace.Wrap(err)
		}

	} else {
		if err = service.UpdateCurrentPriceLot(ctx, betLot.CardID, betLot.BetAmount); err != nil {
			return ErrMarketplace.Wrap(err)
		}
		if lot.EndTime.Sub(time.Now().UTC()) < time.Minute {
			if err = service.UpdateEndTimeLot(ctx, betLot.CardID, time.Now().UTC().Add(time.Minute)); err != nil {
				return ErrMarketplace.Wrap(err)
			}
		}
	}

	return nil
}

// WinLot changes owner of the item and transfers money.
func (service *Service) WinLot(ctx context.Context, winLot WinLot) error {
	if err := service.UpdateStatusLot(ctx, winLot.CardID, winLot.Status); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	// TODO: transfer money to the old cardholder from new user. If userID == shopperID not transfer mb.

	if winLot.Type == TypeCard {
		if err := service.cards.UpdateStatus(ctx, winLot.CardID, cards.StatusActive); err != nil {
			return ErrMarketplace.Wrap(err)
		}

		if winLot.UserID != winLot.ShopperID {
			if err := service.cards.UpdateUserID(ctx, winLot.CardID, winLot.ShopperID); err != nil {
				return ErrMarketplace.Wrap(err)
			}
		}
	}
	// TODO: check other items.

	return nil
}

// UpdateShopperIDLot updates shopper id of lot.
func (service *Service) UpdateShopperIDLot(ctx context.Context, id, shopperID uuid.UUID) error {
	return ErrMarketplace.Wrap(service.marketplace.UpdateShopperIDLot(ctx, id, shopperID))
}

// UpdateStatusLot updates status of lot.
func (service *Service) UpdateStatusLot(ctx context.Context, id uuid.UUID, status Status) error {
	return ErrMarketplace.Wrap(service.marketplace.UpdateStatusLot(ctx, id, status))
}

// UpdateCurrentPriceLot updates current price of lot.
func (service *Service) UpdateCurrentPriceLot(ctx context.Context, id uuid.UUID, currentPrice big.Int) error {
	return ErrMarketplace.Wrap(service.marketplace.UpdateCurrentPriceLot(ctx, id, currentPrice))
}

// UpdateEndTimeLot updates end time of lot.
func (service *Service) UpdateEndTimeLot(ctx context.Context, id uuid.UUID, endTime time.Time) error {
	return ErrMarketplace.Wrap(service.marketplace.UpdateEndTimeLot(ctx, id, endTime))
}

// Delete deletes lot in the database.
func (service *Service) Delete(ctx context.Context, cardID uuid.UUID) error {
	return ErrMarketplace.Wrap(service.marketplace.Delete(ctx, cardID))
}

// PublicKey returns public key for specific network.
func PublicKey(ctx context.Context, networkId string) ([]byte, error) {
}

// BridgeOut initiates outbound bridge transaction.
func (service *Service) BridgeOut(ctx context.Context, req TokenOutRequest) ([]byte, error) {
	respPubKey, err := PublicKey(ctx, TypeCasper)
	if err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}

	publicKey := keypair.PublicKey{
		Tag:        keypair.KeyTagEd25519,
		PubKeyData: respPubKey,
	}

	standardPayment := new(big.Int).SetUint64(service.config.GasLimit)

	deployParams := sdk.NewDeployParams(publicKey, strings.ToLower(service.config.ChainName.String()), nil, 0)
	payment := sdk.StandardPayment(standardPayment)

	// token contract.
	tokenContractFixedBytes := types.FixedByteArray(req.Token)
	tokenContract := types.CLValue{
		Type:      types.CLTypeByteArray,
		ByteArray: &tokenContractFixedBytes,
	}
	tokenContractBytes, err := serialization.Marshal(tokenContract)
	if err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}

	// amount.
	amount := types.CLValue{
		Type: types.CLTypeU256,
		U256: req.Amount,
	}
	amountBytes, err := serialization.Marshal(amount)
	if err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}

	// transaction_id.
	transactionID := types.CLValue{
		Type: types.CLTypeU256,
		U256: req.TransactionID,
	}
	transactionIDBytes, err := serialization.Marshal(transactionID)
	if err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}

	//  source chain.
	sourceChain := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.From.NetworkName,
	}
	sourceChainBytes, err := serialization.Marshal(sourceChain)
	if err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}

	// source address.
	sourceAddress := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.From.Address,
	}
	sourceAddressBytes, err := serialization.Marshal(sourceAddress)
	if err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}

	// recipient.
	var recipientHashBytes [32]byte
	copy(recipientHashBytes[:], req.To)

	recipient := types.CLValue{
		Type: types.CLTypeKey,
		Key: &types.Key{
			Type:    types.KeyTypeAccount,
			Account: recipientHashBytes,
		},
	}
	recipientBytes, err := serialization.Marshal(recipient)
	if err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}

	args := map[string]sdk.Value{
		"token_contract": {
			IsOptional:  false,
			Tag:         types.CLTypeByteArray,
			StringBytes: hex.EncodeToString(tokenContractBytes),
		},
		"amount": {
			Tag:         types.CLTypeU256,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(amountBytes),
		},
		"transaction_id": {
			Tag:         types.CLTypeU256,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(transactionIDBytes),
		},
		"source_chain": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(sourceChainBytes),
		},
		"source_address": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(sourceAddressBytes),
		},
		"recipient": {
			Tag:         types.CLTypeKey,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(recipientBytes),
		},
	}

	keyOrder := []string{
		"token_contract",
		"amount",
		"transaction_id",
		"source_chain",
		"source_address",
		"recipient",
	}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(service.config.BridgeContractAddress)
	if err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "bridge_out", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	reqSign := SignRequest{
		NetworkId: TypeCasper,
		Data:      deploy.Hash,
	}
	signature, err := service.bridge.Sign(ctx, reqSign)
	if err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signature,
	}

	approval := sdk.Approval{
		Signer:    publicKey,
		Signature: signatureKeypair,
	}

	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := service.contract.PutDeploy(*deploy)
	if err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}

	txHash, err := hex.DecodeString(hash)

	return txHash, ErrMarketplace.Wrap(err)
}
