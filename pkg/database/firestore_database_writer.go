package database

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	nft_indexer "nft-indexer"
)

type FirestoreDatabaseWriter struct {
	client *firestore.Client
}

func NewFirestoreDatabaseWriter(ctx context.Context, config *nft_indexer.Configuration) (*FirestoreDatabaseWriter, error) {
	serviceAccount := config.Gcloud.Firestore.ServiceAccount

	if serviceAccount.Type == "" {
		serviceAccount.Type = "service_account"
	}

	// TODO: find out why we can't read firebase creds from yaml
	/*b, err := json.Marshal(serviceAccount)
	if err != nil {
		return nil, err
	}*/

	clientOptions := option.WithCredentialsFile("creds.json")

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

func (f *FirestoreDatabaseWriter) Write(ctx context.Context, collection *NFTCollection) error {
	collection.Address = Normalize(collection.Address)
	_, err := f.client.Collection("sleeyaxTestCollections").Doc(fmt.Sprintf("%s:%s", collection.ChainId, collection.Address)).Set(ctx, collection, firestore.Merge())
	return err
}

func (f *FirestoreDatabaseWriter) Close() error {
	return f.client.Close()
}
