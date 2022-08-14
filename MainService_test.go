package Products_Age

import (
	"testing"
	"time"
)

const datePattern string = "2006-01-02 15:04:05"

// TestFilterOrders calls FilterOrders with a list of valid orders and valid dates
func TestFilterOrders(t *testing.T) {
	t1, _ := time.Parse(datePattern, "2022-08-13 00:00:00")
	t2, _ := time.Parse(datePattern, "2021-08-30 00:00:00")

}
