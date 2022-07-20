package main

import (
	"log"
	nft_indexer "nft-indexer"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
)

func main() {
	// read config file
	config, err := nft_indexer.ParseConfig("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	address := "0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D" // BAYC
	network := ethereum.MainNetwork                         // 1

	// create and connect ethereum provider
	provider := ethereum.NewProvider(config)
	if err = provider.Connect(network); err != nil {
		log.Fatalln(err)
	}
	defer provider.Close()

	// create contract
	contract := ethereum.NewContract(address, network, provider)

	// parse the contract into an ERC-721 compatible token contract
	token, err := ethereum.NewTokenContract(contract, database.ERC721)
	if err != nil {
		log.Fatalln(err)
	}

	creator, err := token.GetCreator()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(creator)
}
