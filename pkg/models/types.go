package models

type Response struct {
	Status int64       `json:"status"`
	Msg    string      `json:"msg"`
	Obj    interface{} `json:"obj"`
}

type ResponsePage struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

func ResponseOK(msg string, obj interface{}) *Response {
	return &Response{
		Status: 200,
		Msg:    msg,
		Obj:    obj,
	}
}

func ResponseError(msg string, obj interface{}) *Response {
	return &Response{
		Status: 500,
		Msg:    msg,
		Obj:    obj,
	}
}
