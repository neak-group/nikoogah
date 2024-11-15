package session

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/neak-group/nikoogah/internal/infra/keystorefx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type SessionService struct {
	keyStore keystorefx.KeyStoreConn
	logger   *zap.Logger
}

type SessionServiceParams struct {
	fx.In

	Keystore keystorefx.KeyStoreConn
	Logger   *zap.Logger
}

func ProvideSessionService(p SessionServiceParams) *SessionService {
	return &SessionService{
		keyStore: p.Keystore,
		logger:   p.Logger,
	}
}

type Session struct {
	SessionID string
	SessionData
}

type SessionData struct {
	UserID      string `json:"-"`
	FullName    string
	PhoneNumber string
	UserState   string
	DeviceInfo
	Exp       time.Time
	LastLogin time.Time
}

type DeviceInfo struct {
	UserAgent string
	IPAddress string
}

const sessionKeyPrefix string = "sess"

func (ss *SessionService) NewSession(ctx context.Context, userID string, fullName string, deviceInfo DeviceInfo) (sessionID *Session, err error) {
	rc, err := ss.keyStore.KSClient(ctx)
	if err != nil {
		return nil, err
	}

	token, err := generateToken(32) // Generates a 256-bit token
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	sessionKey := fmt.Sprintf("%s:%s", sessionKeyPrefix, token)
	expiry := time.Now().Add(30 * 24 * time.Hour)
	loginTime := time.Now()
	_, err = rc.HSet(ctx, sessionKey, map[string]interface{}{
		"UserID":    userID,
		"FullName":  fullName,
		"UserAgent": deviceInfo.UserAgent,
		"IPAddress": deviceInfo.IPAddress,
		"Exp":       expiry.Format(time.RFC3339),
		"LastLogin": loginTime.Format(time.RFC3339),
	}).Result()
	if err != nil {
		return nil, err
	}

	err = rc.ExpireAt(ctx, sessionKey, expiry).Err()
	if err != nil {
		return nil, err
	}

	return &Session{
		SessionID: token,
		SessionData: SessionData{
			UserID:     userID,
			FullName:   fullName,
			DeviceInfo: deviceInfo,
			Exp:        expiry,
			LastLogin:  loginTime,
		},
	}, nil
}

func (ss *SessionService) ValidateSession(ctx context.Context, sessionID string) (session *Session, err error) {
	rc, err := ss.keyStore.KSClient(ctx)
	if err != nil {
		return nil, err
	}

	sessionKey := fmt.Sprintf("%s:%s", sessionKeyPrefix, sessionID)
	// Retrieve session data
	data, err := rc.HGetAll(ctx, sessionKey).Result()
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		ss.logger.Info("session data retrieved from redis", zap.Any("data", data), zap.String("sessKey", sessionKey))
		return nil, fmt.Errorf("session not found")
	}

	// Parse retrieved data into session struct
	exp, err := time.Parse(time.RFC3339, data["Exp"])
	if err != nil {
		return nil, err
	}
	lastLogin, err := time.Parse(time.RFC3339, data["LastLogin"])
	if err != nil {
		return nil, err
	}

	userID, ok := data["UserID"]
	if !ok {
		return nil, fmt.Errorf("user id is not valid")
	}

	session = &Session{
		SessionID: sessionID,
		SessionData: SessionData{
			UserID:   userID,
			FullName: data["UserName"],
			DeviceInfo: DeviceInfo{
				UserAgent: data["UserAgent"],
				IPAddress: data["IPAddress"],
			},
			Exp:       exp,
			LastLogin: lastLogin,
		},
	}

	return session, nil
}

func generateToken(byteLength int) (string, error) {
	b := make([]byte, byteLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", err // Properly handle the error
	}
	// Use URL-safe encoding with Base64
	return base64.URLEncoding.EncodeToString(b), nil
}

func (ss *SessionService) Nullify(ctx context.Context, sessionID string) error {
	rc, err := ss.keyStore.KSClient(ctx)
	if err != nil {
		return err
	}

	sessionKey := fmt.Sprintf("%s:%s", sessionKeyPrefix, sessionID)

	err = rc.Del(ctx, sessionKey).Err()
	if err != nil {
		return err
	}

	return nil
}
