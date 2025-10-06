package sessions

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type manager struct {
	Store SessionStore
	Ttl   int
}

func NewSessionStore(store SessionStore, ttl int) Manager {
	return &manager{
		Store: store,
	}
}

func (m *manager) CreateSession(ctx context.Context, userID string) (string, error) {

	sessionID := uuid.NewString()
	data := map[string]string{
		"user_id":    userID,
		"created_at": time.Now().Format(time.RFC3339),
	}

	err := m.Store.Set(ctx, sessionID, data, m.Ttl)
	if err != nil {
		return "", err
	}

	return sessionID, nil

}

func (m *manager) GetSession(ctx context.Context, sessionID string) (map[string]string, error) {
	data, err := m.Store.Get(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	return data, nil

}
func (m *manager) DeleteSession(ctx context.Context, sessionID string) error {
	return m.Store.Delete(ctx, sessionID)
}
