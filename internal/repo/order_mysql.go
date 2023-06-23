package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	orderdomain "github.com/ddd-db-tx/internal/domain/order_domain"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

type OrderMysqlRepo struct {
	db *sql.DB
}

func NewOrderMysqlRepo(db *sql.DB) orderdomain.RepoInterface {
	return &OrderMysqlRepo{db}
}

func (r *OrderMysqlRepo) InsertOne(ctx context.Context, provideEntity func(id string) orderdomain.Order) (*orderdomain.Order, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	orderID := strings.Replace(uuid.NewString(), "-", "", -1)
	domainOrder := provideEntity(orderID)

	err = execQueryOrRollback(
		tx,
		"INSERT INTO `order` (id, type, total) VALUES (UUID_TO_BIN(?), ?, ?)",
		orderID,
		domainOrder.Typ(),
		domainOrder.Total(),
	)
	if err != nil {
		return nil, fmt.Errorf("insert order to mysql: %w", err)
	}

	tr := domainOrder.Transport()
	err = execQueryOrRollback(
		tx,
		`INSERT INTO order_transport (order_id, name, type) VALUES (UUID_TO_BIN(?), ?, ?)`,
		orderID,
		tr.Name(),
		tr.Typ(),
	)
	if err != nil {
		return nil, fmt.Errorf("insert order transport to mysql: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit transaction: %w", err)
	}

	return &domainOrder, nil
}

func execQueryOrRollback(tx *sql.Tx, q string, args ...any) error {
	var errs error

	_, err := tx.Exec(q, args...)

	if err != nil {
		if rollbackErr := tx.Rollback(); err != nil {
			errs = multierror.Append(errs, rollbackErr)
		}

		errs = multierror.Append(err)

		return errs
	}

	return nil
}
