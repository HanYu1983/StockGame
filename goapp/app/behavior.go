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
	OTBuy orderType = iota
	OTSell orderType = iota
)


type Order struct {
	Key string
	PlayerKey string
	Type orderType
	Price float32
	Count int
	Time time.Time
}

func (o Order) IsSolve() bool{
	return o.Count == 0
}

func CanBuy(a Order, b Order) bool {
	return a.Price >= b.Price
}

func Buy(a Order, b Order) (na Order, nb Order, count int) {
	na = a
	nb = b
	if a.Count > b.Count {
		count = b.Count
		na.Count = a.Count - count
		nb.Count = 0
	}else if a.Count < b.Count {
		count = a.Count
		nb.Count = b.Count - count
		na.Count = 0
	}else{
		count = a.Count
		na.Count = 0
		nb.Count = 0
	}
	return
}

type Deal struct {
	Buy Order
	Sell Order
	Count int
	Price float32
	Time time.Time
}

func IsSolve(deal Deal) bool {
	return deal.Buy.Count == 0 && deal.Sell.Count == 0
}

type IPlatform interface {
	Request(order Order) bool
	PerformTransaction() []Deal
}