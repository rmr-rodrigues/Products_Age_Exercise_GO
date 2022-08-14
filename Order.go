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
