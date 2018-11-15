package handler

import (
	"inwinstack/cgmh/apiserver/pkg/dao"
	"inwinstack/cgmh/apiserver/pkg/version"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GlobalHandler struct {
	dao *dao.DataAccess

	// subhandler
	Auth *AuthHandler
	User *UserHandler
	Form *FormHandler
}

func New(dao *dao.DataAccess) *GlobalHandler {
	h := &GlobalHandler{dao: dao}
	h.Auth = &AuthHandler{dao: dao}
	h.User = &UserHandler{dao: dao}
	h.Form = &FormHandler{dao: dao}
	return h
}

func (h *GlobalHandler) GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"version": version.GetVersion(),
	})
}

func (h *GlobalHandler) GetHealthz(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func (h *GlobalHandler) GetDAO() *dao.DataAccess {
	return h.dao
}
