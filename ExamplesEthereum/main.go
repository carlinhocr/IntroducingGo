package main

import (
	"context"
	"fmt"
	"github.com/celo-org/celo-blockchain/ethclient"
	//"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	//var ethereum_mainnet string = put string from credentials.txt file
	//client, err := ethclient.Dial(ethereum_mainnet)
	//var celo_alfajores_url string = "https://alfajores-forno.celo-testnet.org/"
	//client, err := ethclient.Dial(celo_alfajores_url)
	//file, err := os.Open("./ExamplesEthereum/credentials.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//b1 := make([]byte, 20)
	//n1, err := file.Read(b1)
	//fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	//fmt.Print(file)

	var celo_forno_url string = "https://forno.celo.org/"
	client, err := ethclient.Dial(celo_forno_url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Latest block number: ", header.Number.String(), " for network ", celo_forno_url) // 5671744
}
