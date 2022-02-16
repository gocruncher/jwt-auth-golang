package keyval_service

import (
	"bitbucket.org/idomteam/idom-api/auth/internal/apperror"
	"bitbucket.org/idomteam/idom-api/auth/pkg/logging"
	"context"
	"errors"
	"fmt"
)

var _ Service = &service{}

type service struct {
	storage Storager
	logger  logging.Logger
}

func NewService(noteStorage Storager, logger logging.Logger) Service {
	return &service{
		storage: noteStorage,
		logger:  logger,
	}
}

type Service interface {
	Test(ctx context.Context) string
	GetKey(ctx context.Context, key string) (string, error)
	SetKey(ctx context.Context, key string, val string)  error
}

func (s service) Test(ctx context.Context) string {
	return "ok"
}

func (s service) GetKey(ctx context.Context, key string) (string, error) {
	val, err := s.storage.Get(ctx, key)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return "", err
		}
		return "", fmt.Errorf("failed to get key. error: %w", err)
	}
	return val, nil
}

func (s service) SetKey(ctx context.Context, key, val string)  error {
	err := s.storage.Add(ctx, key, val)
	if err != nil {
		return  fmt.Errorf("failed to get key. error: %w", err)
	}
	return nil
}

