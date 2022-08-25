package main

import (
	"context"
	"flag"
	"github.com/pkg/errors"
	"log"
	"nft-indexer/pkg/config"
	"nft-indexer/pkg/database"
	indexer "nft-indexer/pkg/indexer"
	"nft-indexer/pkg/indexer/ethereum"
)

func main() {
	// read CLI flags
	var useFirestore, useConsole, useFile bool
	flag.BoolVar(&useFirestore, "firestore", false, "Write results to firestore")
	flag.BoolVar(&useConsole, "console", false, "Write results to console")
	flag.BoolVar(&useFile, "file", false, "Write results to a JSON file")
	flag.Parse()

	if !useFirestore && !useConsole && !useFile {
		log.Fatalln("Missing flags.\nUsage: ./indexer [-console, -firestore, -file]\nExamples:\n./indexer -console\n./indexer -firestore\n./indexer -console -firestore")
	}

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

	// connect to DB
	var db database.Writer
	if useFirestore {
		db, err = database.NewFirestoreDatabaseWriter(ctx, c)
		if err != nil {
			log.Fatalln(err)
		}
		defer db.Close()
	} else if useConsole {
		db = database.NewConsoleWriter()
	} else if useFile {
		panic("TODO: implement me")
	}

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

		// write results to whatever output formats were specified
		if indexResult.Collection != nil {
			if err = db.WriteNFTCollection(ctx, indexResult.Collection); err != nil {
				log.Println(err)
				return
			}
		}
		if indexResult.Stats != nil {
			if err = db.WriteStats(ctx, indexResult.Stats); err != nil {
				log.Println(err)
				return
			}
		}
	}
}
