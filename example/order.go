package example

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type OrderID uuid.UUID

type Item struct {
	title string
	price decimal.Decimal
	count int
}

type Order struct {
	id         OrderID
	items      []Item
	totalPrice decimal.Decimal
	customerID uuid.UUID
	createdAt  time.Time
	isPaid     bool
}
