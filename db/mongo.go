package db

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoCLient struct {
	*mongo.Client
	logr zerolog.Logger
}

func NewMongoClient(connectionURI string, logr zerolog.Logger) (*MongoCLient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &MongoCLient{client, logr}, nil
}

func (r *MongoCLient) Close() {
	if err := r.Disconnect(context.Background()); err != nil {
		r.logr.Error().Err(err).Msg("failed to disconnect mongodb client")
	}
}
