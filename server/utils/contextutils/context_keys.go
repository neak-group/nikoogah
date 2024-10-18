package contextutils

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type ContextKeyUserID struct{}

func GetUserIDFromCtx(ctx context.Context) (uuid.UUID, error) {
	v := ctx.Value(ContextKeyUserID{})
	if v == nil {
		return uuid.Nil, fmt.Errorf("no valid user id context found")
	}
	rawID, ok := v.(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("no valid user id context found")
	}

	id, err := uuid.Parse(rawID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("no valid user id context found")
	}

	return id, nil
}

func SetUserIDCtx(ctx context.Context, userID string) context.Context {
	ctx = context.WithValue(ctx, ContextKeyUserID{}, userID)
	return ctx
}

type ContextKeyRequestID struct{}

type ContextKeyUserRole struct{}
