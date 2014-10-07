package app

import (
	"time"
)

type IPlayer interface {
	GetKey() string
	GetStockList() []IStock
}

type INews interface {

}

type ICompany interface {
	
}

type IStock interface {
	GetPrice() int
	GetCount() int
	GetCompany() ICompany
}

type orderType int

const (
	buy orderType = iota
	sell orderType = iota
)

type dealType int

const (
	pending dealType = iota
	wait dealType = iota
	ok dealType = iota
)

type Order struct {
	PlayerKey string
	Type orderType
	Price float32
	Count int
	Time time.Time
	DealType dealType
}

func CanBuy(a Order, b Order) bool {
	return a.Price > b.Price
}

func Buy(a Order, b Order) (na Order, nb Order) {
	na = a
	nb = b
	if a.Count > b.Count {
		na.Count = a.Count - b.Count
		nb.Count = 0
		na.DealType = wait
		nb.DealType = ok
	}else{
		nb.Count = b.Count - a.Count
		na.Count = 0
		nb.DealType = wait
		na.DealType = ok
	}
	return
}

type Deal struct {
	Buy Order
	Sell Order
	Price float32
	Time time.Time
}

type IPlatform interface {
	Request(order Order) bool
	PerformTransaction() []Deal
}