package Products_Age

import (
	"time"
)

type Order struct {
	customerName    string
	customerPhone   string
	shippingAddress string
	grandTotal      float64
	placedAt        time.Time
	items           []Item
}

func (order Order) GetPlacedAt() time.Time {
	return order.placedAt
}
