package sessions

import "context"

type SessionStore interface {
	Set(ctx context.Context, sessionID string, data map[string]string, ttlSeconds int) error
	Get(ctx context.Context, sessionID string) (map[string]string, error)
	Delete(ctx context.Context, sessionID string) error
}

type Manager interface {
	CreateSession(ctx context.Context, userID string) (string, error)
	GetSession(ctx context.Context, userID string) (map[string]string, error)
	DeleteSession(ctx context.Context, userID string) error
}
