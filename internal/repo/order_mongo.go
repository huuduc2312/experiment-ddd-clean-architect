package repo

import (
	"context"
	"fmt"
	"time"

	orderdomain "github.com/ddd-db-tx/internal/domain/order_domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderMongoRepo struct {
	db *mongo.Database
}

func NewOrderMongoRepo(db *mongo.Database) orderdomain.RepoInterface {
	return &OrderMongoRepo{db}
}

type OrderModel struct {
	ID        primitive.ObjectID    `bson:"_id"`
	Type      orderdomain.OrderType `bson:"type"`
	Transport TransportMongoModel   `bson:"transport"`
	Items     []OrderItemModel      `bson:"items"`
	Total     int                   `bson:"total"`
	CreatedAt time.Time             `bson:"createdAt"`
}

type TransportMongoModel struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Type string             `bson:"type"`
}

type OrderItemModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	ProductID string             `bson:"productID"`
	Qty       int                `bson:"qty"`
}

func (r *OrderMongoRepo) InsertOne(ctx context.Context, provideEntity func(id string) orderdomain.Order) (*orderdomain.Order, error) {
	id := primitive.NewObjectID()
	domainOrder := provideEntity(id.Hex())

	if _, err := r.db.Collection("orders").InsertOne(
		ctx,
		OrderModel{
			ID:   id,
			Type: domainOrder.Typ(),
			Transport: TransportMongoModel{
				Name: domainOrder.Transport().Name(),
				Type: string(domainOrder.Transport().Typ()),
			},
			Items:     []OrderItemModel{},
			Total:     0,
			CreatedAt: time.Now(),
		},
	); err != nil {
		return nil, fmt.Errorf("insert order aggregate to mongodb: %w", err)
	}

	return &domainOrder, nil
}
