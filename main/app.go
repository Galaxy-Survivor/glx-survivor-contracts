package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	NFT "app/nft"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
		FromBlock: big.NewInt(3),
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
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println("\n\n********************************")
		fmt.Println("BlockHash", vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println("BlockNumber", vLog.BlockNumber)   // 2394201
		fmt.Println("TxHash", vLog.TxHash.Hex())       // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6
		fmt.Println("Data", vLog.Data)                 // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		var eventRet struct {
			address common.Address
			id      *big.Int
		}

		err := contractAbi.UnpackIntoInterface(&eventRet, "event_MintedAircraft", vLog.Data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// for i := range interfaces {
		// 	fmt.Println((interfaces[i])) // bar
		// }
		fmt.Printf("Address %v - ID %v\n", eventRet.address.Hash(), eventRet.id) // foo
		// fmt.Println(string(inter.Value[:])) // bar

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}
