package indexer

import (
	"context"
	"nft-indexer/pkg/config"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
	"nft-indexer/pkg/indexer/thirdparty/opensea"
	"nft-indexer/pkg/utils"
	"time"
)

func GetCollectionMetadata(ctx context.Context, config *config.Configuration, tokenContract *ethereum.TokenContract, collection *database.NFTCollection, sink *Sink) {
	step := database.CollectionMetadata

	os, err := opensea.NewOpenSea(utils.RandomItem(config.OpenSea.ApiKeys))
	if err != nil {
		sink.WriteError(err, step)
		return
	}

	osCollection, err := os.GetNFTCollection(tokenContract.Contract().Address.String())
	if err != nil {
		sink.WriteError(err, step)
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
	collection.State.Create = database.Create{
		Step:      database.TokenMetadata,
		UpdatedAt: time.Now().Unix(),
	}

	sink.Write(IndexResult{Collection: collection, Step: step})
}
