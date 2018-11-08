package httpwrapper

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message"`
	Reason  string      `json:"reason,omitempty"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(c *gin.Context, code int, payload interface{}, err error) {
	resp := &Response{Code: code, Message: getMsg(code)}
	if payload != nil {
		resp.Data = payload
	}

	if err != nil {
		resp.Reason = err.Error()
	}
	c.JSON(code, resp)
}
