package main

import (
	"fmt"
	"github.com/rmr-rodrigues/Products_Age_Exercise_GO/internal"
	"os"
	"time"
)

func main() {
	// Application arguments example: "2018-01-01 00:00:00" "2019-01-01 00:00:00" "(1-3, 4-6, 7-12, >12)"

	const datePattern string = "2006-01-02 15:04:05"
	args := os.Args[1:]
	defaultIntervals := "(1-3, 4-6, 7-12, >12)"

	fmt.Println("")
	if len(args) >= 2 && len(args) <= 3 {
		if len(args) == 3 {
			defaultIntervals = args[2]
		}
		startDate := args[0]
		endDate := args[1]

		var startDateTime, er1 = time.Parse(datePattern, startDate)
		var endDateTime, er2 = time.Parse(datePattern, endDate)

		if er1 == nil && er2 == nil {
			orders := internal.FilterOrders(internal.OrdersList, startDateTime, endDateTime)
			if len(orders) > 0 {
				allIntervals := internal.ParseStringToIntervals(defaultIntervals)
				ages := internal.GetOrdersMap(orders)
				finalMap := internal.SumIntervalsValues(allIntervals, ages)
				if len(finalMap) > 0 {
					fmt.Println(internal.FinalMapToString(finalMap))
				} else {
					fmt.Println("There are no old products being sold!")
				}
			} else {
				fmt.Println("There are no orders in the interval of dates!")
			}
		} else {
			fmt.Println("There seems to be some problem with the dates, try again!")
		}
	} else {
		fmt.Println("There seems to be some problem with the arguments, try again!")
	}
}
