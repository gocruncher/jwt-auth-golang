package user

import (
	"context"
)

type Storager interface {
	Create(ctx context.Context, user User) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	FindOne(ctx context.Context, uuid string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, uuid string) error
}
