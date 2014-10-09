package app

import (
	"reflect"
	"sort"
)

func CheckSolvedDeal(deals []Deal) []Deal{
	var buf []Deal
	solved := []Deal{}
	for _, deal := range deals {
		buf = append( buf, deal )
		if deal.Buy.IsSolve() && deal.Sell.IsSolve() {
			solved = append( solved, buf... )
			buf = []Deal{}
		}
	}
	return solved
}

func CheckDeal(orders []Order) (remainOrders []Order, deals []Deal){
	buyList := FilterWith( orders, IsOTBuy )
	sellList := FilterWith( orders, IsOTSell )
	sort.Sort(ByPriciple(buyList))
	sort.Sort(ByPriciple(sellList))
	list1, list2, deals := DoDeal( buyList, sellList, []Deal{} )
	return append(list1, list2...), deals
}

func IsOTBuy(order Order) bool{
	return order.Type == OTBuy
}

func IsOTSell(order Order) bool{
	return order.Type == OTSell
}


type ByPriciple []Order
func (a ByPriciple) Len() int{ 
	return len(a) 
}
func (a ByPriciple) Swap(i, j int){
	a[i], a[j] = a[j], a[i] 
}
func (a ByPriciple) Less(i, j int) bool {
	if a[i].Type == OTBuy {
		if a[i].Price == a[j].Price {
			return a[i].Time.Before(a[i].Time)
		}else{
			return a[i].Price > a[j].Price
		}
	}else{
		if a[i].Price == a[j].Price {
			return a[i].Time.Before(a[i].Time)
		}else{
			return a[i].Price < a[j].Price
		}
	}
}


func DoDeal(buyList []Order, sellList []Order, dealList []Deal) (nextBuyList []Order, nextSellList []Order, nextDealList []Deal) {
	nb, ns, nd := DealWith( buyList, sellList, dealList )
	isChanged := [3]int{len(nb), len(ns), len(nd)} != [3]int{len(buyList), len(sellList), len(dealList)}
	if isChanged {
		return DoDeal(nb, ns, nd )
	}else{
		return nb, ns, nd
	}
}

func DealWith(buyList []Order, sellList []Order, dealList []Deal) (nextBuyList []Order, nextSellList []Order, nextDealList []Deal) {
	if len(buyList) == 0 || len(sellList) == 0 {
		return buyList, sellList, dealList
	}
	mostBuy := buyList[0]
	mostSell := sellList[0]
	if CanBuy( mostBuy, mostSell ) {
		nextBuy, nextSell, count := Buy( mostBuy, mostSell )
		deal := Deal{Buy:nextBuy, Sell:nextSell, Price:nextSell.Price, Count: count}
		nextBuyList = func()[]Order{
			if nextBuy.Count == 0 {
				return buyList[1:]
			}else{
				return append([]Order{nextBuy}, buyList[1:]...)
			}
		}()
		nextSellList = func()[]Order{
			if nextSell.Count == 0 {
				return sellList[1:]
			}else{
				return append([]Order{nextSell}, sellList[1:]...)
			}
		}()
		nextDealList = append( dealList, deal )
		return
	}
	return buyList, sellList, dealList
}

func FilterWith(orders []Order, filter func(order Order)bool) []Order {
	var ret []Order
	for _, order := range orders {
		if filter(order) {
			ret = append( ret, order )
		}
	}
	return ret
}

func Map(f interface{}, xs interface{}) interface{} {
	vf := reflect.ValueOf(f)
	vxs := reflect.ValueOf(xs)
	tys := reflect.SliceOf(vf.Type().Out(0))
	vys := reflect.MakeSlice(tys, vxs.Len(), vxs.Len())
	for i := 0; i < vxs.Len(); i++ {
		y := vf.Call([]reflect.Value{vxs.Index(i)})[0]
		vys.Index(i).Set(y)
    }
	return vys.Interface()
}