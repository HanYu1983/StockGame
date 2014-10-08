package hello

import (
    "fmt"
    "net/http"
	"lib/tool"
	"app"
)

func init() {
    FuncFrontControl := func(w http.ResponseWriter, r *http.Request){
        tool.FrontControl(w, r,
            tool.ActionMap{
				
            },
        )
    }
    http.HandleFunc("/", handler)
    http.HandleFunc("/Func", FuncFrontControl)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
	app.Test(w, r)
}