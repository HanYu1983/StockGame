package app

import (
	"lib/tool"
	"time"
)

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