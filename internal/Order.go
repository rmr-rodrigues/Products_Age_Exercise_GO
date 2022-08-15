package internal

import (
	"time"
)

type Order struct {
	CustomerName    string
	CustomerPhone   string
	ShippingAddress string
	GrandTotal      float64
	PlacedAt        time.Time
	Items           []Item
}

func (order Order) GetPlacedAt() time.Time {
	return order.PlacedAt
}

func (order Order) GetProductsAges() []int {
	var ages = make([]int, 0)
	for _, item := range order.Items {
		months := DatesDifferenceInMonths(item.Product.CreationDate, order.PlacedAt)
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
