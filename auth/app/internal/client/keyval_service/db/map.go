package db

import (
	"bitbucket.org/idomteam/idom-api/auth/internal/client/keyval_service"
	"bitbucket.org/idomteam/idom-api/auth/pkg/logging"
	"bitbucket.org/idomteam/idom-api/auth/pkg/mapdb"
	"context"
	"time"
)

var _ keyval_service.Storager = &db{}

type db struct {
	storage *mapdb.MapDB
	logger     logging.Logger
}

func NewStorage(storage *mapdb.MapDB, logger logging.Logger) keyval_service.Storager {
	return &db{
		storage: storage,
		logger:  logger,
	}
}

func (s *db) Add(ctx context.Context, key, val string ) (err error) {
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	s.storage.Data[key] = val
	return nil
}

func (s *db) Get(ctx context.Context, key string ) (rsp string , err error) {
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.storage.Data[key].(string), nil
}
