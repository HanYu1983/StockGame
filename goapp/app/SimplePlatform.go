package app

import (
	"reflect"
)

type SimplePlatform struct {
	OrderList []Order
} 

func (p *SimplePlatform) Request(order Order) bool{
	p.OrderList = append( p.OrderList, order )
	return true
}

func (p *SimplePlatform) PerformTransaction() []Deal{
	var dealList []Deal
	var buyList []Order
	var sellList []Order
	/*for {
		nd, nb, ns := DealWith( buyList, sellList, dealList )
		if nd, nd, ns == buyList, sellList, dealList {
		
		}
	}*/
	return nil
}

func DealWith(buyList []Order, sellList []Order, dealList []Deal) (nextBuyList []Order, nextSellList []Order, nextDealList []Deal) {
	if len(buyList) == 0 || len(sellList) == 0 {
		return buyList, sellList, dealList
	}
	mostBuy := buyList[0]
	mostSell := sellList[0]
	if CanBuy( mostBuy, mostSell ) {
		nextBuy, nextSell := Buy( mostBuy, mostSell )
		deal := Deal{Buy:nextBuy, Sell:nextSell, Price:nextSell.Price}
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
		nextDealList = append( nextDealList, deal )
		return
	}
	return buyList, sellList, dealList
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