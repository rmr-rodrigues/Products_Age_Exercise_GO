package Products_Age

import (
	"time"
)

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

func IntervalToSlice(start int, end int) []int {
	//list := make([]int, 0)
	var l []int

	if end > start && start > 0 {
		for i := start; i <= end; i++ {
			l = append(l, i)
		}
	}
	return l
}

func SumAgesByInterval(interval []int, agesMap map[int]int) int {
	sum := 0

	for _, value := range interval {
		if v, found := agesMap[value]; found {
			sum = sum + v
		}
	}
	return sum
}

func SumIntervalsValues(intervals []Interval, m map[int]int) map[Interval]int {
	keys := getKeysFromMap(m)
	var finalMap = map[Interval]int{}
	for _, interval := range intervals {
		// if it's an interval greater than
		if interval.start > interval.end && interval.end == -1 {
			intervalAges := 0
			for _, k := range keys {
				if k >= interval.start+1 {
					intervalAges = intervalAges + m[k]
				}
			}
			if intervalAges > 0 {
				finalMap[interval] = intervalAges
			}
		} else { // if it's a normal interval
			i := IntervalToSlice(interval.start, interval.end)
			intervalAges := SumAgesByInterval(i, m)
			if intervalAges > 0 {
				finalMap[interval] = intervalAges
			}
		}
	}
	return finalMap
}

func getKeysFromMap(m map[int]int) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
