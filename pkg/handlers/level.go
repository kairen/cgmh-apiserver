package handler

import (
	"inwinstack/cgmh/apiserver/pkg/dao"
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/gin-gonic/gin"
)

type LevelHandler struct {
	dao *dao.DataAccess
}

func (h *LevelHandler) List(c *gin.Context) {
	if !isAdmin(c, h.dao) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

	levels, err := h.dao.Level.FindAll()
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, levels)
}

func (h *LevelHandler) Create(c *gin.Context) {
	if !isAdmin(c, h.dao) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

	level := &models.Level{}
	if err := c.ShouldBindJSON(level); err != nil {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if h.dao.Level.IsExistByName(level.Name) {
		http.BadRequest(c, http.ErrorResourceExist)
		return
	}

	if err := h.dao.Level.Insert(level); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}

func (h *LevelHandler) Update(c *gin.Context) {
	if !isAdmin(c, h.dao) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

	level := &models.Level{}
	if err := c.ShouldBindJSON(&level); err != nil || level.ID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if h.dao.Level.IsExistByName(level.Name) {
		http.BadRequest(c, http.ErrorResourceExist)
		return
	}

	if err := h.dao.Level.Update(level); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, level)
}

func (h *LevelHandler) Delete(c *gin.Context) {
	if !isAdmin(c, h.dao) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

	obj := &struct {
		ID string `json:"id" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&obj); err != nil || obj.ID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.dao.Level.Remove(obj.ID); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}
