package createorder

import (
	"context"
	"fmt"

	orderdomain "github.com/ddd-db-tx/internal/domain/order_domain"
	transportdomain "github.com/ddd-db-tx/internal/domain/transport_domain"
)

type OrderParam struct {
	Typ         orderdomain.OrderType `json:"type"`
	Items       []OrderItem           `json:"item"`
	TransportID string                `json:"transportID"`
}

type OrderItem struct {
	ProductID string `json:"productID"`
	Qty       int    `json:"qty"`
}

type TransportResp struct {
	ID   string                        `json:"id"`
	Name string                        `json:"name"`
	Typ  transportdomain.TransportType `json:"type"`
}

type OrderResult struct {
}

func Handle(
	ctx context.Context,
	oRepo orderdomain.RepoInterface,
	trRepo transportdomain.RepoInterface,
	orderParam OrderParam,
) (*orderdomain.Order, error) {
	transport, err := trRepo.FindOne(ctx, orderParam.TransportID)
	if err != nil {
		return nil, fmt.Errorf("find transport: %w", err)
	}

	domainOrder, err := oRepo.InsertOne(
		ctx,
		func(id string) orderdomain.Order {
			return orderdomain.New(
				id,
				orderParam.Typ,
				*transport,
			)
		},
	)
	if err != nil {
		return nil, fmt.Errorf("insert order: %w", err)
	}

	return domainOrder, nil
}
