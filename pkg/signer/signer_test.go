// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package signer_test

import (
	"math/big"
	"testing"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision/pkg/signer"
)

func TestSignature(t *testing.T) {
	privateKey := "5aefce0a2d473f59578fa7dee6a122d6509af1e0f79fcbee700dfcfeddabe4cc"

	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	require.NoError(t, err)

	tokenID := uuid.MustParse("94b94d50-d001-4f88-b7cf-1763b39044b1")

	contractAddress := signer.Address("0xc3d2bdedf7f309e2908ecf90d0dfb44acf2a8077cf053a05779fb15bbbfdbfb9")

	wallet := "0x56f088767D91badc379155290c4205c7b917a36E"
	casperWallet := "0x9060c0820b5156b1620c8e3344d17f9fad5108f5dc2672f2308439e84363c88e"

	expectedSignature := evmsignature.Signature("5603fd54241708d2c717518f94c64c1d0faea8481638655f8b1575e0583771e52451efd861a258a9ce1b13340d90a964c213586026ac2ed00fdadfaed952d78c1c")
	expectedCasperSignature := evmsignature.Signature("ee64aa00a97ba58f17eb0aadd93b6b70e650b4b634080ba2dd9621e96e99ef83710f623d0338c3f24e9e0e4ed48e67d3b788e20c5e9feffb92623c03f88174d81b")
	expectedNonceSignature := evmsignature.Signature("db8876e5d43be963a84ee4c86257e479bab9759bbac3eb9bf5f5d85f30b21be400c425617eb8542eaf60db9259f736fd5983fd3f9136e53ea0fe490f1e3df9011c")
	expectedCasperNonceSignature := evmsignature.Signature("c9cac1d7d80fbc431b54fce157d7d863d52147db4f27270deb3c3bc41e1fc2482178571df8be156d656d46f9a81132173561eecd6b3313ae511312fa7e8c4fbb1b")

	value := *big.NewInt(100)
	nonce := int64(5)

	t.Run("GenerateSignatureWithValue", func(t *testing.T) {
		signature, err := signer.GenerateSignatureWithValue(signer.Address(wallet), contractAddress, tokenID, privateKeyECDSA)
		assert.Equal(t, expectedSignature, signature)
		require.NoError(t, err)
	})

	t.Run("GenerateCasperSignatureWithValue", func(t *testing.T) {
		signature, err := signer.GenerateSignatureWithValue(signer.Address(casperWallet), contractAddress, tokenID, privateKeyECDSA)
		require.NoError(t, err)
		assert.Equal(t, expectedCasperSignature, signature)
	})

	t.Run("GenerateSignatureWithValueAndNonce", func(t *testing.T) {
		signature, err := signer.GenerateSignatureWithValueAndNonce(signer.Address(wallet), contractAddress, &value, nonce, privateKeyECDSA)
		require.NoError(t, err)
		assert.Equal(t, expectedNonceSignature, signature)
	})

	t.Run("GenerateCasperSignatureWithValueAndNonce", func(t *testing.T) {
		signature, err := signer.GenerateSignatureWithValueAndNonce(signer.Address(casperWallet), contractAddress, &value, nonce, privateKeyECDSA)
		require.NoError(t, err)
		assert.Equal(t, expectedCasperNonceSignature, signature)
	})

	t.Run("Negative GenerateSignatureWithValueAndNonce", func(t *testing.T) {
		nonce = -5
		_, err := signer.GenerateSignatureWithValueAndNonce(signer.Address(casperWallet), contractAddress, &value, nonce, privateKeyECDSA)
		require.Error(t, err)

		assert.True(t, signer.ErrCreateSignature.Has(err))
	})

}
