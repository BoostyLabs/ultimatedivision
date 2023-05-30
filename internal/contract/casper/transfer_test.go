// Copyright (C) 2021-2023 Creditor Corp. Group.
// See LICENSE for copying information.

package casper_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	casper_ed25519 "github.com/casper-ecosystem/casper-golang-sdk/keypair/ed25519"
	"github.com/stretchr/testify/require"

	"ultimatedivision/internal/contract/casper"
	contract "ultimatedivision/pkg/contractcasper"
)

func TestCasper_SetAcceptOffer(t *testing.T) {

	var (
		casperNodeAddress = "http://136.243.187.84:7777/rpc"

		privateKeyEd25519ForTransaction = "1228fcc08c02bfe100543a2581f60b0ad0e09f4c53e81641f09a70850880a256c1a19239600bf8293462d99f2ec19d5d1b443c760f9bdb4d720554585b26139a"
		pathToPublicKeyFile             = "./public_key.in"

		//bridgeContractPackageHash = "9299f58df67c2eff01e97f362996d35ab5393167e58c58360b1721cce95a7bbc"

		// DragonNFT_contract_hash.
		//bridgeContractPackageHash = "805b58313894dbef4ab184267484580fae7758d9591d6bc8da9b283cb4083cb8"
		// market_contract_hash.
		bridgeContractPackageHash = "277df7a58cfdc33a55176576ae9e5ab6c26b3b959aa9b18d5700cabea0c50598"
	)

	ctx := context.Background()
	casperClient := contract.New(casperNodeAddress)

	privateKeyForTransferSigningBytes, err := hex.DecodeString(privateKeyEd25519ForTransaction)
	require.NoError(t, err)

	// -----------
	publicKey := make([]byte, 32)
	copy(publicKey, privateKeyForTransferSigningBytes[32:])
	pair := casper_ed25519.ParseKeyPair(publicKey, privateKeyForTransferSigningBytes[:32])

	transfer := casper.NewTransfer(casperClient, func(b []byte) ([]byte, error) {
		casperSignature := pair.Sign(b)
		return casperSignature.SignatureData, nil
	})

	publicKeyFile, err := os.ReadFile(pathToPublicKeyFile)
	require.NoError(t, err)
	// -----------

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

func TestReadPrivateKeyFromFile(t *testing.T) {
	//t.Skip("for manual testing")
	pr, err := casper_ed25519.ParsePrivateKeyFile("./UltimateMint4_secret_key.pem")
	require.NoError(t, err)
	fmt.Println(hex.EncodeToString(pr))
	t.Error()
}

func TestCasper_SetFinalListing(t *testing.T) {
	var (
		casperNodeAddress = "http://136.243.187.84:7777/rpc"

		//privateKeyEd25519ForTransaction = "1228fcc08c02bfe100543a2581f60b0ad0e09f4c53e81641f09a70850880a256c1a19239600bf8293462d99f2ec19d5d1b443c760f9bdb4d720554585b26139a"
		privateAccountKey = "861bf3e31b748978c0f86cf19eb33816d5dff7f10482f60bae0148ae52ec6826"
		publicAccountKey  = "ad794c8f3da55845a5422506b4bd01ed1ae4e57378a82216d25d7351854b563d"
		//publicAccountKey = "01ad794c8f3da55845a5422506b4bd01ed1ae4e57378a82216d25d7351854b563d"

		//privateKeyEd25519ForTransaction = "861bf3e31b748978c0f86cf19eb33816d5dff7f10482f60bae0148ae52ec6826ad794c8f3da55845a5422506b4bd01ed1ae4e57378a82216d25d7351854b563d"

		//pathToPrivat                    = "./UltimateMint4_secret_key.pem"
		privateKeyEd25519ForTransaction = publicAccountKey + privateAccountKey

		bridgeContractPackageHash = "277df7a58cfdc33a55176576ae9e5ab6c26b3b959aa9b18d5700cabea0c50598"
		nftContractHash           = "contract-805b58313894dbef4ab184267484580fae7758d9591d6bc8da9b283cb4083cb8"
		tokenID                   = "7a78a717-bf19-4e7b-902f-7d4f334cb4c6"
		//tokenID = "1"
	)

	//fmt.Println("deploy-->", deploy)

	ctx := context.Background()
	casperClient := contract.New(casperNodeAddress)

	// -----------
	privateKeyForTransferSigningBytes, err := hex.DecodeString(privateKeyEd25519ForTransaction)
	require.NoError(t, err)

	privateAccountKeyBytes, err := hex.DecodeString(privateAccountKey)
	require.NoError(t, err)
	publicAccountKeyBytes, err := hex.DecodeString(publicAccountKey)
	require.NoError(t, err)

	fmt.Println("privateKeyForTransferSigningBytes-----_!!!!!!", privateKeyForTransferSigningBytes[32:])

	pair := casper_ed25519.ParseKeyPair(publicAccountKeyBytes, privateAccountKeyBytes)

	transfer := casper.NewTransfer(casperClient, func(b []byte) ([]byte, error) {
		casperSignature := pair.Sign(b)
		return casperSignature.SignatureData, nil
	})
	// -----------
	txHash, err := transfer.SetFinalListing(ctx, casper.SetFinalListingRequest{
		NftContractHash:             nftContractHash,
		PublicKey:                   pair.PublicKey(),
		ChainName:                   "casper-test",
		StandardPaymentForBridgeOut: 2500000000, // 2.5 CSPR.
		BridgeContractPackageHash:   bridgeContractPackageHash,
		TokenID:                     tokenID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, txHash)
}

func TestCasper_SetFinalListing2(t *testing.T) {
	var (
		casperNodeAddress = "http://136.243.187.84:7777/rpc"

		privateAccountKey = "861bf3e31b748978c0f86cf19eb33816d5dff7f10482f60bae0148ae52ec6826"
		publicAccountKey  = "bbf5e5d84de2a446b11880bb5533e7ba85b3830171f3d457527fe60c49333f32"

		bridgeContractPackageHash = "277df7a58cfdc33a55176576ae9e5ab6c26b3b959aa9b18d5700cabea0c50598"
		nftContractHash           = "contract-805b58313894dbef4ab184267484580fae7758d9591d6bc8da9b283cb4083cb8"
		tokenID                   = "7a78a717-bf19-4e7b-902f-7d4f334cb4c6"
		//tokenID = "1"
	)

	ctx := context.Background()
	casperClient := contract.New(casperNodeAddress)

	// -----------
	privateAccountKeyBytes, err := hex.DecodeString(privateAccountKey)
	require.NoError(t, err)
	publicAccountKeyBytes, err := hex.DecodeString(publicAccountKey)
	require.NoError(t, err)

	pair := casper_ed25519.ParseKeyPair(publicAccountKeyBytes, privateAccountKeyBytes)
	//pair := casper_ed25519.ParseKeyPair(privateAccountKeyBytes, publicAccountKeyBytes)

	transfer := casper.NewTransfer(casperClient, func(b []byte) ([]byte, error) {
		casperSignature := pair.Sign(b)
		return casperSignature.SignatureData, nil
	})
	// -----------
	txHash, err := transfer.SetFinalListing(ctx, casper.SetFinalListingRequest{
		NftContractHash:             nftContractHash,
		PublicKey:                   pair.PublicKey(),
		ChainName:                   "casper-test",
		StandardPaymentForBridgeOut: 2500000000, // 2.5 CSPR.
		BridgeContractPackageHash:   bridgeContractPackageHash,
		TokenID:                     tokenID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, txHash)
}
