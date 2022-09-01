package indexer

import (
	"context"
	"fmt"
	"log"
	"math"
	"nft-indexer/pkg/config"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
	"nft-indexer/pkg/utils"
	"time"
)

const maxSupply = 30_000

func AggregateMetadata(ctx context.Context, config *config.Configuration, tokenContract *ethereum.TokenContract, collection *database.NFTCollection, sink *Sink) {
	if collection.NumNfts > maxSupply {
		log.Println(fmt.Sprintf("colleciton %s has too many tokens to aggregate metadata", collection.Address))

		collection.NumTraitTypes = 0
		collection.NumOwners = 0
	} else {
		var attributes = make(map[string]database.Attribute)

		incrementTrait := func(value string, traitType string, displayType database.DisplayType) {
			if traitType == "" {
				traitType = value
			}

			// initialize traitType if it doesn't exist
			if _, ok := attributes[traitType]; !ok {
				attributes[traitType] = database.Attribute{
					DisplayType: displayType,
					Type:        traitType,
					Slug:        utils.ToSearchFriendly(traitType),
				}
			}

			// initialize values map if it doesn't exist
			if attr, ok := attributes[traitType]; ok {
				if attr.Values == nil {
					attr.Values = make(map[string]database.AttributeMetadata)
				}
				attributes[traitType] = attr
			}

			// initialize value if it doesn't exist
			if _, ok := attributes[traitType].Values[value]; !ok {
				attributes[traitType].Values[value] = database.AttributeMetadata{
					Type:      traitType,
					TypeSlug:  utils.ToSearchFriendly(traitType),
					Value:     value,
					ValueSlug: utils.ToSearchFriendly(value),
				}
			}

			// increment counts
			if attr, ok := attributes[traitType]; ok {
				attr.Count++
				attr.Percent = math.Round((float64(attr.Count)/float64(len(collection.Tokens)))*float64(100)*float64(100)) / float64(100)
				if val, ok := attr.Values[value]; ok {
					val.Count++
					val.Percent = math.Round((float64(val.Count)/float64(len(collection.Tokens)))*float64(100)*float64(100)) / float64(100)
					val.RarityScore = 1 / (val.Percent / 100)
					attr.Values[value] = val
				}
				attributes[traitType] = attr
			}
		}

		for _, token := range collection.Tokens {
			for _, attribute := range token.Metadata.Attributes {
				incrementTrait(attribute.Value, attribute.TraitType, attribute.DisplayType)
			}
		}

		collection.Attributes = attributes
	}

	collection.State.Create = database.Create{
		Step:      database.CompleteStep,
		UpdatedAt: time.Now().Unix(),
	}

	sink.Write(IndexResult{Collection: collection, Step: database.AggregateMetadataStep})
}
