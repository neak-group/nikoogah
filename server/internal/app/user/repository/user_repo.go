package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app/user/entity"
)

type UserRepository interface {
	FetchUser(ctx context.Context, userID uuid.UUID) (*entity.User, error)
	FetchUserByPhone(ctx context.Context, phone string) (*entity.User, error)
	SaveUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	ChangeUserState(ctx context.Context, userID uuid.UUID, newState entity.UserState) error
}
