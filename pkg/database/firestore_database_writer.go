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

const (
	nftCollectionsCollection = "sleeyaxTestCollections"
	maxBatchWritesPerRequest = 500
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

	// Write tokens a separate subcollection.
	// Note that the field on the collection struct is ignored, so we can safely write the values to a subcollection manually.
	if len(collection.Tokens) > 0 {
		batch := f.client.Batch()

		for i, token := range collection.Tokens {
			chunk := i + 1

			ref := f.client.Collection(nftCollectionsCollection).Doc(fmt.Sprintf("%s:%s", token.ChainId, token.CollectionAddress)).Collection("nfts").Doc(token.TokenId)
			m := toFirestoreMap(token)
			batch.Set(ref, m, firestore.MergeAll)

			// commit batch per x items (or if we are at the last item), so we don't exceed firestore's payload size limit
			if chunk%maxBatchWritesPerRequest == 0 || chunk == len(collection.Tokens) {
				_, err := batch.Commit(ctx)
				if err != nil {
					return err
				}
				batch = f.client.Batch()
			}
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
