// Copyright (C) 2021-2023 Creditor Corp. Group.
// See LICENSE for copying information.

package casper

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/casper-ecosystem/casper-golang-sdk/keypair"
	"github.com/casper-ecosystem/casper-golang-sdk/sdk"
	"github.com/casper-ecosystem/casper-golang-sdk/serialization"
	"github.com/casper-ecosystem/casper-golang-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	contract "ultimatedivision/pkg/contractcasper"
)

// SetSignerRequest describes values to calls setSigner method.
type SetSignerRequest struct {
	PublicKey                   keypair.PublicKey
	ChainName                   string
	StandardPaymentForBridgeOut int64
	BridgeContractPackageHash   string
	Value                       string
}

// SetFinalListingRequest describes values to calls final listing method.
type SetFinalListingRequest struct {
	PublicKey                   keypair.PublicKey
	ChainName                   string
	NftContractHash             string
	TokenID                     string
	StandardPaymentForBridgeOut int64
	BridgeContractPackageHash   string
}

// Transfer describes sign func to sign transaction and casper client to send transaction.
type Transfer struct {
	casper contract.Casper

	sign func([]byte) ([]byte, error)
}

// NewTransfer is constructor for Transfer.
func NewTransfer(casper contract.Casper, sign func([]byte) ([]byte, error)) *Transfer {
	return &Transfer{
		casper: casper,
		sign:   sign,
	}
}

// SetFinalListing send NFT to winner.
func (t *Transfer) SetFinalListing(ctx context.Context, req SetFinalListingRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	// ціна на газ ліміт.
	payment := sdk.StandardPayment(big.NewInt(req.StandardPaymentForBridgeOut))
	fmt.Println("payment ------->", payment)
	// NFT contract hash, where we mint NFT.
	nftContract := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.NftContractHash,
	}
	nftContractBytes, err := serialization.Marshal(nftContract)
	if err != nil {
		return "", err
	}

	tokenIDFromStringByteSlice := []byte("$\u0000\u0000\u0000" + req.TokenID)
	fmt.Println("tokenIDFromStringByteSlice ------->", tokenIDFromStringByteSlice)
	args := map[string]sdk.Value{
		"nft_contract_hash": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(nftContractBytes),
		},
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDFromStringByteSlice),
		},
	}
	fmt.Println("args ------->", args)
	keyOrder := []string{"nft_contract_hash", "token_id"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	// contract that we are using.
	contractHexBytes, err := hex.DecodeString(req.BridgeContractPackageHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "final_listing", *runtimeArgs)
	fmt.Println("session ------->", payment)
	deploy := sdk.MakeDeploy(deployParams, payment, session)
	fmt.Println("deploy ------->", payment)
	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)
	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}
	fmt.Println("hash ------->", hash)

	return hash, nil
}

