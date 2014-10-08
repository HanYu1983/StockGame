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

type dealType int

const (
	DTPending dealType = iota
	DTWait dealType = iota
	DTOK dealType = iota
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
	return a.Price >= b.Price
}

func Buy(a Order, b Order) (na Order, nb Order) {
	na = a
	nb = b
	if a.Count > b.Count {
		na.Count = a.Count - b.Count
		nb.Count = 0
		na.DealType = DTWait
		nb.DealType = DTOK
	}else if a.Count < b.Count {
		nb.Count = b.Count - a.Count
		na.Count = 0
		nb.DealType = DTWait
		na.DealType = DTOK
	}else{
		na.Count = 0
		nb.Count = 0
		na.DealType = DTOK
		nb.DealType = DTOK
	}
	return
}

type Deal struct {
	Buy Order
	Sell Order
	Price float32
	Time time.Time
}

func IsSolve(deal Deal) bool {
	return deal.Buy.DealType == DTOK && deal.Sell.DealType == DTOK
}

type IPlatform interface {
	Request(order Order) bool
	PerformTransaction() []Deal
}