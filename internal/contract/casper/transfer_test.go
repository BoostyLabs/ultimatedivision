// Copyright (C) 2021-2023 Creditor Corp. Group.
// See LICENSE for copying information.

package casper_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"testing"

	casper_ed25519 "github.com/casper-ecosystem/casper-golang-sdk/keypair/ed25519"
	"github.com/ethereum/go-ethereum/common"
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
	t.Skip("for manual testing")
	pr, err := casper_ed25519.ParsePrivateKeyFile("./new.pem")
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

////
////
////
////
////
////
////
////

func TestCasper_NftContract(t *testing.T) {
	t.Skip("for manual testing")

	var (
		casperNodeAddress = "http://116.202.169.210:7777/rpc"

		privateAccountKey = "4f0ffa6c3925d02127a9e9213c7c21215dbc288e0ee61770e6adf7752324c282"
		publicAccountKey  = "ad794c8f3da55845a5422506b4bd01ed1ae4e57378a82216d25d7351854b563d"
		// accountHash = "04d18b95474c8d7962a69bb386c788f1d7785b2cf3c26d9c7644516cf9652295"

		nftContractHash = "05560ca94e73f35c5b9b8a0f8b66e56238169e60ae421fb7b71c7ac3c6c744e2"                      // nft contract
		tokenID         = "7a78a717-bf19-4e7b-902f-7d4f334cb4c6"                                                  // hard code
		recipient       = "account-hash-04d18b95474c8d7962a69bb386c788f1d7785b2cf3c26d9c7644516cf9652295"         // my account
		spender         = "contract-package-wasm701ed1a382367a6016f3b389f75177030fd583c5b8838b4c04e92da6b4a11928" // market PACKAGE contract
	)

	ctx := context.Background()
	casperClient := contract.New(casperNodeAddress)

	// private -----------
	privateAccountKeyBytes, err := hex.DecodeString(privateAccountKey)
	require.NoError(t, err)
	publicAccountKeyBytes, err := hex.DecodeString(publicAccountKey)
	require.NoError(t, err)

	pair := casper_ed25519.ParseKeyPair(publicAccountKeyBytes, privateAccountKeyBytes)

	transfer := casper.NewTransfer(casperClient, func(b []byte) ([]byte, error) {
		casperSignature := pair.Sign(b)
		return casperSignature.SignatureData, nil
	})
	// -----------

	t.Run("mint one", func(t *testing.T) {
		txHash, err := transfer.MintOne(ctx, casper.MintOneRequest{
			PublicKey:              pair.PublicKey(),
			ChainName:              "casper-test",
			StandardPayment:        4100000000, // 4.1 CSPR.
			NFTContractPackageHash: nftContractHash,
			TokenID:                tokenID,
			Recipient:              recipient,
		})
		require.NoError(t, err)
		require.Empty(t, txHash)
	})

	t.Run("approve", func(t *testing.T) {
		txHash, err := transfer.ApproveNFT(ctx, casper.ApproveNFTRequest{
			PublicKey:              pair.PublicKey(),
			ChainName:              "casper-test",
			StandardPayment:        2500000000, // 2.5 CSPR.
			NFTContractPackageHash: nftContractHash,
			Spender:                spender,
			TokenID:                tokenID,
		})
		require.NoError(t, err)
		require.Empty(t, txHash)
	})
}

