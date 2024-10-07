package userrepo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app/user/entity"
	"github.com/neak-group/nikoogah/internal/infra/mongofx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type userMongoRepository struct {
	logger *zap.Logger

	mongoClient     mongofx.MongoDBConn
	usersCollection string
}

func (repo *userMongoRepository) FetchUser(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	db, err := repo.mongoClient.GetDB(ctx)
	if err != nil {
		repo.logger.Error("failed to get MongoDB connection", zap.Error(err))
		return nil, err
	}

	collection := db.Collection(repo.usersCollection)

	var user entity.User
	err = collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			repo.logger.Info("user not found", zap.String("userID", userID.String()))
			return nil, nil
		}
		repo.logger.Error("failed to fetch user", zap.Error(err))
		return nil, err
	}

	return &user, nil
}

func (repo *userMongoRepository) FetchUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	db, err := repo.mongoClient.GetDB(ctx)
	if err != nil {
		repo.logger.Error("failed to get MongoDB connection", zap.Error(err))
		return nil, err
	}

	collection := db.Collection(repo.usersCollection)

	var user entity.User
	err = collection.FindOne(ctx, bson.M{"phone_number": phone}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			repo.logger.Info("user not found by phone", zap.String("phone", phone))
			return nil, nil
		}
		repo.logger.Error("failed to fetch user by phone", zap.Error(err))
		return nil, err
	}

	return &user, nil
}

func (repo *userMongoRepository) SaveUser(ctx context.Context, user *entity.User) error {
	db, err := repo.mongoClient.GetDB(ctx)
	if err != nil {
		repo.logger.Error("failed to get MongoDB connection", zap.Error(err))
		return err
	}

	collection := db.Collection(repo.usersCollection)

	filter := bson.M{"_id": user.ID}
	update := bson.M{
		"$set": bson.M{
			"first_name":        user.FirstName,
			"last_name":         user.LastName,
			"phone_number":      user.PhoneNumber,
			"phone_verified_at": user.PhoneVerifiedAt,
			"national_code":     user.NationalCode,
			"avatar_path":       user.AvatarPath,
			"resume_path":       user.ResumePath,
			"user_state":        user.UserState,
			"updated_at":        time.Now(),
		},
		"$setOnInsert": bson.M{
			"created_at": user.CreatedAt,
		},
	}

	opts := options.Update().SetUpsert(true)
	_, err = collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		repo.logger.Error("failed to save user", zap.Error(err))
		return err
	}

	return nil
}

// DeleteUser removes a user by their UUID from MongoDB.
func (repo *userMongoRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	db, err := repo.mongoClient.GetDB(ctx)
	if err != nil {
		repo.logger.Error("failed to get MongoDB connection", zap.Error(err))
		return err
	}

	collection := db.Collection(repo.usersCollection)

	_, err = collection.DeleteOne(ctx, bson.M{"_id": userID})
	if err != nil {
		repo.logger.Error("failed to delete user", zap.Error(err))
		return err
	}

	return nil
}

func (repo *userMongoRepository) ChangeUserState(ctx context.Context, userID uuid.UUID, newState entity.UserState) error {
	db, err := repo.mongoClient.GetDB(ctx)
	if err != nil {
		repo.logger.Error("failed to get MongoDB connection", zap.Error(err))
		return err
	}

	collection := db.Collection(repo.usersCollection)

	update := bson.M{
		"$set": bson.M{
			"user_state": newState,
			"updated_at": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": userID}, update)
	if err != nil {
		repo.logger.Error("failed to change user state", zap.Error(err))
		return err
	}

	return nil
}
