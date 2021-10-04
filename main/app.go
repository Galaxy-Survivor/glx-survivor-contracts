package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	NFT "app/nft"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	args := os.Args[1:]
	client, err := ethclient.Dial(args[0])
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(args[1])
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(1),
		ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// contractAbi, err := abi.JSON(strings.NewReader(string(NFT.AircraftNFTABI)))

	aircraftNFT, err := NFT.NewAircraftNFT(contractAddress, client)

	fmt.Println(len(logs))
	for _, vLog := range logs {
		fmt.Println("\n\n********************************")
		fmt.Println("BlockHash", vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println("BlockNumber", vLog.BlockNumber)   // 2394201
		fmt.Println("TxHash", vLog.TxHash.Hex())       // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6
		fmt.Println("Data", vLog.Data)                 // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		// var event struct {
		// 	address common.Address
		// 	id      *big.Int
		// }

		event, err := aircraftNFT.AircraftNFTFilterer.ParseEventMintedAircraft(vLog)
		if err != nil {
			fmt.Println("Error", err)
			continue
		}
		// for i := range interfaces {
		//      fmt.Println((interfaces[i])) // bar
		// }
		fmt.Printf("Address %v - ID %v\n", event.Operator.Hash(), event.AircraftId) // foo
		// fmt.Println(string(inter.Value[:])) // bar
		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}
		fmt.Println("TOPIC: ", topics[0])
		fmt.Println("====================")
	}
	// NFT.WatchEventMintedAircraft(client)
}
