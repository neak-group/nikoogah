package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/neak-group/nikoogah/internal/app/charity/charity/entity"
	"github.com/neak-group/nikoogah/internal/infra/mongofx"
	"github.com/neak-group/nikoogah/utils/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type CharityMongoRepository struct {
	Logger *zap.Logger

	MongoClient         mongofx.MongoDBConn
	CharityDatabase     string
	CharitiesCollection string
}

func (r *CharityMongoRepository) getCollection(ctx context.Context) *mongo.Collection {
	db, err := r.MongoClient.GetDB(ctx, r.CharityDatabase)
	if err != nil {
		return nil
	}
	return db.Collection(r.CharitiesCollection)
}

// FindCharityTierID finds the charity tier ID by name
func (r *CharityMongoRepository) FindCharityTierID(ctx context.Context, name string) (uuid.UUID, error) {
	db, err := r.MongoClient.GetDB(ctx, r.CharityDatabase)
	if err != nil {
		if db == nil {
			err = fmt.Errorf("no database connection found")
			r.Logger.Error("Error getting MongoDB database", zap.Error(err))

		} else {

			r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		}
		return uuid.Nil, err
	}

	collection := db.Collection(r.CharitiesCollection)
	var result entity.Charity
	filter := bson.M{"name": name}

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return uuid.Nil, nil
		}
		r.Logger.Error("Error finding charity by name", zap.Error(err))
		return uuid.Nil, err
	}

	charityTierID, err := uuid.Parse(result.CharityTierID)
	if err != nil {
		r.Logger.Error("Error parsing charity tier ID", zap.Error(err))
		return uuid.Nil, err
	}

	return charityTierID, nil
}

// FetchCharity fetches charity by ID
func (r *CharityMongoRepository) FetchCharity(ctx context.Context, id uuid.UUID) (*entity.Charity, error) {
	db, err := r.MongoClient.GetDB(ctx, r.CharityDatabase)
	if err != nil {
		if db == nil {
			err = fmt.Errorf("no database connection found")
			r.Logger.Error("Error getting MongoDB database", zap.Error(err))

		} else {

			r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		}
		return nil, err
	}

	collection := db.Collection(r.CharitiesCollection)
	var charity entity.Charity
	filter := bson.M{"id": id}

	err = collection.FindOne(ctx, filter).Decode(&charity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		r.Logger.Error("Error fetching charity by ID", zap.Error(err))
		return nil, err
	}

	return &charity, nil
}

// CreateCharity creates a new charity
func (r *CharityMongoRepository) CreateCharity(ctx context.Context, charity *entity.Charity) (uuid.UUID, error) {
	db, err := r.MongoClient.GetDB(ctx, r.CharityDatabase)
	if err != nil {
		if db == nil {
			err = fmt.Errorf("no database connection found")
			r.Logger.Error("Error getting MongoDB database", zap.Error(err))

		} else {

			r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		}
		return uuid.Nil, err
	}

	collection := db.Collection(r.CharitiesCollection)
	charity.ID = uuid.New()
	charity.CreatedAt = time.Now()
	charity.UpdatedAt = charity.CreatedAt

	_, err = collection.InsertOne(ctx, charity)
	if err != nil {
		r.Logger.Error("Error creating new charity", zap.Error(err))
		return uuid.Nil, err
	}

	return charity.ID, nil
}

// SaveCharity updates an existing charity
func (r *CharityMongoRepository) SaveCharity(ctx context.Context, charity *entity.Charity) (uuid.UUID, error) {
	db, err := r.MongoClient.GetDB(ctx, r.CharityDatabase)
	if err != nil {
		if db == nil {
			err = fmt.Errorf("no database connection found")
			r.Logger.Error("Error getting MongoDB database", zap.Error(err))

		} else {

			r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		}
		return uuid.Nil, err
	}

	collection := db.Collection(r.CharitiesCollection)
	charity.UpdatedAt = time.Now()

	filter := bson.M{"id": charity.ID}
	update := bson.M{"$set": charity}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		r.Logger.Error("Error updating charity", zap.Error(err))
		return uuid.Nil, err
	}

	return charity.ID, nil
}

// FindRepresentativeByUserID finds a representative by user ID
func (r *CharityMongoRepository) FindRepresentativeByUserID(ctx context.Context, userID uuid.UUID) (*entity.Representative, error) {
	db, err := r.MongoClient.GetDB(ctx, r.CharityDatabase)
	if err != nil {
		if db == nil {
			err = fmt.Errorf("no database connection found")
			r.Logger.Error("Error getting MongoDB database", zap.Error(err))

		} else {

			r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		}
		return nil, err
	}

	collection := db.Collection(r.CharitiesCollection)
	var charity entity.Charity
	filter := bson.M{"representatives.user_id": userID}

	err = collection.FindOne(ctx, filter).Decode(&charity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		r.Logger.Error("Error finding representative by user ID", zap.Error(err))
		return nil, err
	}

	for _, rep := range charity.Representatives {
		if rep.UserID == userID {
			return rep, nil
		}
	}

	return nil, nil
}

// FindExistingRepresentativeByUserID checks if a representative exists by user ID
func (r *CharityMongoRepository) FindExistingRepresentativeByUserID(ctx context.Context, userID uuid.UUID) (bool, error) {
	db, err := r.MongoClient.GetDB(ctx, r.CharityDatabase)
	if err != nil {
		if db == nil {
			err = fmt.Errorf("no database connection found")
			r.Logger.Error("Error getting MongoDB database", zap.Error(err))

		} else {

			r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		}
		return false, err
	}

	collection := db.Collection(r.CharitiesCollection)
	filter := bson.M{"representatives.user_id": userID}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		r.Logger.Error("Error counting representatives", zap.Error(err))
		return false, err
	}

	return count > 0, nil
}
