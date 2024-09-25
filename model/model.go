package model

type Request struct {
	Array  []int `json:"numbers"`
	Target int   `json:"Target"`
}
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Resp interface{} `json:"response,omitempty"`
}
