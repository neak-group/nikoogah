package services

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app/user/entity"
)

type SessionService interface {
	NewSession(ctx context.Context, userID string, fullName string, deviceInfo entity.DeviceInfo) (sessionID *entity.Session, err error)
	Nullify(ctx context.Context, sessionID string) error
	ValidateSession(ctx context.Context, sessionID string) (session *entity.Session, err error)
}
