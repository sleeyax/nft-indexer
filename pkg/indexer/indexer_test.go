package indexer

import (
	"context"
	"nft-indexer/pkg/database"
	"testing"
)

type MockStep struct {
	CurrentStep database.CreationFlow
	NextStep    database.CreationFlow
	Executed    bool
}

func (m *MockStep) Execute(context context.Context, collection *database.NFTCollection) (database.CreationFlow, error) {
	m.Executed = true
	return m.NextStep, nil
}

func TestIndex(t *testing.T) {
	m1 := &MockStep{CurrentStep: database.CollectionCreator, NextStep: database.CollectionMetadata}
	m2 := &MockStep{CurrentStep: database.CollectionMetadata, NextStep: database.Complete}
	m3 := &MockStep{CurrentStep: database.Complete}

	collection := &database.NFTCollection{}

	indexer := new(Indexer)

	if err := indexer.Start(context.Background(), collection, map[database.CreationFlow]Step{
		m1.CurrentStep: m1,
		m2.CurrentStep: m2,
		m3.CurrentStep: m3,
	}); err != nil {
		t.Fatal(err)
	}

	if m1.Executed == false {
		t.Fail()
	}
	if m2.Executed == false {
		t.Fail()
	}
	if m3.Executed == false {
		t.Fail()
	}
}
