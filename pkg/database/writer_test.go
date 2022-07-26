package database

import (
	"context"
	"testing"
)

type WriterMock struct {
	writes []*NFTCollection
}

func (w *WriterMock) Write(ctx context.Context, collection *NFTCollection) error {
	w.writes = append(w.writes, collection)
	return nil
}

func (w *WriterMock) Close() error {
	return nil
}

func TestWriter(t *testing.T) {
	ctx := context.Background()

	writer := new(WriterMock)
	defer writer.Close()

	if err := writer.Write(ctx, &NFTCollection{
		ChainId: "1",
		Address: "0x123",
	}); err != nil {
		t.Fatal(err)
	}
	if err := writer.Write(ctx, &NFTCollection{
		ChainId: "5",
		Address: "0x321",
	}); err != nil {
		t.Fatal(err)
	}

	if len(writer.writes) != 2 {
		t.Fatalf("expected 2 writes, got %d", len(writer.writes))
	}
}