// SetAcceptOffer send NFT to winner.
func (t *Transfer) SetAcceptOffer(ctx context.Context, req SetSignerRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPaymentForBridgeOut))

	value := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.Value,
	}
	valueBytes, err := serialization.Marshal(value)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"signer": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(valueBytes),
		},
	}

	//amountBytes, err := serialization.Marshal(nftContractHash)
	//if err != nil {
	//	return "", err
	//}
	//
	//args2 := map[string]sdk.Value{
	//	"nft_contract_hash": {
	//		Tag:         types.CLTypeString,
	//		IsOptional:  false,
	//		StringBytes: hex.EncodeToString(amountBytes),
	//	},
	//	"token_id": {
	//		Tag:         types.CLTypeString,
	//		IsOptional:  false,
	//		StringBytes: hex.EncodeToString(gasCommissionBytes),
	//	},
	//	"erc20_contract": {
	//		Tag:         types.CLTypeU256,
	//		IsOptional:  false,
	//		StringBytes: hex.EncodeToString(deadlineBytes),
	//	},
	//}

	keyOrder := []string{"nft_contract_hash", "token_id", "erc20_contract"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.BridgeContractPackageHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "accept_offer", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

////
////
////
////
////
////
////
////

type MintOneRequest struct {
	PublicKey              keypair.PublicKey
	ChainName              string
	StandardPayment        int64
	NFTContractPackageHash string
	TokenID                string
	Recipient              string
}

func (t *Transfer) MintOne(ctx context.Context, req MintOneRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	tokenID := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.TokenID,
	}
	tokenIDBytes, err := serialization.Marshal(tokenID)
	if err != nil {
		return "", err
	}

	recipient := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.Recipient,
	}
	recipientBytes, err := serialization.Marshal(recipient)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDBytes),
		},
		"recipient": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(recipientBytes),
		},
	}

	keyOrder := []string{"token_id", "recipient"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.NFTContractPackageHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "mint_one", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

type ApproveNFTRequest struct {
	PublicKey              keypair.PublicKey
	ChainName              string
	StandardPayment        int64
	NFTContractPackageHash string
	Spender                string
	TokenID                string
}

func (t *Transfer) ApproveNFT(ctx context.Context, req ApproveNFTRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	spender := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.Spender,
	}
	spenderBytes, err := serialization.Marshal(spender)
	if err != nil {
		return "", err
	}

	tokenID := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.TokenID,
	}
	tokenIDBytes, err := serialization.Marshal(tokenID)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"spender": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(spenderBytes),
		},
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDBytes),
		},
	}

	keyOrder := []string{"spender", "token_id"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.NFTContractPackageHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "approve", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

type CreateListingRequest struct {
	PublicKey          keypair.PublicKey
	ChainName          string
	StandardPayment    int64
	MarketContractHash string
	NFTContractHash    string
	TokenID            string
	MinBidPrice        *big.Int
	RedemptionPrice    *big.Int
	AuctionDuration    *big.Int
}

func (t *Transfer) CreateListing(ctx context.Context, req CreateListingRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	nftContractHash := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.NFTContractHash,
	}
	nftContractHashBytes, err := serialization.Marshal(nftContractHash)
	if err != nil {
		return "", err
	}

	tokenID := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.TokenID,
	}
	tokenIDBytes, err := serialization.Marshal(tokenID)
	if err != nil {
		return "", err
	}

	minBidPrice := types.CLValue{
		Type: types.CLTypeU256,
		U256: req.MinBidPrice,
	}
	minBidPriceBytes, err := serialization.Marshal(minBidPrice)
	if err != nil {
		return "", err
	}

	redemptionPrice := types.CLValue{
		Type: types.CLTypeU256,
		U256: req.RedemptionPrice,
	}
	redemptionPriceBytes, err := serialization.Marshal(redemptionPrice)
	if err != nil {
		return "", err
	}

	auctionDuration := types.CLValue{
		Type: types.CLTypeU256,
		U256: req.AuctionDuration,
	}
	auctionDurationBytes, err := serialization.Marshal(auctionDuration)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"nft_contract_hash": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(nftContractHashBytes),
		},
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDBytes),
		},
		"min_bid_price": {
			Tag:         types.CLTypeU256,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(minBidPriceBytes),
		},
		"redemption_price": {
			Tag:         types.CLTypeU256,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(redemptionPriceBytes),
		},
		"auction_duration": {
			Tag:         types.CLTypeU256,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(auctionDurationBytes),
		},
	}

	keyOrder := []string{"nft_contract_hash", "token_id", "min_bid_price", "redemption_price", "auction_duration"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.MarketContractHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "create_listing", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

type ApproveTokensRequest struct {
	PublicKey          keypair.PublicKey
	ChainName          string
	StandardPayment    int64
	TokensContractHash string
	Spender            common.Hash
	Amount             *big.Int
}

func (t *Transfer) ApproveTokens(ctx context.Context, req ApproveTokensRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	var spenderHashBytes [32]byte
	copy(spenderHashBytes[:], req.Spender.Bytes())

	spender := types.CLValue{
		Type: types.CLTypeKey,
		Key: &types.Key{
			Type: types.KeyTypeHash,
			Hash: spenderHashBytes,
		},
	}
	spenderBytes, err := serialization.Marshal(spender)
	if err != nil {
		return "", err
	}

	amount := types.CLValue{
		Type: types.CLTypeU256,
		U256: req.Amount,
	}
	amountBytes, err := serialization.Marshal(amount)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"spender": {
			Tag:         types.CLTypeKey,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(spenderBytes),
		},
		"amount": {
			Tag:         types.CLTypeU256,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(amountBytes),
		},
	}

	keyOrder := []string{"spender", "amount"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.TokensContractHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "approve", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

type MakeOfferRequest struct {
	PublicKey          keypair.PublicKey
	ChainName          string
	StandardPayment    int64
	MarketContractHash string
	NFTContractHash    string
	TokenID            string
	Amount             *big.Int
}

func (t *Transfer) MakeOffer(ctx context.Context, req MakeOfferRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	nftContractHash := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.NFTContractHash,
	}
	nftContractHashBytes, err := serialization.Marshal(nftContractHash)
	if err != nil {
		return "", err
	}

	tokenID := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.TokenID,
	}
	tokenIDBytes, err := serialization.Marshal(tokenID)
	if err != nil {
		return "", err
	}

	amount := types.CLValue{
		Type: types.CLTypeU256,
		U256: req.Amount,
	}
	amountBytes, err := serialization.Marshal(amount)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"nft_contract_hash": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(nftContractHashBytes),
		},
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDBytes),
		},
		"offer_price": {
			Tag:         types.CLTypeU256,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(amountBytes),
		},
	}

	keyOrder := []string{"nft_contract_hash", "token_id", "offer_price"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.MarketContractHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "make_offer", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

type AcceptOfferRequest struct {
	PublicKey          keypair.PublicKey
	ChainName          string
	StandardPayment    int64
	MarketContractHash string
	NFTContractHash    string
	TokenID            string
}

func (t *Transfer) AcceptOffer(ctx context.Context, req AcceptOfferRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	nftContractHash := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.NFTContractHash,
	}
	nftContractHashBytes, err := serialization.Marshal(nftContractHash)
	if err != nil {
		return "", err
	}

	tokenID := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.TokenID,
	}
	tokenIDBytes, err := serialization.Marshal(tokenID)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"nft_contract_hash": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(nftContractHashBytes),
		},
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDBytes),
		},
	}

	keyOrder := []string{"nft_contract_hash", "token_id"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.MarketContractHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "accept_offer", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

type BuyListingRequest struct {
	PublicKey          keypair.PublicKey
	ChainName          string
	StandardPayment    int64
	MarketContractHash string
	NFTContractHash    string
	TokenID            string
}

func (t *Transfer) BuyListing(ctx context.Context, req BuyListingRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	nftContractHash := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.NFTContractHash,
	}
	nftContractHashBytes, err := serialization.Marshal(nftContractHash)
	if err != nil {
		return "", err
	}

	tokenID := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.TokenID,
	}
	tokenIDBytes, err := serialization.Marshal(tokenID)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"nft_contract_hash": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(nftContractHashBytes),
		},
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDBytes),
		},
	}

	keyOrder := []string{"nft_contract_hash", "token_id"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.MarketContractHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "buy_listing", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}

