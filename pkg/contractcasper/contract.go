// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

package contract

import (
	"context"
	"encoding/hex"
	"encoding/json"

	"github.com/casper-ecosystem/casper-golang-sdk/keypair"
	"github.com/casper-ecosystem/casper-golang-sdk/sdk"
	"github.com/zeebo/errs"
)

// ErrContract indicates that there was an error in the contract package.
var ErrContract = errs.Class("contract package")

// BridgeInRequest describes values to initiate inbound bridge transaction.
type BridgeInRequest struct {
	Deploy         string
	RPCNodeAddress string
}

// StringNetworkAddress describes an address for some network.
type StringNetworkAddress struct {
	NetworkName string
	Address     string
}

// BridgeInResponse describes bridge in tx hash.
type BridgeInResponse struct {
	Txhash string
}

// BridgeIn initiates inbound bridge transaction.
func BridgeIn(ctx context.Context, req BridgeInRequest) (BridgeInResponse, error) {
	request := struct {
		Deploy struct {
			Hash      sdk.Hash                  `json:"hash"`
			Header    *sdk.DeployHeader         `json:"header"`
			Payment   *sdk.ExecutableDeployItem `json:"payment"`
			Session   *sdk.ExecutableDeployItem `json:"session"`
			Approvals []struct {
				Signer    string `json:"signer"`
				Signature string `json:"signature"`
			} `json:"approvals"`
		}
	}{}

	err := json.Unmarshal([]byte(req.Deploy), &request)
	if err != nil {
		return BridgeInResponse{}, ErrContract.Wrap(err)
	}

	pubKeyData, err := hex.DecodeString(request.Deploy.Approvals[0].Signer[2:])
	if err != nil {
		return BridgeInResponse{}, ErrContract.Wrap(err)
	}

	signer := keypair.PublicKey{
		Tag:        request.Deploy.Header.Account.Tag,
		PubKeyData: pubKeyData,
	}

	signatureData, err := hex.DecodeString(request.Deploy.Approvals[0].Signature[2:])
	if err != nil {
		return BridgeInResponse{}, ErrContract.Wrap(err)
	}

	signature := keypair.Signature{
		Tag:           request.Deploy.Header.Account.Tag,
		SignatureData: signatureData,
	}

	approval := sdk.Approval{
		Signer:    signer,
		Signature: signature,
	}

	deploy := sdk.Deploy{
		Hash:      request.Deploy.Hash,
		Header:    request.Deploy.Header,
		Payment:   request.Deploy.Payment,
		Session:   request.Deploy.Session,
		Approvals: []sdk.Approval{approval},
	}

	casperClient := sdk.NewRpcClient(req.RPCNodeAddress)
	deployResp, err := casperClient.PutDeploy(deploy)
	if err != nil {
		return BridgeInResponse{}, ErrContract.Wrap(err)
	}

	resp := BridgeInResponse{
		Txhash: deployResp.Hash,
	}

	return resp, nil
}
