package handler

import (
	"inwinstack/cgmh/apiserver/pkg/dao"
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	dao *dao.DataAccess
}

func (h *UserHandler) Get(c *gin.Context) {
	user, err := h.dao.User.FindByUUID(c.Param("uuid"))
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, user)
}

func (h *UserHandler) List(c *gin.Context) {
	users, err := h.dao.User.FindAll()
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, users)
}

func (h *UserHandler) Update(c *gin.Context) {
	user := &models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil || user.UUID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.dao.User.Update(user); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	user := &struct {
		UUID string `json:"uuid" binding:"required"`
	}{}
	err := c.ShouldBindJSON(&user)
	if err != nil || user.UUID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.dao.User.RemoveByUUID(user.UUID); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}
