package transportdomain

import (
	"errors"
	"time"
)

type TransportType string

const (
	TransportTypeShipCOD    TransportType = "ship_code"
	TransportTypeOnlPayment TransportType = "onl_payment"
)

type Transport struct {
	id        string
	name      string
	typ       TransportType
	createdAt time.Time
}

var (
	ErrOrderNameMustNotEmpty = errors.New("order name must not empty")
)

func (t Transport) ID() string {
	return t.id
}

func (t Transport) Name() string {
	return t.name
}

func (t *Transport) SetName(name string) error {
	if name == "" {
		return ErrOrderNameMustNotEmpty
	}

	t.name = name

	return nil
}

func (t Transport) Typ() TransportType {
	return t.typ
}

func (t Transport) CreatedAt() time.Time {
	return t.createdAt
}

func NewTransport(id, name string, typ TransportType) (*Transport, error) {
	if name == "" {
		return nil, ErrOrderNameMustNotEmpty
	}

	// validate type

	return &Transport{id: id, name: name, typ: typ, createdAt: time.Now()}, nil
}
