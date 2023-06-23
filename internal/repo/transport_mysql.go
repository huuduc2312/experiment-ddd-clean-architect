package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	transportdomain "github.com/ddd-db-tx/internal/domain/transport_domain"
)

type TransportMysqlRepo struct {
	db *sql.DB
}

func NewTransportMysqlRepo(db *sql.DB) transportdomain.RepoInterface {
	return &TransportMysqlRepo{db}
}

type TransportModel struct {
	ID        string
	Name      string
	Type      string
	CreatedAt time.Time
}

func (r *TransportMysqlRepo) FindOne(ctx context.Context, id string) (*transportdomain.Transport, error) {
	var tr TransportModel
	err := r.db.
		QueryRow("SELECT id, name, type FROM transport WHERE id = UUID_TO_BIN(?)", id).
		Scan(&tr.ID, &tr.Name, &tr.Type)
	if err != nil {
		return nil, fmt.Errorf("select transport from mysql: %w", err)
	}

	domainTransport, err := transportdomain.NewTransport(tr.ID, tr.Name, transportdomain.TransportType(tr.Type))
	if err != nil {
		return nil, fmt.Errorf("create new domain transport from db data: %w", err)
	}

	return domainTransport, nil
}
