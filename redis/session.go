package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"pvms/sessions"
	"time"

	"github.com/redis/go-redis/v9"
)

type SessionStore struct {
	ClusterClient *redis.ClusterClient
}

func NewSessionStore(client *redis.ClusterClient) *SessionStore {
	return &SessionStore{
		ClusterClient: client,
	}
}

func (ss *SessionStore) Set(ctx context.Context, sessionID string, data sessions.SessionData, ttlSeconds int) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ss.ClusterClient.Set(ctx, sessionID, bytes, time.Duration(ttlSeconds)*time.Second).Err()
}

func (ss *SessionStore) Get(ctx context.Context, sessionID string) (*sessions.SessionData, error) {
	val, err := ss.ClusterClient.Get(ctx, sessionID).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("get session %v", err)
	}

	var data sessions.SessionData
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (ss *SessionStore) Delete(ctx context.Context, sessionID string) error {
	if err := ss.ClusterClient.Del(ctx, sessionID).Err(); err != nil {
		return err
	}

	return nil
}
