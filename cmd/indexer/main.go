package main

import (
	"context"
	"log"
	nft_indexer "nft-indexer"
	"nft-indexer/pkg/database"
	nftindexer "nft-indexer/pkg/indexer"
	"nft-indexer/pkg/indexer/ethereum"
)

func main() {
	// read config file
	config, err := nft_indexer.ParseConfig("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	// define collection to index
	collection := &database.NFTCollection{
		Address:       "0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D", // BAYC
		ChainId:       string(ethereum.MainNetwork),
		TokenStandard: database.ERC721,
	}

	ctx := context.Background()

	// connect to firestore DB
	db, err := database.NewFirestoreDatabaseWriter(ctx, config)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	ch := make(chan nftindexer.IndexResult)

	// start indexer in separate goroutine
	idx, err := nftindexer.New(config, nftindexer.Ethereum, ethereum.MainNetwork)
	go idx.Start(ctx, collection, ch)

	// keep looping in the main goroutine until the indexer goroutine closed the channel
	// NOTE: ideally this should be run in a separate goroutine for better scalability.
	for {
		// read incoming index results from the indexer's channel & write them to DB
		indexResult, ok := <-ch
		if !ok {
			return // channel was closed by the indexer, so we should stop this goroutine as well
		}

		// indexer returned an error and has exited
		if indexResult.Error != nil {
			log.Println(indexResult.Error)
			return
		}

		// log.Println(indexResult.Collection)
		if err = db.Write(ctx, collection); err != nil {
			log.Println(err)
			return
		}
	}
}
