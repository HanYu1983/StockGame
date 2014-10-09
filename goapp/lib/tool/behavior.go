package tool

import (
	"net/http"
)

type IRepository interface {
    Create(po interface{}) int64
    Update(key int64, po interface{})
    Read(key int64) interface{}
    Delete(key int64)
    GetAll() []interface{}
}

type ICookieManager interface {
    GetValue()(bool, string)
    SetValue(v string)
    Clear()
}

type ILogger interface {
	Log(msg interface{})
}

type ISystem interface {
	GetRequest() *http.Request
	GetResponse() http.ResponseWriter
	Log(msg interface{})
}