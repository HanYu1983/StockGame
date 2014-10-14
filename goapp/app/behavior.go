package app

import (
	"time"
)

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

type Deal struct {
	Buy Order
	Sell Order
	Count int
	Price float32
	Time time.Time
}

type IPlatform interface {
	Request(order Order) bool
	PerformTransaction() []Deal
	GetRequestWithPlayerKey(key string) []Order
}

type ITransactionData interface {
	AddData(deals []Deal)
	GetDataWithRange(after time.Time, before time.Time) []Deal
}

type KLine struct {
	Low float32
	High float32
	Start float32
	End float32
	Count int
	Time time.Time
}