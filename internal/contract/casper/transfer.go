// Copyright (C) 2021-2023 Creditor Corp. Group.
// See LICENSE for copying information.

package casper

import (
	"context"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/casper-ecosystem/casper-golang-sdk/keypair"
	"github.com/casper-ecosystem/casper-golang-sdk/sdk"
	"github.com/casper-ecosystem/casper-golang-sdk/serialization"
	"github.com/casper-ecosystem/casper-golang-sdk/types"

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
