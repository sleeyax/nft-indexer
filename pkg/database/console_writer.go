package database

import (
	"context"
	"encoding/json"
	"log"
)

// ConsoleWriter is a simple database.Writer implementation that writes data to stdout instead.
// This is especially useful in development.
type ConsoleWriter struct{}

// NewConsoleWriter makes a new instance of ConsoleWriter.
func NewConsoleWriter() ConsoleWriter {
	return ConsoleWriter{}
}

func (c ConsoleWriter) Print(data interface{}) error {
	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	log.Println(string(b))

	return nil
}

func (c ConsoleWriter) WriteNFTCollection(ctx context.Context, collection *NFTCollection) error {
	return c.Print(collection)
}

func (c ConsoleWriter) WriteStats(ctx context.Context, stats *NftCollectionStats) error {
	return c.Print(stats)
}

func (c ConsoleWriter) Close() error {
	return nil
}
