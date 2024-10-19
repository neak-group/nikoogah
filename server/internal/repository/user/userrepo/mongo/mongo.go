package mongo

import (
	"context"
	"time"

	"github.com/neak-group/nikoogah/internal/app/user/entity"
	"github.com/neak-group/nikoogah/internal/infra/mongofx"
	"github.com/neak-group/nikoogah/utils/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type UserMongoRepository struct {
	Logger *zap.Logger

	MongoClient     mongofx.MongoDBConn
	UserDatabase    string
	UsersCollection string
}

func (repo *UserMongoRepository) FetchUser(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	db, err := repo.MongoClient.GetDB(ctx, repo.UserDatabase)
	if err != nil {
		repo.Logger.Error("failed to get MongoDB connection", zap.Error(err))
		return nil, err
	}

	collection := db.Collection(repo.UsersCollection)

	var user entity.User
	err = collection.FindOne(ctx, bson.M{"id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			repo.Logger.Info("user not found", zap.String("userID", userID.String()))
			return nil, nil
		}
		repo.Logger.Error("failed to fetch user", zap.Error(err))
		return nil, err
	}

	return &user, nil
}

func (repo *UserMongoRepository) FetchUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	db, err := repo.MongoClient.GetDB(ctx, repo.UserDatabase)
	if err != nil {
		repo.Logger.Error("failed to get MongoDB connection", zap.Error(err))
		return nil, err
	}

	collection := db.Collection(repo.UsersCollection)

	var user entity.User

	err = collection.FindOne(ctx, bson.M{"phone_number.phone_number": phone}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			repo.Logger.Info("user not found by phone", zap.String("phone", phone))
			return nil, nil
		}
		repo.Logger.Error("failed to fetch user by phone", zap.Error(err))
		return nil, err
	}

	return &user, nil
}

func (repo *UserMongoRepository) SaveUser(ctx context.Context, user *entity.User) error {
	db, err := repo.MongoClient.GetDB(ctx, repo.UserDatabase)
	if err != nil {
		repo.Logger.Error("failed to get MongoDB connection", zap.Error(err))
		return err
	}

	collection := db.Collection(repo.UsersCollection)

	filter := bson.M{"id": user.ID}
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
		repo.Logger.Error("failed to save user", zap.Error(err))
		return err
	}

	return nil
}

// DeleteUser removes a user by their UUID from MongoDB.
func (repo *UserMongoRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	db, err := repo.MongoClient.GetDB(ctx, repo.UserDatabase)
	if err != nil {
		repo.Logger.Error("failed to get MongoDB connection", zap.Error(err))
		return err
	}

	collection := db.Collection(repo.UsersCollection)

	_, err = collection.DeleteOne(ctx, bson.M{"id": userID})
	if err != nil {
		repo.Logger.Error("failed to delete user", zap.Error(err))
		return err
	}

	return nil
}

func (repo *UserMongoRepository) ChangeUserState(ctx context.Context, userID uuid.UUID, newState entity.UserState) error {
	db, err := repo.MongoClient.GetDB(ctx, repo.UserDatabase)
	if err != nil {
		repo.Logger.Error("failed to get MongoDB connection", zap.Error(err))
		return err
	}

	collection := db.Collection(repo.UsersCollection)

	update := bson.M{
		"$set": bson.M{
			"user_state": newState,
			"updated_at": time.Now(),
		},
	}

	repo.Logger.Info("user id when update", zap.String("userID", userID.String()))
	_, err = collection.UpdateOne(ctx, bson.M{"id": userID}, update)
	if err != nil {
		repo.Logger.Error("failed to change user state", zap.Error(err))
		return err
	}

	return nil
}
