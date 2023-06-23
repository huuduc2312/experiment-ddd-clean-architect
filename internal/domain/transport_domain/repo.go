package transportdomain

import "context"

type RepoInterface interface {
	FindOne(ctx context.Context, id string) (*Transport, error)
}
