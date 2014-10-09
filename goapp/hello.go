package hello

import (
    "fmt"
    "net/http"
	"lib/tool"
	"app"
)

func init() {
    TestFrontControl := func(w http.ResponseWriter, r *http.Request){
        tool.FrontControl(w, r,
            tool.ActionMap{
				"TestCheckDeal": app.TestCheckDeal,
            },
        )
    }
    http.HandleFunc("/", handler)
    http.HandleFunc("/Func", TestFrontControl)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}