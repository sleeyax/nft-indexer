package main

import (
	"context"
	"log"
	nft_indexer "nft-indexer"
	"nft-indexer/pkg/database"
	indexer2 "nft-indexer/pkg/indexer"
	"nft-indexer/pkg/indexer/ethereum"
)

func main() {
	// read config file
	config, err := nft_indexer.ParseConfig("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	collection := &database.NFTCollection{
		Address:       "0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D", // BAYC
		ChainId:       string(ethereum.MainNetwork),
		TokenStandard: database.ERC721,
	}

	idx, err := indexer2.New(config, indexer2.Ethereum, ethereum.MainNetwork)
	if err = idx.Start(context.Background(), collection); err != nil {
		log.Fatalln(err)
	}
}
