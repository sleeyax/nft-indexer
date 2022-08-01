package database

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"nft-indexer/pkg/config"
	"strings"
)

type FirestoreDatabaseWriter struct {
	client *firestore.Client
}

func NewFirestoreDatabaseWriter(ctx context.Context, cfg *config.Configuration) (*FirestoreDatabaseWriter, error) {
	serviceAccount := cfg.Gcloud.Firestore.ServiceAccount

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

func (f *FirestoreDatabaseWriter) Write(ctx context.Context, collection *NFTCollection, opts ...firestore.SetOption) error {
	m := toFirestoreMap(collection)

	_, err := f.client.Collection("sleeyaxTestCollections").Doc(fmt.Sprintf("%s:%s", collection.ChainId, collection.Address)).Set(ctx, m, opts...)

	return err
}

func (f *FirestoreDatabaseWriter) Close() error {
	return f.client.Close()
}
