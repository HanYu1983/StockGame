package app

import (
	"lib/tool"
)

type SimplePlatform struct {
	OrderList []Order
	Logger tool.ILogger
} 

func (p *SimplePlatform) Request(order Order) bool{
	p.OrderList = append( p.OrderList, order )
	return true
}

func (p *SimplePlatform) PerformTransaction() []Deal{
	buyList := FilterWith( p.OrderList, func(order Order) bool{
		return order.Type == OTBuy
	})
	sellList := FilterWith( p.OrderList, func(order Order) bool{
		return order.Type == OTSell
	})
	_, _, deals := DoDeal( buyList, sellList, []Deal{} )
	return deals
}
