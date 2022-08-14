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

func DatesDifferenceInMonths(startDate time.Time, endDate time.Time) int {
	datesDifference := endDate.Sub(startDate)

	if datesDifference <= 0 {
		return 0
	} else {
		return int(datesDifference.Hours() / 24 / 30)
	}
}
