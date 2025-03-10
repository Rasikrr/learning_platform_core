package session

import (
	"context"
	"errors"
)

var ErrSessionNotFound = errors.New("session not found")

const (
	sessionKey = "_session"
)

func GetFromCtx(ctx context.Context) (*Session, error) {
	ses, ok := ctx.Value(sessionKey).(*Session)
	if !ok {
		return nil, ErrSessionNotFound
	}
	return ses, nil
}