func TestCasper_MarketContract(t *testing.T) {
	t.Skip("for manual testing")
	var (
		casperNodeAddress = "http://116.202.169.210:7777/rpc"

		privateAccountKey = "4f0ffa6c3925d02127a9e9213c7c21215dbc288e0ee61770e6adf7752324c282"
		publicAccountKey  = "ad794c8f3da55845a5422506b4bd01ed1ae4e57378a82216d25d7351854b563d"
		// accountHash = "04d18b95474c8d7962a69bb386c788f1d7785b2cf3c26d9c7644516cf9652295"

		marketContractHash = "feed638f60f5a2840656d86e0e51dc62c092e79d980ba8dc281387dbb8f80c42"          // market contract
		nftContractHash    = "contract-05560ca94e73f35c5b9b8a0f8b66e56238169e60ae421fb7b71c7ac3c6c744e2" // nft contract
		tokenID            = "7a78a717-bf19-4e7b-902f-7d4f334cb4c6"                                      // hard code
		minBidPrice        = big.NewInt(3)                                                               // wei.
		redemptionPrice    = big.NewInt(10)                                                              // wei.
		auctionDuration    = big.NewInt(300000)                                                          // mb seconds.
	)

	ctx := context.Background()
	casperClient := contract.New(casperNodeAddress)

	// private -----------
	privateAccountKeyBytes, err := hex.DecodeString(privateAccountKey)
	require.NoError(t, err)
	publicAccountKeyBytes, err := hex.DecodeString(publicAccountKey)
	require.NoError(t, err)

	pair := casper_ed25519.ParseKeyPair(publicAccountKeyBytes, privateAccountKeyBytes)

	transfer := casper.NewTransfer(casperClient, func(b []byte) ([]byte, error) {
		casperSignature := pair.Sign(b)
		return casperSignature.SignatureData, nil
	})
	// -----------

	t.Run("create listing", func(t *testing.T) {
		txHash, err := transfer.CreateListing(ctx, casper.CreateListingRequest{
			PublicKey:          pair.PublicKey(),
			ChainName:          "casper-test",
			StandardPayment:    5700000000, // 5.7 CSPR.
			MarketContractHash: marketContractHash,
			NFTContractHash:    nftContractHash,
			TokenID:            tokenID,
			MinBidPrice:        minBidPrice,
			RedemptionPrice:    redemptionPrice,
			AuctionDuration:    auctionDuration,
		})
		require.NoError(t, err)
		require.Empty(t, txHash)
	})

	t.Run("make offer", func(t *testing.T) {
		txHash, err := transfer.MakeOffer(ctx, casper.MakeOfferRequest{
			PublicKey:          pair.PublicKey(),
			ChainName:          "casper-test",
			StandardPayment:    4400000000, // 4.4 CSPR.
			MarketContractHash: marketContractHash,
			NFTContractHash:    nftContractHash,
			TokenID:            tokenID,
			Amount:             minBidPrice,
		})
		require.NoError(t, err)
		require.Empty(t, txHash)
	})

	t.Run("accept offer", func(t *testing.T) {
		txHash, err := transfer.AcceptOffer(ctx, casper.AcceptOfferRequest{
			PublicKey:          pair.PublicKey(),
			ChainName:          "casper-test",
			StandardPayment:    10000000000, // 10 CSPR.
			MarketContractHash: marketContractHash,
			NFTContractHash:    nftContractHash,
			TokenID:            tokenID,
		})
		require.NoError(t, err)
		require.Empty(t, txHash)
	})

	t.Run("buy listing", func(t *testing.T) {
		txHash, err := transfer.BuyListing(ctx, casper.BuyListingRequest{
			PublicKey:          pair.PublicKey(),
			ChainName:          "casper-test",
			StandardPayment:    10000000000, // 10 CSPR.
			MarketContractHash: marketContractHash,
			NFTContractHash:    nftContractHash,
			TokenID:            tokenID,
		})
		require.NoError(t, err)
		require.Empty(t, txHash)
	})

	t.Run("final listing", func(t *testing.T) {
		txHash, err := transfer.FinalListing(ctx, casper.FinalListingRequest{
			PublicKey:          pair.PublicKey(),
			ChainName:          "casper-test",
			StandardPayment:    10000000000, // 10 CSPR.
			MarketContractHash: marketContractHash,
			NFTContractHash:    nftContractHash,
			TokenID:            tokenID,
		})
		require.NoError(t, err)
		require.Empty(t, txHash)
	})
}

func TestCasper_TokensContract(t *testing.T) {
	t.Skip("for manual testing")
	var (
		casperNodeAddress = "http://116.202.169.210:7777/rpc"

		privateAccountKey = "4f0ffa6c3925d02127a9e9213c7c21215dbc288e0ee61770e6adf7752324c282"
		publicAccountKey  = "ad794c8f3da55845a5422506b4bd01ed1ae4e57378a82216d25d7351854b563d"
		// accountHash = "04d18b95474c8d7962a69bb386c788f1d7785b2cf3c26d9c7644516cf9652295"

		tokensContractHash = "5aed0843516b06e4cbf56b1085c4af37035f2c9c1f18d7b0ffd7bbe96f91a3e0" // tokens contract
		spender            = "701ed1a382367a6016f3b389f75177030fd583c5b8838b4c04e92da6b4a11928" // market PACKAGE contract
		amount             = big.NewInt(100)                                                    // wei.
	)

	ctx := context.Background()
	casperClient := contract.New(casperNodeAddress)

	// private -----------
	privateAccountKeyBytes, err := hex.DecodeString(privateAccountKey)
	require.NoError(t, err)
	publicAccountKeyBytes, err := hex.DecodeString(publicAccountKey)
	require.NoError(t, err)

	pair := casper_ed25519.ParseKeyPair(publicAccountKeyBytes, privateAccountKeyBytes)

	transfer := casper.NewTransfer(casperClient, func(b []byte) ([]byte, error) {
		casperSignature := pair.Sign(b)
		return casperSignature.SignatureData, nil
	})
	// -----------

	t.Run("approve tokens", func(t *testing.T) {
		txHash, err := transfer.ApproveTokens(ctx, casper.ApproveTokensRequest{
			PublicKey:          pair.PublicKey(),
			ChainName:          "casper-test",
			StandardPayment:    100000000, // 0.1 CSPR.
			TokensContractHash: tokensContractHash,
			Spender:            common.HexToHash(spender),
			Amount:             amount,
		})
		require.NoError(t, err)
		require.Empty(t, txHash)
	})
}
