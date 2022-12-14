package internal

import (
	"reflect"
	"testing"
	"time"
)

var t1, _ = time.Parse(DatePattern, "2018-01-01 00:00:00")
var t2, _ = time.Parse(DatePattern, "2019-01-01 00:00:00")

func TestFilterOrders(t *testing.T) {

	t.Run("filters 3 orders out of 5", func(t *testing.T) {
		ordersList := []Order{Order1, Order2, Order3, Order4, Order5}

		want := []Order{Order1, Order3, Order4}

		got := FilterOrders(ordersList, t1, t2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("returns an empty slice if there are no orders in the interval", func(t *testing.T) {
		ordersList := []Order{Order5}

		var want = make([]Order, 0)
		got := FilterOrders(ordersList, t1, t2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestDatesDifferenceInMonths(t *testing.T) {
	t.Run("calculates de difference between to sequential dates", func(t *testing.T) {
		want := 12
		got := DatesDifferenceInMonths(t1, t2)

		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("returns 0 fore two equal dates", func(t *testing.T) {
		want := 0
		got := DatesDifferenceInMonths(t2, t2)

		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("returns 0 if startDate bigger than endDate", func(t *testing.T) {
		want := 0
		got := DatesDifferenceInMonths(t2, t1)

		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestIntervalToSlice(t *testing.T) {
	t.Run("for a valid interval returns a slice with all the values in that interval", func(t *testing.T) {

		want := []int{1, 2, 3, 4, 5}
		got := IntervalToSlice(1, 5)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("for an interval with start and end values equals returns an empty slice", func(t *testing.T) {

		var want []int
		got := IntervalToSlice(1, 1)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("for an interval with a start value bigger than the end value returns an empty slice", func(t *testing.T) {

		var want []int
		got := IntervalToSlice(4, 1)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAgesByInterval(t *testing.T) {
	t.Run("for a valid interval and a map with values it returns the sum of all values", func(t *testing.T) {

		m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
		i := []int{1, 2, 3, 4, 5}
		want := 15
		got := SumAgesByInterval(i, m)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("for a valid interval and a map without values it returns 0", func(t *testing.T) {

		m := map[int]int{}
		i := []int{1, 2, 3, 4, 5}
		want := 0
		got := SumAgesByInterval(i, m)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("for a valid interval and a map with values out of the interval it returns 0", func(t *testing.T) {

		m := map[int]int{3: 3, 4: 4, 5: 5}
		i := []int{1, 2}
		want := 0
		got := SumAgesByInterval(i, m)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumIntervalsValues(t *testing.T) {
	t.Run("for a slice of valid intervals and a map with values, it returns a map with the sum of values for each interval and the intervals as keys", func(t *testing.T) {

		m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6}
		intervals := []Interval{{1, 3}, {4, 6}}
		want := map[Interval]int{Interval{Start: 1, End: 3}: 6, Interval{Start: 4, End: 6}: 15}
		got := SumIntervalsValues(intervals, m)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("for an interval greater than, returns a map with the interval as key", func(t *testing.T) {

		m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6}
		intervals := []Interval{{1, -1}}
		want := map[Interval]int{Interval{Start: 1, End: -1}: 20}
		got := SumIntervalsValues(intervals, m)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("for an interval without values it returns an empty map", func(t *testing.T) {

		m := map[int]int{1: 1, 2: 2, 3: 3}
		intervals := []Interval{{4, 6}}
		want := map[Interval]int{}
		got := SumIntervalsValues(intervals, m)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestParseStringToIntervals(t *testing.T) {
	t.Run("for a correct string has to return a slice with all the intervals", func(t *testing.T) {

		// strings examples: "(1-3, 4-6, 7-12, >12)", "(1-3)", (>1)
		intervals := "(1-3, 4-6, 7-12, >12)"
		want := []Interval{{1, 3}, {4, 6}, {7, 12}, {12, -1}}
		got := ParseStringToIntervals(intervals)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("parse just a string with one normal interval", func(t *testing.T) {

		// strings examples: "(1-3, 4-6, 7-12, >12)", "(1-3)", (>1)
		intervals := "(1-3)"
		want := []Interval{{1, 3}}
		got := ParseStringToIntervals(intervals)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("parse just a string with one greater than interval", func(t *testing.T) {

		// strings examples: "(1-3, 4-6, 7-12, >12)", "(1-3)", (>1)
		intervals := "(>3)"
		want := []Interval{{3, -1}}
		got := ParseStringToIntervals(intervals)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestAddMaps(t *testing.T) {

	t.Run("combine 2 maps returning the first map combined with the values of the second map", func(t *testing.T) {
		m1 := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
		m2 := map[int]int{1: 1, 2: 2, 3: 3}

		want := map[int]int{1: 2, 2: 4, 3: 6, 4: 4, 5: 5}

		got := addMaps(m1, m2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("if first map empty, returns the second map", func(t *testing.T) {
		m1 := map[int]int{}
		m2 := map[int]int{1: 1, 2: 2, 3: 3}

		want := map[int]int{1: 1, 2: 2, 3: 3}

		got := addMaps(m1, m2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("if second map empty, returns the first map", func(t *testing.T) {
		m1 := map[int]int{1: 1, 2: 2, 3: 3}
		m2 := map[int]int{}

		want := map[int]int{1: 1, 2: 2, 3: 3}

		got := addMaps(m1, m2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