type FinalListingRequest struct {
	PublicKey          keypair.PublicKey
	ChainName          string
	StandardPayment    int64
	MarketContractHash string
	NFTContractHash    string
	TokenID            string
}

func (t *Transfer) FinalListing(ctx context.Context, req FinalListingRequest) (string, error) {
	deployParams := sdk.NewDeployParams(req.PublicKey, strings.ToLower(req.ChainName), nil, 0)
	payment := sdk.StandardPayment(big.NewInt(req.StandardPayment))

	nftContractHash := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.NFTContractHash,
	}
	nftContractHashBytes, err := serialization.Marshal(nftContractHash)
	if err != nil {
		return "", err
	}

	tokenID := types.CLValue{
		Type:   types.CLTypeString,
		String: &req.TokenID,
	}
	tokenIDBytes, err := serialization.Marshal(tokenID)
	if err != nil {
		return "", err
	}

	args := map[string]sdk.Value{
		"nft_contract_hash": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(nftContractHashBytes),
		},
		"token_id": {
			Tag:         types.CLTypeString,
			IsOptional:  false,
			StringBytes: hex.EncodeToString(tokenIDBytes),
		},
	}

	keyOrder := []string{"nft_contract_hash", "token_id"}
	runtimeArgs := sdk.NewRunTimeArgs(args, keyOrder)

	contractHexBytes, err := hex.DecodeString(req.MarketContractHash)
	if err != nil {
		return "", err
	}

	var contractHashBytes [32]byte
	copy(contractHashBytes[:], contractHexBytes)
	session := sdk.NewStoredContractByHash(contractHashBytes, "final_listing", *runtimeArgs)

	deploy := sdk.MakeDeploy(deployParams, payment, session)

	signedTx, err := t.sign(deploy.Hash)
	if err != nil {
		return "", err
	}

	signatureKeypair := keypair.Signature{
		Tag:           keypair.KeyTagEd25519,
		SignatureData: signedTx,
	}
	approval := sdk.Approval{
		Signer:    req.PublicKey,
		Signature: signatureKeypair,
	}
	deploy.Approvals = append(deploy.Approvals, approval)

	hash, err := t.casper.PutDeploy(*deploy)
	if err != nil {
		return "", err
	}

	return hash, nil
}
