// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package signer_test

import (
	"github.com/BoostyLabs/evmsignature"
	"github.com/stretchr/testify/assert"
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

	expectedSignature := evmsignature.Signature("55f0b9370fcdd1f45b83ffc384a351022aafbf2033dfec6899980905cf3a8e36269a3bf0c8c2a51e4976eed96cd587af7390c80eddd7d92880ecd837d30f2b091c")
	expectedCasperSignature := evmsignature.Signature("cc58b2fdbae61bac7ef6044c586dd74fbe4ff0a5289dca109c5fa8a7ed8ffd777719c4975e4f1309528cbe06b01eddc5f32b9479a4aaf208aab551eec3a8f1a51c")
	expectedNonceSignature := evmsignature.Signature("3ceb98f80fc5c8525f07c13e60a8cb54ebbd051546ae8cd1446353567494f5571e3c0620b1e28fceaac4d50d9d4acfb13de4858f5548c7cc5489d21354d0de3a1b")
	expectedCasperNonceSignature := evmsignature.Signature("090072221ac59344060ca70defd38339585ff4d6d54d9c70cf4591aaea3262575ce482b3c2e037f298a06c153c219ef0f36297c4475c4bba2f1730f4735ae4971b")

	value := *big.NewInt(100)

	t.Run("GenerateSignatureWithValue", func(t *testing.T) {
		signature, err := signer.GenerateSignatureWithValue(signer.Address(wallet), contractAddress, tokenID1, privateKeyECDSA)
		assert.Equal(t, expectedSignature, signature)
		require.NoError(t, err)
	})

	t.Run("GenerateCasperSignatureWithValue", func(t *testing.T) {
		signature, err := signer.GenerateSignatureWithValue(signer.Address(casperWallet), contractAddress, tokenID1, privateKeyECDSA)
		require.NoError(t, err)
		assert.Equal(t, expectedCasperSignature, signature)
	})

	t.Run("GenerateSignatureWithValueAndNonce", func(t *testing.T) {
		signature, err := signer.GenerateSignatureWithValueAndNonce(signer.Address(wallet), contractAddress, &value, tokenID1, privateKeyECDSA)
		require.NoError(t, err)
		assert.Equal(t, expectedNonceSignature, signature)
	})

	t.Run("GenerateCasperSignatureWithValueAndNonce", func(t *testing.T) {
		signature, err := signer.GenerateSignatureWithValueAndNonce(signer.Address(casperWallet), contractAddress, &value, tokenID1, privateKeyECDSA)
		require.NoError(t, err)
		assert.Equal(t, expectedCasperNonceSignature, signature)
	})

}
