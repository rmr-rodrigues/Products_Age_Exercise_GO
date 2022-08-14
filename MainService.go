package Products_Age

import "time"

func FilterOrders(orders []Order, startDate time.Time, endDate time.Time) []Order {

	var filteredOrders []Order

	for _, order := range orders {
		if order.placedAt.After(startDate) && order.placedAt.Before(endDate) {
			filteredOrders = append(filteredOrders, order)
		}
	}
	return filteredOrders
}
