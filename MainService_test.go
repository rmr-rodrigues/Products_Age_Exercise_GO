package Products_Age

import (
	"reflect"
	"testing"
	"time"
)

var t1, _ = time.Parse(datePattern, "2018-01-01 00:00:00")
var t2, _ = time.Parse(datePattern, "2019-01-01 00:00:00")

// TODO review this line TestFilterOrders calls FilterOrders with a list of valid orders and valid dates
func TestFilterOrders(t *testing.T) {

	t.Run("filters 3 orders out of 5", func(t *testing.T) {
		ordersList := []Order{order1, order2, order3, order4, order5}

		want := []Order{order1, order3, order4}

		got := FilterOrders(ordersList, t1, t2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("returns an empty slice it there are no orders in the interval", func(t *testing.T) {
		ordersList := []Order{order5}

		var want []Order
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
