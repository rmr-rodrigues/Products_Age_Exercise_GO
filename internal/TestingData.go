package internal

import (
	"time"
)

const DatePattern string = "2006-01-02 15:04:05"

var date20180101, _ = time.Parse(DatePattern, "2018-01-01 00:00:00")
var date20180102, _ = time.Parse(DatePattern, "2018-01-02 00:00:00")
var date20180201, _ = time.Parse(DatePattern, "2018-02-01 00:00:00")
var date20180301, _ = time.Parse(DatePattern, "2018-03-01 00:00:00")
var date20180501, _ = time.Parse(DatePattern, "2018-05-01 00:00:00")
var date20181201, _ = time.Parse(DatePattern, "2018-12-01 00:00:00")
var date20181215, _ = time.Parse(DatePattern, "2018-12-15 00:00:00")
var date20190101, _ = time.Parse(DatePattern, "2019-01-01 00:00:00")
var date20180625, _ = time.Parse(DatePattern, "2018-06-25 00:00:00")
var date20190601, _ = time.Parse(DatePattern, "2019-06-01 00:00:00")

var product1 = Product{Name: "Product1", Category: "Category1", Weight: 2.5, Price: 5.0, CreationDate: date20180102}
var product2 = Product{Name: "Product2", Category: "Category1", Weight: 2.5, Price: 5.0, CreationDate: date20180201}
var product3 = Product{Name: "Product3", Category: "Category1", Weight: 2.5, Price: 5.0, CreationDate: date20180301}
var product4 = Product{Name: "Product4", Category: "Category1", Weight: 2.5, Price: 5.0, CreationDate: date20181201}
var product5 = Product{Name: "Product1", Category: "Category1", Weight: 2.5, Price: 5.0, CreationDate: date20180501}
var product6 = Product{Name: "Product1", Category: "Category1", Weight: 2.5, Price: 5.0, CreationDate: date20181215}
var product7 = Product{Name: "Product1", Category: "Category1", Weight: 2.5, Price: 5.0, CreationDate: date20180101}

var itemDate20180102 = Item{Cost: 10.0, ShippingFee: 5.0, TaxAmount: 1.0, Product: product1}
var itemDate20180201 = Item{Cost: 10.0, ShippingFee: 5.0, TaxAmount: 1.0, Product: product2}
var itemDate20180301 = Item{Cost: 10.0, ShippingFee: 5.0, TaxAmount: 1.0, Product: product3}
var itemDate20181231 = Item{Cost: 10.0, ShippingFee: 5.0, TaxAmount: 1.0, Product: product4}
var itemDate20180501 = Item{Cost: 10.0, ShippingFee: 5.0, TaxAmount: 1.0, Product: product5}
var itemDate20181215 = Item{Cost: 10.0, ShippingFee: 5.0, TaxAmount: 1.0, Product: product6}
var itemDate20180101 = Item{Cost: 10.0, ShippingFee: 5.0, TaxAmount: 1.0, Product: product7}

var Order1 = Order{CustomerName: "Client 1", CustomerPhone: "++351 999999999", ShippingAddress: "Address 1", GrandTotal: 1.0, PlacedAt: date20181201,
	Items: []Item{itemDate20180102, itemDate20180201, itemDate20180301}}
var Order2 = Order{CustomerName: "Client 2", CustomerPhone: "++351 999999998", ShippingAddress: "Address 2", GrandTotal: 2.0, PlacedAt: date20190101,
	Items: []Item{itemDate20180102, itemDate20181231}}
var Order3 = Order{CustomerName: "Client 3", CustomerPhone: "++351 999999998", ShippingAddress: "Address 2", GrandTotal: 2.0, PlacedAt: date20181201,
	Items: []Item{itemDate20180501}}
var Order4 = Order{CustomerName: "Client 1", CustomerPhone: "++351 999999999", ShippingAddress: "Address 1", GrandTotal: 1.0, PlacedAt: date20181201,
	Items: []Item{itemDate20180501, itemDate20180301}}
var Order5 = Order{CustomerName: "Client 2", CustomerPhone: "++351 999999998", ShippingAddress: "Address 2", GrandTotal: 2.0, PlacedAt: date20190101,
	Items: []Item{itemDate20181231, itemDate20180501, itemDate20181231}}
var Order6 = Order{CustomerName: "Client 2", CustomerPhone: "++351 999999998", ShippingAddress: "Address 2", GrandTotal: 2.0, PlacedAt: date20190101,
	Items: []Item{itemDate20181215}}
var Order7 = Order{CustomerName: "Client 2", CustomerPhone: "++351 999999998", ShippingAddress: "Address 2", GrandTotal: 2.0, PlacedAt: date20190101,
	Items: []Item{itemDate20180101, itemDate20180501, itemDate20180102}}
var Order8 = Order{CustomerName: "Client 2", CustomerPhone: "++351 999999998", ShippingAddress: "Address 2", GrandTotal: 2.0, PlacedAt: date20180625,
	Items: []Item{itemDate20180101, itemDate20180501, itemDate20180102}}
var Order9 = Order{CustomerName: "Client 2", CustomerPhone: "++351 999999998", ShippingAddress: "Address 2", GrandTotal: 2.0, PlacedAt: date20190601,
	Items: []Item{itemDate20180101, itemDate20180501, itemDate20180102}}

var OrdersList = []Order{Order1, Order2, Order3, Order4, Order5, Order6, Order7, Order8, Order9}
