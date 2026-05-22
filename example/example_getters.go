package example

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

func (i *Item) Title() string {
	return i.title
}

func (i *Item) Price() decimal.Decimal {
	return i.price
}

func (i *Item) Count() int {
	return i.count
}

func (o *Order) Id() OrderID {
	return o.id
}

func (o *Order) Items() []Item {
	return o.items
}

func (o *Order) TotalPrice() decimal.Decimal {
	return o.totalPrice
}

func (o *Order) CustomerID() uuid.UUID {
	return o.customerID
}

func (o *Order) CreatedAt() time.Time {
	return o.createdAt
}

func (o *Order) IsPaid() bool {
	return o.isPaid
}
