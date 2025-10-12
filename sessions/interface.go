package sessions

import "context"

type SessionData struct {
	UserID    string `json:"user_id"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type SessionStore interface {
	Set(ctx context.Context, sessionID string, data SessionData, ttlSeconds int) error
	Get(ctx context.Context, sessionID string) (*SessionData, error)
	Delete(ctx context.Context, sessionID string) error
}

type Manager interface {
	CreateSession(ctx context.Context, userID string, role string) (string, error)
	GetSession(ctx context.Context, userID string) (*SessionData, error)
	DeleteSession(ctx context.Context, userID string) error
}
