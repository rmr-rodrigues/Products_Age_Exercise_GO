package Products_Age

import (
	"regexp"
	"strconv"
	"strings"
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

func parseStringToIntervals(s string) []Interval {

	validateIntervalRegex, _ := regexp.Compile("^\\d+-\\d+")
	validateGreaterThanIntervalRegex, _ := regexp.Compile(">\\d+") //regexp.Compile("^>\\d+$")

	var result []Interval

	if len(s) >= 4 {
		// string examples: "(1-3, 4-6, 7-12, >12)", "(1-3)", (>1)
		s1 := strings.TrimSpace(s)
		s2 := strings.TrimLeft(s1, "(")
		s3 := strings.TrimRight(s2, ")")
		s4 := strings.Split(s3, ",")

		for _, value := range s4 {
			s5 := strings.TrimSpace(value)

			if validateIntervalRegex.MatchString(s5) {
				result = append(result, stringToNormalInterval(s5))
			} else if validateGreaterThanIntervalRegex.MatchString(s) { // TODO verify the regex
				result = append(result, stringToGreaterThanInterval(s5))
			}

		}
	}
	return result
}

func stringToNormalInterval(s string) Interval {
	s1 := strings.Split(s, "-")

	start, _ := strconv.Atoi(s1[0])
	end, _ := strconv.Atoi(s1[1])
	return Interval{start, end}
}

func stringToGreaterThanInterval(s string) Interval {
	s1 := strings.TrimLeft(s, ">")

	start, _ := strconv.Atoi(s1)
	end := -1
	return Interval{start, end}
}
