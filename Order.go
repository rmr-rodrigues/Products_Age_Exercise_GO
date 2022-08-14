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

func (order Order) GetProductsAges() []int {
	var ages = make([]int, 0)
	for _, item := range order.items {
		months := DatesDifferenceInMonths(item.Product.creationDate, order.placedAt)
		if months > 0 {
			ages = append(ages, months)
		}
	}
	return ages
}

func (order Order) GetMapAges() map[int]int {
	resultMap := make(map[int]int)

	for _, productAge := range order.GetProductsAges() {

		if value, found := resultMap[productAge]; found {
			resultMap[productAge] = value + 1
		} else if !found {
			resultMap[productAge] = 1
		}
	}
	return resultMap
}
