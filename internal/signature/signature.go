// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package signature

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"
)

// ErrCreateSignature indicates that an error occurred while creating a signature.
var ErrCreateSignature = errs.Class("signature package error")

type NFTStoreSignature struct {
	EncodedMethod   string
	WalletAddress   common.Address
	ContractAddress common.Address
	TokenID         int64
	Value           *big.Int
	PrivateKey      *ecdsa.PrivateKey
}

// GenerateNFTStoreSignature generates signature for user's wallet to buy nft in store.
func GenerateNFTStoreSignature(nftStoreSignature NFTStoreSignature) (evmsignature.Signature, error) {
	var values [][]byte

	encodedMethodSelector := crypto.Keccak256Hash([]byte(nftStoreSignature.EncodedMethod)).Bytes()[:4]

	tokenIDStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", nftStoreSignature.TokenID))
	tokenIDByte, err := hex.DecodeString(string(tokenIDStringWithZeros))
	if err != nil {
		return "", errs.Wrap(err)
	}

	valueStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", nftStoreSignature.Value))
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", errs.Wrap(err)
	}

	values = append(values, encodedMethodSelector, nftStoreSignature.WalletAddress.Hash().Bytes(), nftStoreSignature.ContractAddress.Hash().Bytes(),
		tokenIDByte, valueByte)

	createSignature := evmsignature.CreateSignature{
		Values:     values,
		PrivateKey: nftStoreSignature.PrivateKey,
	}

	signatureByte, err := evmsignature.MakeSignature(createSignature)
	if err != nil {
		return "", errs.Wrap(err)
	}

	signature, err := evmsignature.ReformSignature(signatureByte)

	return signature, errs.Wrap(err)
}
