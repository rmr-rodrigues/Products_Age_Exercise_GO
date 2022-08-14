package Products_Age

import (
	"reflect"
	"testing"
	"time"
)

// TestFilterOrders calls FilterOrders with a list of valid orders and valid dates
func TestFilterOrders(t *testing.T) {
	t1, _ := time.Parse(datePattern, "2018-01-01 00:00:00")
	t2, _ := time.Parse(datePattern, "2019-01-01 00:00:00")

	t.Run("filters 3 orders out of 5", func(t *testing.T) {
		ordersList := []Order{order1, order2, order3, order4, order5}

		want := []Order{order1, order3, order4}

		got := FilterOrders(ordersList, t1, t2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
