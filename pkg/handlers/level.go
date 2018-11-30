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

func (h *LevelHandler) Get(c *gin.Context) {
	level, err := h.svc.Level.FindByID(c.Param("id"))
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, level)
}

func (h *LevelHandler) List(c *gin.Context) {
	if !checkAdmin(c, h.svc) {
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
	if !checkAdmin(c, h.svc) {
		return
	}

	level := &model.Level{}
	if err := c.ShouldBindJSON(level); err != nil || !level.Validate() {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.svc.Level.Insert(level); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}

func (h *LevelHandler) Update(c *gin.Context) {
	if !checkAdmin(c, h.svc) {
		return
	}

	level := &model.Level{}
	if err := c.ShouldBindJSON(&level); err != nil || level.ID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.svc.Level.Update(level); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, level)
}

func (h *LevelHandler) Delete(c *gin.Context) {
	if !checkAdmin(c, h.svc) {
		return
	}

	obj := &struct {
		ID string `json:"id" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&obj); err != nil || obj.ID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	users, err := h.svc.User.FindUsersByLevelID(obj.ID)
	if len(users) > 0 || err != nil {
		http.BadRequest(c, http.ErrorResourceRefer)
		return
	}

	if err := h.svc.Level.Remove(obj.ID); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}
