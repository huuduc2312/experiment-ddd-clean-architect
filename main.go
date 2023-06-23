package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	orderdomain "github.com/ddd-db-tx/internal/domain/order_domain"
	"github.com/ddd-db-tx/internal/repo"
	createorder "github.com/ddd-db-tx/internal/usecase/create_order"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ddd")
	if err != nil {
		log.Fatal(err)
	}

	// oRepo := repo.NewOrderMysqlRepo(db)
	trRepo := repo.NewTransportMysqlRepo(db)

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	mgDB := mongoClient.Database("ddd")

	oMongoRepo := repo.NewOrderMongoRepo(mgDB)

	domainOrder, err := createorder.Handle(ctx, oMongoRepo, trRepo, createorder.OrderParam{
		Typ:         orderdomain.OrderTypeShipCod,
		Items:       []createorder.OrderItem{},
		TransportID: "4fe1c8ac113b11eea65f00155d357519",
	})

	fmt.Println("Res", domainOrder)
	fmt.Println("Err", err)
}
