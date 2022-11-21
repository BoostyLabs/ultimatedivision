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

	tokenID1 := uuid.New()

	contractAddress := signer.Address("0x510D43D563A792C5F13fF2C69fFcE1f0226056BD")

	wallet := "0x56f088767D91badc379155290c4205c7b917a36E"

	value := big.Int{}

	t.Run("GenerateSignatureWithValue", func(t *testing.T) {
		_, err := signer.GenerateSignatureWithValue(signer.Address(wallet), contractAddress, tokenID1, privateKeyECDSA)
		require.NoError(t, err)
	})

	t.Run("GenerateSignatureWithValueAndNonce", func(t *testing.T) {
		_, err := signer.GenerateSignatureWithValueAndNonce(signer.Address(wallet), contractAddress, &value, tokenID1, privateKeyECDSA)
		require.NoError(t, err)
	})

}
