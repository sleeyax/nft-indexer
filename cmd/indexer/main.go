package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/pkg/errors"
	"log"
	"nft-indexer/pkg/config"
	"nft-indexer/pkg/database"
	indexer "nft-indexer/pkg/indexer"
	"nft-indexer/pkg/indexer/ethereum"
)

func main() {
	// read config file
	c, err := config.ParseConfig("config.yaml")
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
	db, err := database.NewFirestoreDatabaseWriter(ctx, c)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	ch := make(chan indexer.IndexResult)

	// start indexer in separate goroutine
	idx, err := indexer.New(c, indexer.Ethereum, ethereum.MainNetwork)
	go idx.Start(ctx, collection, ch)

	// keep looping in the main goroutine until the indexer goroutine closed the channel
	// NOTE: ideally this should be run in a separate goroutine for better scalability.
	for {
		// read incoming index results from the indexer's channel & write them to DB
		indexResult, ok := <-ch
		if !ok {
			return // channel was closed by the indexer, so we should stop this goroutine as well
		}

		// an indexer step returned an error, skip current loop
		if indexResult.Error != nil {
			log.Println(errors.WithMessage(indexResult.Error, "indexing step error"))
			break
		}

		// an indexer step returned a warning, continue
		if indexResult.Warning != nil {
			log.Println(errors.WithMessage(indexResult.Warning, "indexing step warning"))
		}

		var writeOptions []firestore.SetOption
		if collection.State.Create.Step == database.Unindexed {
			writeOptions = append(writeOptions, firestore.MergeAll)
		}
		log.Println(indexResult.Collection)
		log.Println(indexResult.Stats)
		/*if err = db.Write(ctx, collection, writeOptions...); err != nil {
			log.Println(err)
			return
		}*/
	}
}
