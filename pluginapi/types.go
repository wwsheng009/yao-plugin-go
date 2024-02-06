package pluginapi

import "net/http"

// Response HTTP Response
type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Headers http.Header `json:"headers"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}
