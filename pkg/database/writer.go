package database

import (
	"context"
	"io"
)

type Writer interface {
	io.Closer

	// WriteNFTCollection nftCollections a NFT collection to a database or some other kind of storage.
	WriteNFTCollection(ctx context.Context, collection *NFTCollection) error

	// WriteStats nftCollections NFT collection stats to a database or some other kind of storage.
	WriteStats(ctx context.Context, stats *NftCollectionStats) error
}
