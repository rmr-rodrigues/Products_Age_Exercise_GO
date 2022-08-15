package internal

import (
	"reflect"
	"testing"
)

func TestGetProductsAges(t *testing.T) {

	t.Run("returns a slice with the ages of all products in the order", func(t *testing.T) {
		want := []int{11, 10, 9}

		got := Order1.GetProductsAges()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("returns an empty slice if the products age are less than one month", func(t *testing.T) {
		var want = make([]int, 0)
		got := Order6.GetProductsAges()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestGetMapAges(t *testing.T) {
	t.Run("for a slice of ages return a map with the age as a key and the quantity of ages as the value", func(t *testing.T) {
		want := map[int]int{11: 1, 10: 1, 9: 1}
		got := Order1.GetMapAges()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("for a slice with different products with the same age", func(t *testing.T) {
		want := map[int]int{8: 1, 12: 2}
		got := Order7.GetMapAges()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("for an empty slice of ages return an empty map", func(t *testing.T) {
		want := make(map[int]int)
		got := Order6.GetMapAges()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
