package main

import (
	"log"
	nft_indexer "nft-indexer"
)

func main() {
	config, err := nft_indexer.Configure()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(config.Alchemy.JsonRpc.MainNet[0])
}
