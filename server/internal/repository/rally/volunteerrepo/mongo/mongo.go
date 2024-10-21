package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/neak-group/nikoogah/internal/app/rally/volunteer/entity"
	"github.com/neak-group/nikoogah/internal/infra/mongofx"
	"github.com/neak-group/nikoogah/utils/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type VolunteerMongoRepository struct {
	Logger *zap.Logger

	MongoClient          mongofx.MongoDBConn
	RallyDatabase        string
	VolunteersCollection string
}

// FetchVolunteer fetches a volunteer by their UserID
func (r *VolunteerMongoRepository) FetchVolunteer(ctx context.Context, id uuid.UUID) (*entity.Volunteer, error) {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		return nil, err
	}

	collection := db.Collection(r.VolunteersCollection)
	var volunteer entity.Volunteer
	filter := bson.M{"user_id": id}

	err = collection.FindOne(ctx, filter).Decode(&volunteer)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // Volunteer not found
		}
		r.Logger.Error("Error fetching volunteer by UserID", zap.Error(err))
		return nil, err
	}

	return &volunteer, nil
}

// UpdateVolunteer updates or inserts a volunteer record
func (r *VolunteerMongoRepository) UpdateVolunteer(ctx context.Context, volunteer *entity.Volunteer) error {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		return err
	}

	collection := db.Collection(r.VolunteersCollection)

	// Update volunteer if exists, or insert if not
	filter := bson.M{"user_id": volunteer.UserID}
	update := bson.M{
		"$set": bson.M{
			"full_name":                   volunteer.FullName,
			"reputation":                  volunteer.Reputation,
			"resume_file":                 volunteer.ResumeFile,
			"volunteer_transactions":      volunteer.VolunteerTransactions,
			"volunteering_request_number": volunteer.VolunteeringRequestNumber,
			"financial_aids_sum":          volunteer.FinancialAidsSum,
			"updated_at":                  time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		r.Logger.Error("Error updating volunteer", zap.Error(err))
		return err
	}

	return nil
}

// FetchVolunteersByBatchID fetches multiple volunteers based on a list of UserID (batch)
func (r *VolunteerMongoRepository) FetchVolunteersByBatchID(ctx context.Context, ids []uuid.UUID) ([]*entity.Volunteer, error) {
	db, err := r.MongoClient.GetDB(ctx, r.RallyDatabase)
	if err != nil {
		r.Logger.Error("Error getting MongoDB database", zap.Error(err))
		return nil, err
	}

	collection := db.Collection(r.VolunteersCollection)

	// Convert UUID slice to interface slice for BSON
	var uuidInterfaces []interface{}
	for _, id := range ids {
		uuidInterfaces = append(uuidInterfaces, id)
	}

	// Filter using $in operator to find all volunteers with UserIDs in the given list
	filter := bson.M{"user_id": bson.M{"$in": uuidInterfaces}}

	// Set options to limit, sort, or manage projection if necessary (for future enhancements)
	opts := options.Find()

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		r.Logger.Error("Error fetching volunteers by batch ID", zap.Error(err))
		return nil, err
	}
	defer cursor.Close(ctx)

	var volunteers []*entity.Volunteer
	for cursor.Next(ctx) {
		var volunteer entity.Volunteer
		if err := cursor.Decode(&volunteer); err != nil {
			r.Logger.Error("Error decoding volunteer document", zap.Error(err))
			return nil, err
		}
		volunteers = append(volunteers, &volunteer)
	}

	if err := cursor.Err(); err != nil {
		r.Logger.Error("Cursor error during batch fetch", zap.Error(err))
		return nil, err
	}

	return volunteers, nil
}
