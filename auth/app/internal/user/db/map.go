package db

import (
	"bitbucket.org/idomteam/idom-api/auth/internal/user"
	"bitbucket.org/idomteam/idom-api/auth/pkg/logging"
	"bitbucket.org/idomteam/idom-api/auth/pkg/mapdb"
	"context"
	"fmt"
	"strconv"
	"time"
)

var _ user.Storager = &db{}

type db struct {
	storage *mapdb.MapDB
	logger     logging.Logger
}

func NewStorage(storage *mapdb.MapDB, logger logging.Logger) user.Storager {
	return &db{
		storage: storage,
		logger:  logger,
	}
}

func (db *db) Create(ctx context.Context, user user.User) (user.User, error) {
	user.UUID = strconv.Itoa(int(time.Now().UnixNano()))
	db.storage.Data[user.Email] = user
	db.storage.Data[user.UUID] = user
	return user, nil
}

func (db *db) FindOne(ctx context.Context, uuid string) (u user.User, err error) {

	if val, ok := db.storage.Data[uuid]; ok {
		return val.(user.User), nil
	}
	return u, fmt.Errorf("not found by uuid")
}

func (db *db) FindByEmail(ctx context.Context, email string) (u user.User, err error) {
	if val, ok := db.storage.Data[email]; ok {
		return val.(user.User), nil
	}
	return u, fmt.Errorf("not found by email")
}

func (db *db) Update(ctx context.Context, user user.User) error {
	if _, ok := db.storage.Data[user.UUID]; ok {
		db.storage.Data[user.Email] = user
		db.storage.Data[user.UUID] = user
		return nil
	}
	return fmt.Errorf("not found by uuid")
}

func (db *db) Delete(ctx context.Context, uuid string) error {
	if val, ok := db.storage.Data[uuid]; ok {
		delete(db.storage.Data, val.(user.User).Email)
		delete(db.storage.Data, val.(user.User).UUID)
		return nil
	}
	return fmt.Errorf("not found by uuid")
}