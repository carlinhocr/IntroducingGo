package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	var ethereum_mainnet string = "https://mainnet.infura.io/v3/580f184f75ad452eb839aeeacb65b62f"
	client, err := ethclient.Dial(ethereum_mainnet)
	//var celo_url string = "https://alfajores-forno.celo-testnet.org/"
	//client, err := ethclient.Dial(celo_url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 5671744
}
