package handler

import (
	"net/http"

	"inwinstack/cgmh/apiserver/pkg/version"

	"github.com/gin-gonic/gin"
)

func GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"version": version.GetVersion(),
	})
}

func GetHealthz(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
