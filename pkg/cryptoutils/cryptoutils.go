// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cryptoutils

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Address defines address type.
type Address string

// CreateValidAddress creates valid address.
func CreateValidAddress(address Hex) Address {
	return Address(HexPrefix + address[LengthOneBlockInputValue-LengthAddress+LengthHexPrefix:])
}

// IsValidAddress checks if the address is valid.
func (address Address) IsValidAddress() bool {
	return common.IsHexAddress(string(address))
}

// Hex defines hex type.
type Hex string

// isHexCharacter returns bool of c being a valid hexadecimal.
func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// HexPrefix defines the prefix of hex type.
const HexPrefix Hex = "0x"

// PrivateKey defines private key type.
type PrivateKey string

// IsValidPrivateKey validates whether each byte is valid hexadecimal private key.
func (privateKey PrivateKey) IsValidPrivateKey() bool {
	if Length(len(string(privateKey))) != LengthPrivateKey {
		return false
	}
	for _, c := range []byte(string(privateKey)) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}

// PrivateKeyV defines v of private key type.
type PrivateKeyV int

const (
	// PrivateKeyVZero indicates that the v of private key is 0.
	PrivateKeyVZero PrivateKeyV = 0
	// PrivateKeyVOne indicates that the v of private key is 1.
	PrivateKeyVOne PrivateKeyV = 1
	// PrivateKeyVTwentySeven indicates that the v of private key is 27.
	PrivateKeyVTwentySeven PrivateKeyV = 27
	// PrivateKeyVTwentyEight indicates that the v of private key is 28.
	PrivateKeyVTwentyEight PrivateKeyV = 28
)

// Chain defines the list of possible chains in blockchain.
type Chain string

const (
	// ChainEthereum indicates that chain is ethereum.
	ChainEthereum Chain = "ethereum"
	// ChainPolygon indicates that chain is polygon.
	ChainPolygon Chain = "polygon"
)

// ChainID defines the list of possible number chains in blockchain.
type ChainID int

const (
	// ChainIDRinkeby indicates that chain id is 4.
	ChainIDRinkeby ChainID = 4
)

// Length defines the list of possible lengths of blockchain elements.
type Length int

const (
	// LengthPrivateKey defines length private key.
	LengthPrivateKey Length = 64
	// LengthOneBlockInputValue defines the length of one block of input data.
	LengthOneBlockInputValue Length = 64
	// LengthAddress defines the length of address.
	LengthAddress Length = 40
	// LengthHexPrefix defines the length of hex prefix.
	LengthHexPrefix Length = 2
)

// BlockTag defines the list of possible block tags in blockchain.
type BlockTag string

// BlockTagLatest indicates that the last block will be used.
const BlockTagLatest BlockTag = "latest"

// Data entity describes values for data field in transacton.
type Data struct {
	AddressContractMethod Hex
	TokenID               int64
}

// NewDataHex is a constructor for data entity, but returns hex string.
func NewDataHex(data Data) Hex {
	tokenID := createHexStringFixedLength(new(big.Int).SetInt64(data.TokenID))
	return data.AddressContractMethod + tokenID
}

// createHexStringFixedLength creates srings with fixed length and number in hex formate in the end.
func createHexStringFixedLength(value *big.Int) Hex {
	valueString := fmt.Sprintf("%x", value)
	var zeroString string
	for i := 0; i < (int(LengthOneBlockInputValue) - len(valueString)); i++ {
		zeroString += "0"
	}

	return Hex(zeroString + valueString)
}

// Contract entity describes addresses of contract and method.
type Contract struct {
	Address       Address `json:"address"`
	AddressMethod Hex     `json:"addressMethod"`
}
