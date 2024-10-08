package entity

import "time"

type Session struct {
	SessionID string
	SessionData
}

type SessionData struct {
	UserID   string `json:"-"`
	FullName string
	DeviceInfo
	Exp       time.Time
	LastLogin time.Time
}

type DeviceInfo struct {
	UserAgent string
	IPAddress string
}
