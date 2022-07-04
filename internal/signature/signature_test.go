// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package signature_test

import (
	"math/big"
	"testing"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision/internal/signature"
)

func TestSignature(t *testing.T) {
	privateKey := "5aefce0a2d473f59578fa7dee6a122d6509af1e0f79fcbee700dfcfeddabe4cc"

	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	require.NoError(t, err)

	nftStoreSignature := signature.NFTStoreSignature{
		EncodedMethod:   "Signature(address,address,uint256,uint256)",
		WalletAddress:   common.HexToAddress("0x56f088767D91badc379155290c4205c7b917a36E"),
		ContractAddress: common.HexToAddress("0x510D43D563A792C5F13fF2C69fFcE1f0226056BD"),
		TokenID:         5,
		Value:           big.NewInt(100),
		PrivateKey:      privateKeyECDSA,
	}

	expectedSignature := evmsignature.Signature("6833702e092bde75c651b8d9d846b0013e9fcdd7f4b5ead2bc8fcec674d8b980419b48bb9b67595fe5716f81bc8e3b73731bffcbac8587d8a62edb04120daa4f1b")

	t.Run("GenerateNFTStoreSignature", func(t *testing.T) {
		signature, err := signature.GenerateNFTStoreSignature(nftStoreSignature)
		require.NoError(t, err)

		assert.Equal(t, expectedSignature, signature)
	})
}
