package app

import (
	"lib/tool"
)

type SimplePlatform struct {
	OrderList []Order
	Sys tool.ISystem
} 

func (p *SimplePlatform) Request(order Order) bool{
	p.OrderList = append( p.OrderList, order )
	return true
}

func (p *SimplePlatform) PerformTransaction() []Deal{
	remainList, deals := CheckDeal( p.OrderList )
	solved := CheckSolvedDeal(deals)
	isCanSolve := len(solved) == len(deals)
	if isCanSolve {
		p.OrderList = remainList
		return deals
	}
	return []Deal{}
}

func (p *SimplePlatform) GetRequestWithPlayerKey(key string) []Order{
	return FilterWith( p.OrderList, IsPlayerKey(key) )
}
