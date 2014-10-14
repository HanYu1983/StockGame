package hello

import (
    "fmt"
    "net/http"
	"lib/tool"
	"app"
)

func init() {
	testAction := tool.ActionMap{
		"TestCheckDeal": app.TestCheckDeal,
		"TestPlatform": app.TestPlatform,
		"TestGetRequest": app.TestGetRequest,
		"TestEventManager": app.TestEventManager,
	}
    http.HandleFunc("/", handler)
    http.HandleFunc("/Test", tool.FrontControllerWith(testAction))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}