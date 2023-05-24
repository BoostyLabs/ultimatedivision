// Copyright (C) 2021-2023 Creditor Corp. Group.
// See LICENSE for copying information.

package casper_test

import (
	"context"
	"encoding/hex"
	"os"
	"testing"

	casper_ed25519 "github.com/casper-ecosystem/casper-golang-sdk/keypair/ed25519"
	"github.com/stretchr/testify/require"

	"ultimatedivision/internal/contract/casper"
	contract "ultimatedivision/pkg/contractcasper"
)

func TestCasper_SetSigner(t *testing.T) {
	//t.Skip("for manual testing")

	var (
		casperNodeAddress = "http://136.243.187.84:7777/rpc"

		privateKeyEd25519ForTransaction = "1228fcc08c02bfe100543a2581f60b0ad0e09f4c53e81641f09a70850880a256c1a19239600bf8293462d99f2ec19d5d1b443c760f9bdb4d720554585b26139a"
		pathToPublicKeyFile             = "./public_key.in"

		bridgeContractPackageHash = "9299f58df67c2eff01e97f362996d35ab5393167e58c58360b1721cce95a7bbc"
	)

	ctx := context.Background()
	casperClient := contract.New(casperNodeAddress)

	privateKeyForTransferSigningBytes, err := hex.DecodeString(privateKeyEd25519ForTransaction)
	require.NoError(t, err)

	publicKey := make([]byte, 32)
	copy(publicKey, privateKeyForTransferSigningBytes[32:])
	pair := casper_ed25519.ParseKeyPair(publicKey, privateKeyForTransferSigningBytes[:32])

	transfer := casper.NewTransfer(casperClient, func(b []byte) ([]byte, error) {
		casperSignature := pair.Sign(b)
		return casperSignature.SignatureData, nil
	})

	publicKeyFile, err := os.ReadFile(pathToPublicKeyFile)
	require.NoError(t, err)

	txHash, err := transfer.SetSigner(ctx, casper.SetSignerRequest{
		PublicKey:                   pair.PublicKey(),
		ChainName:                   "CASPER-TEST",
		StandardPaymentForBridgeOut: 2500000000, // 2.5 CSPR.
		BridgeContractPackageHash:   bridgeContractPackageHash,
		Value:                       string(publicKeyFile),
	})
	require.NoError(t, err)
	require.NotEmpty(t, txHash)
}

func TestCasper_SetAcceptOffer(t *testing.T) {

	var (
		//methodName = "accept_offer"
		//

		//
		//contractHash = "hash-4ff1e5e37b8720e8049bfff88676d8e27c1037c02e1172a1006c6d2a535607da"
		//contractAddress = "26fa17f8bfb593357a0d948744ea502bb6544400f827d81639cfccbb319fe1b2"
		//nft_contract_hash = "hash-805b58313894dbef4ab184267484580fae7758d9591d6bc8da9b283cb4083cb8"
		//token_id          = ""
		//erc20_contract    = ""

		casperNodeAddress = "http://136.243.187.84:7777/rpc"

		privateKeyEd25519ForTransaction = "1228fcc08c02bfe100543a2581f60b0ad0e09f4c53e81641f09a70850880a256c1a19239600bf8293462d99f2ec19d5d1b443c760f9bdb4d720554585b26139a"
		pathToPublicKeyFile             = "./public_key.in"

		//bridgeContractPackageHash = "9299f58df67c2eff01e97f362996d35ab5393167e58c58360b1721cce95a7bbc"

		// DragonNFT_contract_hash.
		//bridgeContractPackageHash = "805b58313894dbef4ab184267484580fae7758d9591d6bc8da9b283cb4083cb8"
		// market_contract_hash.
		bridgeContractPackageHash = "3f334d12c60f09ce5c1e1f29155de8327d26d6122b6648f6d06cc846c1b65088"
	)

	ctx := context.Background()
	casperClient := contract.New(casperNodeAddress)

	privateKeyForTransferSigningBytes, err := hex.DecodeString(privateKeyEd25519ForTransaction)
	require.NoError(t, err)

	publicKey := make([]byte, 32)
	copy(publicKey, privateKeyForTransferSigningBytes[32:])
	pair := casper_ed25519.ParseKeyPair(publicKey, privateKeyForTransferSigningBytes[:32])

	transfer := casper.NewTransfer(casperClient, func(b []byte) ([]byte, error) {
		casperSignature := pair.Sign(b)
		return casperSignature.SignatureData, nil
	})

	publicKeyFile, err := os.ReadFile(pathToPublicKeyFile)
	require.NoError(t, err)

	txHash, err := transfer.SetAcceptOffer(ctx, casper.SetSignerRequest{
		PublicKey:                   pair.PublicKey(),
		ChainName:                   "casper-test",
		StandardPaymentForBridgeOut: 2500000000, // 2.5 CSPR.
		BridgeContractPackageHash:   bridgeContractPackageHash,
		Value:                       string(publicKeyFile),
	})
	require.NoError(t, err)
	require.NotEmpty(t, txHash)
}
