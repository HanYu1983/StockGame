package tool

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type ConsoleLogger struct {
	Writer http.ResponseWriter
}

func (l *ConsoleLogger) Log(msg interface{}) {
	str, _ := json.Marshal(msg)
	fmt.Fprintf(l.Writer, "%s\n", str)
}