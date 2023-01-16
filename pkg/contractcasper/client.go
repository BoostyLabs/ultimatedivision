package contract

import (
	"github.com/casper-ecosystem/casper-golang-sdk/sdk"
)

// ensures that rpcClient implement casper.Casper.
var _ Casper = (*rpcClient)(nil)

// rpcClient is a implementation of connector_service.Casper.
type rpcClient struct {
	client *sdk.RpcClient

	rpcNodeAddress string
}

// New is constructor for rpcClient.
func New(rpcNodeAddress string) Casper {
	client := sdk.NewRpcClient(rpcNodeAddress)
	return &rpcClient{
		client:         client,
		rpcNodeAddress: rpcNodeAddress,
	}
}

// GetCurrentBlockNumber returns current block number.
func (r *rpcClient) GetCurrentBlockNumber() (uint64, error) {
	blockResp, err := r.client.GetLatestBlock()
	if err != nil {
		return 0, err
	}

	return uint64(blockResp.Header.Height), nil
}
