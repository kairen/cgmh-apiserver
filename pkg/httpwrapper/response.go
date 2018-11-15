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

func Success(c *gin.Context, payload interface{}) {
	JSON(c, StatusSuccess, payload, nil)
}

func NotFound(c *gin.Context, err error) {
	JSON(c, StatusNotFound, nil, err)
}

func InternalServerError(c *gin.Context, err error) {
	JSON(c, StatusInternalServerError, nil, err)
}

func BadRequest(c *gin.Context, err error) {
	JSON(c, StatusBadRequest, nil, err)
}

func Forbidden(c *gin.Context, err error) {
	JSON(c, StatusForbidden, nil, err)
}

func Unauthorized(c *gin.Context, err error) {
	JSON(c, StatusUnauthorized, nil, err)
}
