package keyval_service

import (
	"context"
)

type Storager interface {
	Get(ctx context.Context, key string) (string, error)
	Add(ctx context.Context, key, val string) error
}
