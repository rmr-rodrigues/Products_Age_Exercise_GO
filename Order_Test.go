package Products_Age

import (
	"reflect"
	"testing"
)

func TestGetProductsAges(t *testing.T) {

	t.Run("returns a slice with the ages of all products in the order", func(t *testing.T) {
		want := []int{12, 11, 10}

		got := GetProductsAges(order1)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
