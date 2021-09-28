package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	NFT "app/nft"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("ws://192.168.1.199:8545")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xe31267c0eedB0F23DB3C62E8eBa3dFf13e18345b")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(NFT.AircraftNFTABI)))
	if err != nil {
		log.Fatal("ERROR PARSE:", err)
	}
	fmt.Println("====================")
	for _, vLog := range logs {
		fmt.Println(">BLOCK HASH: ", vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(">BLOCK NUMBER: ", vLog.BlockNumber)   // 2394201
		fmt.Println(">TX HASH: ", vLog.TxHash.Hex())       // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6
		fmt.Println(">DATA: ", vLog.Data)                  // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		interfaces, err := contractAbi.Unpack("event_MintedAircraft", vLog.Data)
		if err != nil {
			// log.Fatal("ERROR UNPACK: ", err)
			continue
		}
		for i := range interfaces {
			fmt.Println(interfaces[i]) // bar
		}
		fmt.Println(interfaces) // foo
		// fmt.Println(string(inter.Value[:])) // bar

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}
		fmt.Println("TOPIC: ", topics[0])
		fmt.Println("====================")
	}

	eventSignature := []byte("event_MintedAircraft(address,uint256)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println("eventSignature: ", hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}
