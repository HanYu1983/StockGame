package app

import (
	"net/http"
	"lib/tool"
	"time"
)

func TestHandler(){

}

func Test(w http.ResponseWriter, r *http.Request){
	var logger tool.ILogger
	var platform IPlatform
	
	logger = &tool.ConsoleLogger{Writer: w}
	platform = &SimplePlatform{Logger: logger}
	
	orders := []Order{
		Order{Type:OTBuy, Price:3, Count:5, Time: time.Now()},
		Order{Type:OTBuy, Price:2, Count:1, Time: time.Now()},
		
		Order{Type:OTSell, Price:1.2, Count:2, Time: time.Now()},
		Order{Type:OTSell, Price:1.5, Count:2, Time: time.Now()},
		Order{Type:OTSell, Price:2, Count:2, Time: time.Now()},
	}
	
	logger.Log(orders)
	
	for _, order := range orders {
		platform.Request( order )
	}

	deals := platform.PerformTransaction()
	logger.Log(deals)
}