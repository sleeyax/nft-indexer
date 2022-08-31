package indexer

import (
	"context"
	"nft-indexer/pkg/config"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
	"nft-indexer/pkg/indexer/thirdparty/opensea"
	"nft-indexer/pkg/indexer/thirdparty/zora"
	"nft-indexer/pkg/utils"
	"sync"
	"time"
)

func GetCollectionMetadata(ctx context.Context, config *config.Configuration, tokenContract *ethereum.TokenContract, collection *database.NFTCollection, sink *Sink) {
	var wg sync.WaitGroup

	wg.Add(2)

	go writeOpenSeaCollectionMetadata(&wg, utils.RandomItem(config.OpenSea.ApiKeys), tokenContract.Contract().Address.String(), collection, sink)
	go writeAggregatedStats(&wg, config.Zora.ApiKey, tokenContract.Contract().Address.String(), collection, sink)

	wg.Wait()

	// The collection name should be set.
	// If not, it means we couldn't fetch the collection details from OpenSea.
	// This is considered fatal and the error (containing more details about the failure) is already logged in the goroutine.
	if collection.Metadata.Name == "" {
		return
	}

	collection.State.Create = database.Create{
		Step:      database.TokenMetadata,
		UpdatedAt: time.Now().Unix(),
	}

	sink.Write(IndexResult{Collection: collection, Step: database.CollectionMetadata})
}

// writeOpenSeaCollectionMetadata writes collection details and metadata from OpenSea to the collection.
func writeOpenSeaCollectionMetadata(wg *sync.WaitGroup, apiKey string, address string, collection *database.NFTCollection, sink *Sink) {
	defer wg.Done()

	os, err := opensea.NewOpenSea(apiKey)
	if err != nil {
		panic(err)
		return
	}

	osCollection, err := os.GetNFTCollection(address)
	if err != nil {
		sink.WriteError(err, database.CollectionMetadata)
		return
	}

	symbol := osCollection.Symbol
	if len(osCollection.Collection.PrimaryAssetContracts) > 0 {
		symbol = osCollection.Collection.PrimaryAssetContracts[0].Symbol
	}

	collection.Metadata.Name = utils.OrString(osCollection.Collection.Name, osCollection.Name)
	collection.Metadata.Description = utils.OrString(osCollection.Collection.Description, osCollection.Description)
	collection.Metadata.Symbol = symbol
	collection.Metadata.ProfileImage = utils.OrString(osCollection.Collection.ImageUrl, osCollection.Collection.FeaturedImageUrl, osCollection.ImageUrl)
	collection.Metadata.BannerImage = osCollection.Collection.BannerImageUrl
	collection.Metadata.DisplayType = osCollection.Collection.DisplayData.CardDisplayStyle
	collection.Metadata.Links = database.Links{
		Timestamp: time.Now().UnixMilli(),
		Discord:   osCollection.Collection.DiscordUrl,
		Telegram:  utils.OrString(osCollection.Collection.TelegramUrl),
		Twitter:   utils.OrString(osCollection.Collection.TwitterUsername),
		Instagram: utils.OrString(osCollection.Collection.InstagramUsername),
		Wiki:      utils.OrString(osCollection.Collection.WikiUrl),
		Medium:    utils.Ternary(osCollection.Collection.MediumUsername != "", osCollection.Collection.MediumUsername, ""),
		External:  osCollection.Collection.ExternalUrl,
		Slug:      utils.OrString(osCollection.Collection.Slug),
	}
	collection.HasBlueCheck = osCollection.Collection.SafelistRequestStatus == "verified"
	collection.Slug = utils.Ternary(osCollection.Collection.Slug != "", utils.ToSearchFriendly(osCollection.Collection.Slug), "")
}

func writeAggregatedStats(wg *sync.WaitGroup, apiKey string, address string, collection *database.NFTCollection, sink *Sink) {
	defer wg.Done()

	z, err := zora.NewGraphQLClient(apiKey)
	if err != nil {
		panic(err)
		return
	}

	stats, err := z.AggregateStat(address, 10)
	if err != nil {
		sink.WriteError(err, database.CollectionMetadata)
		return
	}

	collection.ZoraStats = &database.ZoraNFTStats{
		ChainId:                   collection.ChainId,
		CollectionAddress:         collection.Address,
		Volume:                    stats.Data.AggregateStat.SalesVolume.ChainTokenPrice,
		NumSales:                  stats.Data.AggregateStat.SalesVolume.TotalCount,
		VolumeUSDC:                stats.Data.AggregateStat.SalesVolume.UsdcPrice,
		NumOwners:                 stats.Data.AggregateStat.OwnerCount,
		NumNfts:                   stats.Data.AggregateStat.NftCount,
		TopOwnersByOwnedNftsCount: stats.Data.AggregateStat.OwnersByCount.Nodes,
		UpdatedAt:                 time.Now().UnixMilli(),
	}
}
