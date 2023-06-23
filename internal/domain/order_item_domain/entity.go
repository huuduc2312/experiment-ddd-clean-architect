package orderitemdomain

import "errors"

type OrderItem struct {
	id        string
	productID string
	qty       int
}

func (oi OrderItem) ID() string {
	return oi.id
}

func (oi OrderItem) ProductID() string {
	return oi.productID
}

func (oi OrderItem) Qty() int {
	return oi.qty
}

var (
	ErrOrderItemQtyNotGreaterThanZero = errors.New("order item qty not greater than 0")
	ErrOrderItemIdIsNull              = errors.New("order item ID is null")
	ErrOrderItemProdIdIsNull          = errors.New("order item product ID is null")
)

func NewOrderItem(id, productID string, qty int) (*OrderItem, error) {
	switch {
	case id == "":
		return nil, ErrOrderItemIdIsNull
	case qty <= 0:
		return nil, ErrOrderItemQtyNotGreaterThanZero
	}

	return &OrderItem{id: id, productID: productID, qty: qty}, nil
}
