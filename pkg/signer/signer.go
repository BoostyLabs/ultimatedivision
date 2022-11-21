// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

package signer

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrCreateSignature indicates that an error occurred while creating a signature.
var ErrCreateSignature = errs.Class("signature package error")

// Address defines address type.
type Address string

// Signature defines signature type.
type Signature string

type NFTStoreSignature struct {
	MethodName      string
	WalletAddress   common.Address
	ContractAddress common.Address
	TokenID         int64
	Value           *big.Int
	PrivateKey      *ecdsa.PrivateKey
}

// GenerateSignatureWithValue generates signature for user's wallet with value.
func GenerateSignatureWithValue(addressWallet Address, addressContract Address, value uuid.UUID, privateKey *ecdsa.PrivateKey) (evmsignature.Signature, error) {
	var values [][]byte
	if err := addressWallet.IsValidAddress(); err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}
	if err := addressContract.IsValidAddress(); err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	addressWalletByte, err := hex.DecodeString(string(addressWallet)[evmsignature.LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	addressContractByte, err := hex.DecodeString(string(addressContract)[evmsignature.LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	valueStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", value))
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	values = append(values, addressWalletByte, addressContractByte, valueByte)
	createSignature := evmsignature.CreateSignature{
		Values:     values,
		PrivateKey: privateKey,
	}

	signatureByte, err := evmsignature.MakeSignature(createSignature)
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	signature, err := evmsignature.ReformSignature(signatureByte)

	return signature, ErrCreateSignature.Wrap(err)
}

// GenerateSignatureWithValueAndNonce generates signature for user's wallet with value and nonce.
func GenerateSignatureWithValueAndNonce(addressWallet Address, addressContract Address, value *big.Int, nonce uuid.UUID, privateKey *ecdsa.PrivateKey) (evmsignature.Signature, error) {
	var values [][]byte
	if err := addressWallet.IsValidAddress(); err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}
	if err := addressContract.IsValidAddress(); err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	addressWalletByte, err := hex.DecodeString(string(addressWallet)[evmsignature.LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	addressContractByte, err := hex.DecodeString(string(addressContract)[evmsignature.LengthHexPrefix:])
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	valueStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", value))
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	nonceStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", nonce))
	nonceByte, err := hex.DecodeString(string(nonceStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	values = append(values, addressWalletByte, addressContractByte, valueByte, nonceByte)
	createSignature := evmsignature.CreateSignature{
		Values:     values,
		PrivateKey: privateKey,
	}

	signatureByte, err := evmsignature.MakeSignature(createSignature)
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	signature, err := evmsignature.ReformSignature(signatureByte)

	return signature, ErrCreateSignature.Wrap(err)
}

// IsValidAddress checks if the address is valid.
func (address Address) IsValidAddress() error {
	if !common.IsHexAddress(string(address)) {
		return ErrCreateSignature.New("")
	}
	return nil
}

// GenerateNFTStoreSignature generates signature for user's wallet to buy nft in store.
func GenerateNFTStoreSignature(nftStoreSignature NFTStoreSignature) (evmsignature.Signature, error) {
	var values [][]byte

	methodSelector := crypto.Keccak256Hash([]byte(nftStoreSignature.MethodName)).Bytes()[:4]

	tokenIDStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", nftStoreSignature.TokenID))
	tokenIDByte, err := hex.DecodeString(string(tokenIDStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	valueStringWithZeros := evmsignature.CreateHexStringFixedLength(fmt.Sprintf("%x", nftStoreSignature.Value))
	valueByte, err := hex.DecodeString(string(valueStringWithZeros))
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	values = append(values, methodSelector, nftStoreSignature.WalletAddress.Hash().Bytes(), nftStoreSignature.ContractAddress.Hash().Bytes(),
		tokenIDByte, valueByte)

	createSignature := evmsignature.CreateSignature{
		Values:     values,
		PrivateKey: nftStoreSignature.PrivateKey,
	}

	signatureByte, err := evmsignature.MakeSignature(createSignature)
	if err != nil {
		return "", ErrCreateSignature.Wrap(err)
	}

	signature, err := evmsignature.ReformSignature(signatureByte)

	return signature, ErrCreateSignature.Wrap(err)
}
