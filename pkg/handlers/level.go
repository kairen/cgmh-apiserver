package handler

import (
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"
	service "inwinstack/cgmh/apiserver/pkg/services"

	"github.com/gin-gonic/gin"
)

type LevelHandler struct {
	svc *service.DataAccess
}

func (h *LevelHandler) List(c *gin.Context) {
	if !isAdmin(c, h.svc) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

	levels, err := h.svc.Level.FindAll()
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, levels)
}

func (h *LevelHandler) Create(c *gin.Context) {
	if !isAdmin(c, h.svc) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

	level := &model.Level{}
	if err := c.ShouldBindJSON(level); err != nil || !level.Validate() {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if h.svc.Level.IsExistByName(level.Name) {
		http.BadRequest(c, http.ErrorResourceExist)
		return
	}

	if err := h.svc.Level.Insert(level); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}

func (h *LevelHandler) Update(c *gin.Context) {
	if !isAdmin(c, h.svc) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

	level := &model.Level{}
	if err := c.ShouldBindJSON(&level); err != nil || level.ID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if h.svc.Level.IsExistByName(level.Name) {
		http.BadRequest(c, http.ErrorResourceExist)
		return
	}

	oldLevel, err := h.svc.Level.FindByID(level.ID.Hex())
	if err != nil {
		http.BadRequest(c, http.ErrorResourceNotFound)
		return
	}

	if err := h.svc.Level.Update(level); err != nil {
		http.InternalServerError(c, err)
		return
	}

	if err := h.svc.User.UpdateLevelsByName(oldLevel.Name, level.Name); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, level)
}

func (h *LevelHandler) Delete(c *gin.Context) {
	if !isAdmin(c, h.svc) {
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

	level, err := h.svc.Level.FindByID(obj.ID)
	if err != nil {
		http.BadRequest(c, http.ErrorResourceNotFound)
		return
	}

	userLevels, _ := h.svc.User.FindUserLevels(level.Name)
	if len(userLevels) > 0 {
		http.BadRequest(c, http.ErrorResourceRefer)
		return
	}

	if err := h.svc.Level.Remove(obj.ID); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}
