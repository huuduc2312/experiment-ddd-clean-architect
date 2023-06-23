package orderdomain

import (
	"context"
)

type RepoInterface interface {
	InsertOne(ctx context.Context, provideEntity func(id string) Order) (*Order, error)
}
