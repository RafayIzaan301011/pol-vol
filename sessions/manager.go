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

func (m *manager) CreateSession(ctx context.Context, userID string, role string) (string, error) {

	sessionID := uuid.NewString()
	data := SessionData{
		UserID:    userID,
		CreatedAt: time.Now().Format(time.RFC3339),
		Role:      role,
	}

	err := m.Store.Set(ctx, sessionID, data, m.Ttl)
	if err != nil {
		return "", err
	}

	return sessionID, nil

}

func (m *manager) GetSession(ctx context.Context, sessionID string) (*SessionData, error) {
	data, err := m.Store.Get(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	return data, nil

}
func (m *manager) DeleteSession(ctx context.Context, sessionID string) error {
	return m.Store.Delete(ctx, sessionID)
}
