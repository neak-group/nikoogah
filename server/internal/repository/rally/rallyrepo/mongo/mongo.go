package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/neak-group/nikoogah/internal/app/rally/rally/entity"
	"github.com/neak-group/nikoogah/internal/infra/mongofx"
	"github.com/neak-group/nikoogah/utils/uuid"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type RallyMongoRepository struct {
	Logger *zap.Logger

	MongoClient     mongofx.MongoDBConn
	RallyDatabase   string
	RallyCollection string
}

// ApplyFilter applies a custom filter (for extensibility)
func (r *RallyMongoRepository) ApplyFilter(ctx context.Context, filter interface{}) {
	// Custom filter implementation, if needed
}

// FetchRally fetches a rally by its ID
func (r *RallyMongoRepository) FetchRally(ctx context.Context, rallyID uuid.UUID) (*entity.Rally, error) {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		return nil, err
	}

	collection := db.Collection(r.RallyCollection)
	var rally entity.Rally
	filter := bson.M{"id": rallyID}

	err = collection.FindOne(ctx, filter).Decode(&rally)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // Rally not found
		}
		r.Logger.Error("Error fetching rally by ID", zap.Error(err))
		return nil, err
	}

	return &rally, nil
}

// FetchRallies fetches all rallies
func (r *RallyMongoRepository) FetchRallies(ctx context.Context) ([]*entity.Rally, error) {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		return nil, err
	}

	collection := db.Collection(r.RallyCollection)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		r.Logger.Error("Error fetching rallies", zap.Error(err))
		return nil, err
	}
	defer cursor.Close(ctx)

	var rallies []*entity.Rally
	for cursor.Next(ctx) {
		var rally entity.Rally
		if err := cursor.Decode(&rally); err != nil {
			r.Logger.Error("Error decoding rally", zap.Error(err))
			return nil, err
		}
		rallies = append(rallies, &rally)
	}

	return rallies, nil
}

// FetchRalliesByFilter fetches rallies based on custom filters
func (r *RallyMongoRepository) FetchRalliesByFilter(ctx context.Context, filters ...interface{}) ([]*entity.Rally, error) {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		return nil, err
	}

	collection := db.Collection(r.RallyCollection)
	// Assuming filters is a list of BSON filters, combine them with $and operator
	filter := bson.M{"$and": filters}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		r.Logger.Error("Error fetching rallies by filter", zap.Error(err))
		return nil, err
	}
	defer cursor.Close(ctx)

	var rallies []*entity.Rally
	for cursor.Next(ctx) {
		var rally entity.Rally
		if err := cursor.Decode(&rally); err != nil {
			r.Logger.Error("Error decoding rally", zap.Error(err))
			return nil, err
		}
		rallies = append(rallies, &rally)
	}

	return rallies, nil
}

// CreateRally creates a new rally
func (r *RallyMongoRepository) CreateRally(ctx context.Context, rally *entity.Rally) (uuid.UUID, error) {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		return uuid.Nil, err
	}

	collection := db.Collection(r.RallyCollection)
	rally.ID = uuid.New()
	rally.CreatedAt = time.Now()
	rally.UpdatedAt = rally.CreatedAt

	_, err = collection.InsertOne(ctx, rally)
	if err != nil {
		r.Logger.Error("Error creating new rally", zap.Error(err))
		return uuid.Nil, err
	}

	return rally.ID, nil
}

// SaveRally updates an existing rally
func (r *RallyMongoRepository) SaveRally(ctx context.Context, rally *entity.Rally) error {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		return err
	}

	collection := db.Collection(r.RallyCollection)
	rally.UpdatedAt = time.Now()

	filter := bson.M{"id": rally.ID}
	update := bson.M{"$set": rally}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		r.Logger.Error("Error saving rally", zap.Error(err))
		return err
	}

	return nil
}

// UpdateParticipations updates human and fund participations in a rally
func (r *RallyMongoRepository) UpdateParticipations(ctx context.Context, rally *entity.Rally, hp *entity.HumanParticipation, fp *entity.FundParticipation) {
	// Update logic for participations (either human or fund)
	// Implement logic to update rally's participation fields
}

// FetchCharityRallyCount fetches the count of rallies for a specific charity
func (r *RallyMongoRepository) FetchCharityRallyCount(ctx context.Context, charityID uuid.UUID) (int, error) {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		return 0, err
	}

	collection := db.Collection(r.RallyCollection)
	filter := bson.M{"charity_id": charityID}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		r.Logger.Error("Error counting rallies for charity", zap.Error(err))
		return 0, err
	}

	return int(count), nil
}

// FetchRalliesByCharityID fetches rallies for a specific charity, optionally filtering only active rallies
func (r *RallyMongoRepository) FetchRalliesByCharityID(ctx context.Context, charityID uuid.UUID, onlyActive bool) ([]*entity.Rally, error) {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		return nil, err
	}

	collection := db.Collection(r.RallyCollection)
	filter := bson.M{"charity_id": charityID}
	if onlyActive {
		filter["state"] = entity.Active
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		r.Logger.Error("Error fetching rallies by charity ID", zap.Error(err))
		return nil, err
	}
	defer cursor.Close(ctx)

	var rallies []*entity.Rally
	for cursor.Next(ctx) {
		var rally entity.Rally
		if err := cursor.Decode(&rally); err != nil {
			r.Logger.Error("Error decoding rally document", zap.Error(err))
			return nil, err
		}
		rallies = append(rallies, &rally)
	}

	return rallies, nil
}

// FetchRallyParticipationCount fetches the count of participations for a specific rally by participation status
func (r *RallyMongoRepository) FetchRallyParticipationCount(ctx context.Context, rallyID uuid.UUID, status entity.ParticipationStatus) (int, error) {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		return 0, err
	}

	collection := db.Collection(r.RallyCollection)
	filter := bson.M{
		"id":                          rallyID,
		"human_participations.status": status,
	}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		r.Logger.Error("Error counting rally participations by status", zap.Error(err))
		return 0, err
	}

	return int(count), nil
}

// FetchTargetFund fetches the target fund amount for a specific rally
func (r *RallyMongoRepository) FetchTargetFund(ctx context.Context, rallyID uuid.UUID) (decimal.Decimal, error) {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		return decimal.Zero, err
	}

	collection := db.Collection(r.RallyCollection)
	var rally entity.Rally
	filter := bson.M{"id": rallyID}

	err = collection.FindOne(ctx, filter).Decode(&rally)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return decimal.Zero, nil // Rally not found
		}
		r.Logger.Error("Error fetching target fund for rally", zap.Error(err))
		return decimal.Zero, err
	}

	return rally.FundAmount, nil
}
