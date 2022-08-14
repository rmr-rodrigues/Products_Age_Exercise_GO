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

func GetProductsAges(order Order) []int {
	var ages = make([]int, 0)
	for _, item := range order.items {
		months := DatesDifferenceInMonths(item.Product.creationDate, order.placedAt)
		if months > 0 {
			ages = append(ages, months)
		}
	}
	return ages
}
