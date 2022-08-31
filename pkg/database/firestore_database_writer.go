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

const nftCollectionsCollection = "sleeyaxTestCollections"

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

func (f *FirestoreDatabaseWriter) Write(ctx context.Context, collection *NFTCollection) error {
	// Write statistics from zora to a separate subcollection.
	// Note that the field on the collection struct is ignored, so we can safely write the values to a subcollection manually.
	if collection.ZoraStats != nil {
		m := toFirestoreMap(collection.ZoraStats)

		_, err := f.client.Collection(nftCollectionsCollection).Doc(fmt.Sprintf("%s:%s", collection.ZoraStats.ChainId, collection.ZoraStats.CollectionAddress)).Collection("collectionStats").Doc("all").Set(ctx, m, firestore.MergeAll)
		if err != nil {
			return err
		}
	}

	var opts []firestore.SetOption
	if collection.State.Create.Step == UnindexedStep {
		opts = append(opts, firestore.MergeAll)
	}

	m := toFirestoreMap(collection)

	_, err := f.client.Collection(nftCollectionsCollection).Doc(fmt.Sprintf("%s:%s", collection.ChainId, collection.Address)).Set(ctx, m, opts...)

	return err
}

func (f *FirestoreDatabaseWriter) Close() error {
	return f.client.Close()
}
