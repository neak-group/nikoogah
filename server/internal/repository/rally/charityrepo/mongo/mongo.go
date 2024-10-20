package mongo

import (
	"context"
	"errors"

	"github.com/neak-group/nikoogah/internal/app/rally/charity/entity"
	"github.com/neak-group/nikoogah/internal/infra/mongofx"
	"github.com/neak-group/nikoogah/utils/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type RlyCharityMongoRepository struct {
	Logger *zap.Logger

	MongoClient         mongofx.MongoDBConn
	RallyDatabase       string
	CharitiesCollection string
}

func (r *RlyCharityMongoRepository) getCollection(ctx context.Context) *mongo.Collection {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		return nil
	}
	return db.Collection(r.CharitiesCollection)
}

// FetchCharity fetches charity by CharityID
func (r *RlyCharityMongoRepository) FetchCharity(ctx context.Context, charityID uuid.UUID) (*entity.Charity, error) {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		return nil, err
	}

	collection := db.Collection(r.CharitiesCollection)
	var charity entity.Charity
	filter := bson.M{"charity_id": charityID}

	err = collection.FindOne(ctx, filter).Decode(&charity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // Charity not found
		}
		r.Logger.Error("Error fetching charity by CharityID", zap.Error(err))
		return nil, err
	}

	return &charity, nil
}

// SaveCharity saves or updates an existing charity
func (r *RlyCharityMongoRepository) SaveCharity(ctx context.Context, charity *entity.Charity) error {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		return err
	}

	collection := db.Collection(r.CharitiesCollection)

	filter := bson.M{"id": charity.CharityID}
	update := bson.M{"$set": charity}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		r.Logger.Error("Error updating charity", zap.Error(err))
		return err
	}
	return nil
}
