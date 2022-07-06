package handler

import (
	"time"

	"github.com/zituocn/gow"
	"github.com/zituocn/gow/lib/logy"
)

type JSONResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Time int         `json:"time"`
	Data interface{} `json:"data,omitempty"`
}

// ServerJSON response json
func ServerJSON(c *gow.Context, statusCode int, args ...interface{}) {
	var (
		err  error
		data interface{}
		msg  string
		code int
	)

	for _, v := range args {
		switch vv := v.(type) {
		case int:
			code = vv
		case string:
			msg = vv
		case error:
			err = vv
		default:
			data = vv
		}
	}
	if err != nil {
		logy.Errorf("%s -> %s", c.Request.URL.String(), err.Error())
	}
	if code == 0 && msg == "" {
		msg = "success"
	}

	resp := &JSONResponse{
		Code: code,
		Msg:  msg,
		Time: int(time.Now().Unix()),
		Data: data,
	}

	c.ServerJSON(statusCode, &resp)
	return
}

// JSON response json
func JSON(c *gow.Context, args ...interface{}) {
	ServerJSON(c, 200, args...)
}
