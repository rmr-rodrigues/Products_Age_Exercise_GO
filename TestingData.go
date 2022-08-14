package Products_Age

import "time"

const datePattern string = "2006-01-02 15:04:05"

var date20180101, _ = time.Parse(datePattern, "2018-01-01 00:00:00")
var date20180102, _ = time.Parse(datePattern, "2018-01-02 00:00:00")
var date20180201, _ = time.Parse(datePattern, "2018-02-01 00:00:00")
var date20180301, _ = time.Parse(datePattern, "2018-03-01 00:00:00")
var date20180501, _ = time.Parse(datePattern, "2018-05-01 00:00:00")
var date20181201, _ = time.Parse(datePattern, "2018-12-01 00:00:00")
var date20181215, _ = time.Parse(datePattern, "2018-12-15 00:00:00")
var date20190101, _ = time.Parse(datePattern, "2019-01-01 00:00:00")

var product1 = Product{"Product1", "Category1", 2.5, 5.0, date20180102}
var product2 = Product{"Product2", "Category1", 2.5, 5.0, date20180201}
var product3 = Product{"Product3", "Category1", 2.5, 5.0, date20180301}
var product4 = Product{"Product4", "Category1", 2.5, 5.0, date20181201}
var product5 = Product{"Product1", "Category1", 2.5, 5.0, date20180501}
var product6 = Product{"Product1", "Category1", 2.5, 5.0, date20181215}

var itemDate20180102 = Item{10.0, 5.0, 1.0, product1}
var itemDate20180201 = Item{10.0, 5.0, 1.0, product2}
var itemDate20180301 = Item{10.0, 5.0, 1.0, product3}
var itemDate20181231 = Item{10.0, 5.0, 1.0, product4}
var itemDate20180501 = Item{10.0, 5.0, 1.0, product5}
var itemDate20181215 = Item{10.0, 5.0, 1.0, product6}

var order1 = Order{"Client 1", "++351 999999999", "Address 1", 1.0, date20181201,
	[]Item{itemDate20180102, itemDate20180201, itemDate20180301}}
var order2 = Order{"Client 2", "++351 999999998", "Address 2", 2.0, date20190101,
	[]Item{itemDate20180102, itemDate20181231}}
var order3 = Order{"Client 3", "++351 999999998", "Address 2", 2.0, date20181201,
	[]Item{itemDate20180501}}
var order4 = Order{"Client 1", "++351 999999999", "Address 1", 1.0, date20181201,
	[]Item{itemDate20180501, itemDate20180301}}
var order5 = Order{"Client 2", "++351 999999998", "Address 2", 2.0, date20190101,
	[]Item{itemDate20181231, itemDate20180501, itemDate20181231}}
var order6 = Order{"Client 2", "++351 999999998", "Address 2", 2.0, date20190101,
	[]Item{itemDate20181215}}
