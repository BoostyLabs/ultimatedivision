// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package signer_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"ultimatedivision/pkg/signer"
)

func TestSignature(t *testing.T) {
	privateKey := "5aefce0a2d473f59578fa7dee6a122d6509af1e0f79fcbee700dfcfeddabe4cc"

	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	require.NoError(t, err)

	tokenID1 := uuid.MustParse("94b94d50-d001-4f88-b7cf-1763b39044b1")

	contractAddress := signer.Address("0x7ebd9ab31b8ed6678812ce3485cd41b703a210228973a0d510e8624721b3af61")

	wallet := "0x56f088767D91badc379155290c4205c7b917a36E"
	casperWallet := "0x9060c0820b5156b1620c8e3344d17f9fad5108f5dc2672f2308439e84363c88e"

	value := big.Int{}

	t.Run("GenerateSignatureWithValue", func(t *testing.T) {
		_, err := signer.GenerateSignatureWithValue(signer.Address(wallet), contractAddress, tokenID1, privateKeyECDSA)
		require.NoError(t, err)
	})

	t.Run("GenerateCasperSignatureWithValue", func(t *testing.T) {
		_, err := signer.GenerateSignatureWithValue(signer.Address(casperWallet), contractAddress, tokenID1, privateKeyECDSA)
		require.NoError(t, err)
	})

	t.Run("GenerateSignatureWithValueAndNonce", func(t *testing.T) {
		_, err := signer.GenerateSignatureWithValueAndNonce(signer.Address(wallet), contractAddress, &value, tokenID1, privateKeyECDSA)
		require.NoError(t, err)
	})

	t.Run("GenerateCasperSignatureWithValueAndNonce", func(t *testing.T) {
		_, err := signer.GenerateSignatureWithValueAndNonce(signer.Address(casperWallet), contractAddress, &value, tokenID1, privateKeyECDSA)
		require.NoError(t, err)
	})

}
