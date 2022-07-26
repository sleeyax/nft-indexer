package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	nft_indexer "nft-indexer"
	"strings"
)

type FirestoreDatabaseWriter struct {
	client *firestore.Client
}

func NewFirestoreDatabaseWriter(ctx context.Context, config *nft_indexer.Configuration) (*FirestoreDatabaseWriter, error) {
	serviceAccount := config.Gcloud.Firestore.ServiceAccount

	var clientOptions option.ClientOption
	if strings.HasSuffix(serviceAccount, ".json") {
		clientOptions = option.WithCredentialsFile(serviceAccount)
	} else {
		clientOptions = option.WithCredentialsJSON([]byte(serviceAccount))
	}

	app, err := firebase.NewApp(ctx, nil, clientOptions)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return &FirestoreDatabaseWriter{client}, nil
}

// toMap converts the NFT collection struct to a map.
//
// Maps unlock more functionality with the go firestore SDK compared to a struct.
func (f *FirestoreDatabaseWriter) toMap(collection *NFTCollection) (map[string]interface{}, error) {
	b, err := json.Marshal(collection)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}

	if err = json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return m, err
}

func (f *FirestoreDatabaseWriter) Write(ctx context.Context, collection *NFTCollection, opts ...firestore.SetOption) error {
	m, err := f.toMap(collection)

	_, err = f.client.Collection("sleeyaxTestCollections").Doc(fmt.Sprintf("%s:%s", collection.ChainId, collection.Address)).Set(ctx, m, opts...)

	return err
}

func (f *FirestoreDatabaseWriter) Close() error {
	return f.client.Close()
}
