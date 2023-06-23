package orderdomain

import (
	"errors"
	"fmt"
	"time"

	orderitemdomain "github.com/ddd-db-tx/internal/domain/order_item_domain"
	transportdomain "github.com/ddd-db-tx/internal/domain/transport_domain"
)

type OrderType string

const (
	OrderTypeShipCod OrderType = "ship_cod"
)

type Order struct {
	id        string
	typ       OrderType
	transport transportdomain.Transport
	items     []orderitemdomain.OrderItem
	total     int
	createdAt time.Time
}

var (
	ErrTransportTypeInvalid = errors.New("transport type invalid")
)

func (o Order) ID() string {
	return o.id
}

func (o Order) Typ() OrderType {
	return o.typ
}

func (o Order) Transport() transportdomain.Transport {
	return o.transport
}

func (o *Order) SetTransport(t transportdomain.Transport) error {
	if o.typ == OrderTypeShipCod && t.Typ() != transportdomain.TransportTypeShipCOD {
		return fmt.Errorf("ship cod order requires a ship cod transport")
	}

	if o.typ == OrderTypeShipCod && t.Typ() != transportdomain.TransportTypeOnlPayment {
		return fmt.Errorf("onl payment order requires a onl payment transport")
	}

	o.transport = t

	return nil
}

func (o *Order) AddOrderItem(orderItem orderitemdomain.OrderItem) error {
	for _, it := range o.items {
		if it.ProductID() == orderItem.ProductID() {
			return fmt.Errorf("same product ID has been added: %s", it.ProductID())
		}
	}

	o.items = append(o.items, orderItem)

	return nil
}

func (o *Order) Total() int {
	return o.total
}

func New(id string, typ OrderType, transport transportdomain.Transport) Order {
	return Order{
		id:        id,
		typ:       typ,
		transport: transport,
		createdAt: time.Now(),
	}
}
