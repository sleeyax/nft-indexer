package database

import (
	"context"
	"io"
)

type Writer interface {
	io.Closer

	// Write nftCollections a NFT collection to a database or some other kind of storage.
	Write(ctx context.Context, collection *NFTCollection) error
}
