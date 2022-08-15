package internal

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func FilterOrders(orders []Order, startDate time.Time, endDate time.Time) []Order {

	var filteredOrders = make([]Order, 0)

	for _, order := range orders {
		if order.PlacedAt.After(startDate) && order.PlacedAt.Before(endDate) {
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
		if interval.Start > interval.End && interval.End == -1 {
			intervalAges := 0
			for _, k := range keys {
				if k >= interval.Start+1 {
					intervalAges = intervalAges + m[k]
				}
			}
			if intervalAges > 0 {
				finalMap[interval] = intervalAges
			}
		} else { // if it's a normal interval
			i := IntervalToSlice(interval.Start, interval.End)
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

func ParseStringToIntervals(s string) []Interval {

	validateIntervalRegex, _ := regexp.Compile("^\\d+-\\d+")
	validateGreaterThanIntervalRegex, _ := regexp.Compile(">\\d+")

	var result []Interval

	if len(s) >= 4 {
		// string examples: "(1-3, 4-6, 7-12, >12)", "(1-3)", "(>1)"
		s1 := strings.TrimSpace(s)
		s2 := strings.TrimLeft(s1, "(")
		s3 := strings.TrimRight(s2, ")")
		s4 := strings.Split(s3, ",")

		for _, value := range s4 {
			s5 := strings.TrimSpace(value)

			if validateIntervalRegex.MatchString(s5) {
				result = append(result, stringToNormalInterval(s5))
			} else if validateGreaterThanIntervalRegex.MatchString(s) {
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

func GetOrdersMap(orders []Order) map[int]int {
	var defaultMap = make(map[int]int)
	for _, order := range orders {
		m := order.GetMapAges()
		defaultMap = addMaps(defaultMap, m)
	}
	return defaultMap
}

func addMaps(map1 map[int]int, map2 map[int]int) map[int]int {
	if len(map1) == 0 {
		return map2
	} else if len(map2) == 0 {
		return map1
	} else {
		for key2, value2 := range map2 {
			if value1, err1 := map1[key2]; err1 {
				map1[key2] = value1 + value2
			} else {
				map1[key2] = value2
			}
		}
		return map1
	}
}

func FinalMapToString(m map[Interval]int) string {
	result := ""
	sortedMapKeys := sortMapKeys(m)
	for _, key := range sortedMapKeys {
		line := ""
		result += line + key.toString() + " months: " + strconv.Itoa(m[key]) + "\n"
	}
	return result
}

func sortMapKeys(m map[Interval]int) []Interval {
	keys := getKeysFromIntervalMap(m)

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i].Start < keys[j].Start
	})

	return keys
}

func getKeysFromIntervalMap(m map[Interval]int) []Interval {
	keys := make([]Interval, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
