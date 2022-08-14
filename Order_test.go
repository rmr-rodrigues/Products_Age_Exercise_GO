package Products_Age

import (
	"reflect"
	"testing"
)

func TestGetProductsAges(t *testing.T) {

	t.Run("returns a slice with the ages of all products in the order", func(t *testing.T) {
		want := []int{11, 10, 9}

		got := order1.GetProductsAges()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("returns an empty slice if the products age are less than one month", func(t *testing.T) {
		var want = make([]int, 0)
		got := order6.GetProductsAges()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
