/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"os"

	"github.com/hyperledgendary/conga-nft-contract/chaincode"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

type serverConfig struct {
	CCID    string
	Address string
}

func main() {
	nftContract := new(chaincode.TokenERC721Contract)
	nftContract.Info.Version = "0.0.1"
	nftContract.Info.Description = "ERC-721 fabric port"
	nftContract.Info.License = new(metadata.LicenseMetadata)
	nftContract.Info.License.Name = "Apache-2.0"
	nftContract.Info.Contact = new(metadata.ContactMetadata)
	nftContract.Info.Contact.Name = "Matias Salimbene"

	chaincode, err := contractapi.NewChaincode(nftContract)
	chaincode.Info.Title = "ERC-721 chaincode"
	chaincode.Info.Version = "0.0.1"

	if err != nil {
		panic("Could not create chaincode from TokenERC721Contract." + err.Error())
	}

	config := serverConfig{
		CCID:    os.Getenv("PACKAGE_ID"),
		Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
	}

	server := &shim.ChaincodeServer{
		CCID:    config.CCID,
		Address: config.Address,
		CC:      chaincode,
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}

	if err := server.Start(); err != nil {
		panic("Failed to start chaincode. " + err.Error())
	}
}
