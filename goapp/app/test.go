package app

import (
	"lib/tool"
	"time"
)

func TestGetRequest(sys tool.ISystem)interface{}{
	platform := IPlatform(&SimplePlatform{Sys:sys})
	
	platform.Request(Order{Key: "b0", PlayerKey:"han", Type:OTBuy, Price:3, Count:5, Time: time.Now()})
	platform.Request(Order{Key: "b1", PlayerKey:"vic", Type:OTBuy, Price:2, Count:1, Time: time.Now()})
	platform.Request(Order{Key: "b1", PlayerKey:"han", Type:OTBuy, Price:2, Count:1, Time: time.Now()})
	
	sys.Log(platform.GetRequestWithPlayerKey("han"))
	
	return tool.CustomView
}

func TestPlatform(sys tool.ISystem)interface{}{
	platform := IPlatform(&SimplePlatform{Sys:sys})
	
	platform.Request(Order{Key: "b0", Type:OTBuy, Price:3, Count:5, Time: time.Now()})
	platform.Request(Order{Key: "b1", Type:OTBuy, Price:2, Count:1, Time: time.Now()})
	platform.Request(Order{Key: "b2", Type:OTBuy, Price:2, Count:4, Time: time.Now()})
	platform.Request(Order{Key: "s2", Type:OTSell, Price:2, Count:2, Time: time.Now()})
	platform.Request(Order{Key: "s3", Type:OTSell, Price:1.5, Count:5, Time: time.Now()})
	platform.Request(Order{Key: "s4", Type:OTSell, Price:1.2, Count:2, Time: time.Now()})
	
	deals := platform.PerformTransaction()
	sys.Log("can not solve")
	sys.Log(deals)
	sys.Log("add last order")
	
	platform.Request(Order{Key: "s5", Type:OTSell, Price:1, Count:1, Time: time.Now()})
	deals = platform.PerformTransaction()
	sys.Log(deals)
	sys.Log(platform)
	
	return tool.CustomView
}

func TestCheckDeal(sys tool.ISystem)interface{}{
	orders := []Order{
		Order{Key: "b0", Type:OTBuy, Price:3, Count:5, Time: time.Now()},
		Order{Key: "b1", Type:OTBuy, Price:2, Count:1, Time: time.Now()},
		Order{Key: "b2", Type:OTBuy, Price:2, Count:1, Time: time.Now()},
		
		Order{Key: "s2", Type:OTSell, Price:2, Count:2, Time: time.Now()},
		Order{Key: "s3", Type:OTSell, Price:1.5, Count:2, Time: time.Now()},
		Order{Key: "s4", Type:OTSell, Price:1.2, Count:2, Time: time.Now()},
		Order{Key: "s5", Type:OTSell, Price:2, Count:3, Time: time.Now()},
	}
	sys.Log("orders")
	sys.Log(orders)
	
	remainList, deals := CheckDeal( orders )
	sys.Log("remainList")
	sys.Log(remainList)
	
	sys.Log("deals")
	sys.Log(deals)
	
	solved := CheckSolvedDeal(deals)
	sys.Log("solved")
	sys.Log(solved)
	return tool.CustomView
}

func TestEventManager(sys tool.ISystem)interface{}{
	evtMgr := tool.IEventManager(&tool.SimpleEventManager{})
	
	playerMgr := 1040
	
	evtMgr.AddSender("playerManager", func(mgr tool.IEventManager){
		sys.Log("send player manager")
		mgr.Send("CreatePlayerManager", playerMgr)
	})
	evtMgr.SendAgain()
	
	evtMgr.AddReceiver("CreatePlayerManager", func(playerMgr int){
		sys.Log("receiver player manager")
		sys.Log(playerMgr)
	})
	evtMgr.SendAgain()
	
	return tool.CustomView
}

