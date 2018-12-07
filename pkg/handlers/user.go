package handler

import (
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/ldap"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc  *service.DataAccess
	ldap *ldap.LDAP
}

func (h *UserHandler) Get(c *gin.Context) {
	if !checkUserUUID(c, h.svc, c.Param("uuid")) {
		return
	}

	user, err := h.svc.User.FindByUUID(c.Param("uuid"))
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, user)
}

func (h *UserHandler) List(c *gin.Context) {
	if !checkAdmin(c, h.svc) {
		return
	}

	users, err := h.svc.User.FindAll()
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, users)
}

func (h *UserHandler) Update(c *gin.Context) {
	user := &model.UserPost{}
	if err := c.ShouldBindJSON(&user); err != nil || !user.Validate() {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if !checkUserUUID(c, h.svc, user.UUID) {
		return
	}

	if err := h.svc.User.Update(user); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, user)
}

func (h *UserHandler) UpdateRole(c *gin.Context) {
	if !checkAdmin(c, h.svc) {
		return
	}

	role := &model.UserRole{}
	if err := c.ShouldBindJSON(&role); err != nil || !role.Validate() {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.svc.User.UpdateRole(role); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, role)
}

func (h *UserHandler) UpdateStatus(c *gin.Context) {
	if !checkAdmin(c, h.svc) {
		return
	}

	stat := &model.UserStatus{}
	if err := c.ShouldBindJSON(&stat); err != nil {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.svc.User.UpdateStatus(stat); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, stat)
}

func (h *UserHandler) UpdateLevel(c *gin.Context) {
	if !checkAdmin(c, h.svc) {
		return
	}

	userLevel := &model.UserLevel{}
	if err := c.ShouldBindJSON(&userLevel); err != nil || !userLevel.Validate() {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if !h.svc.Level.IsExist(userLevel.LevelID) {
		http.BadRequest(c, http.ErrorResourceNotFound)
		return
	}

	if err := h.svc.User.UpdateLevel(userLevel); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, userLevel)
}

func (h *UserHandler) UpdatePoint(c *gin.Context) {
	if !checkAdmin(c, h.svc) {
		return
	}

	point := &model.Point{}
	if err := c.ShouldBindJSON(&point); err != nil || !point.Validate() {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	user, err := h.svc.User.FindByUUID(point.UserUUID)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}

	value := user.Point + point.Value
	if value < 0 {
		http.BadRequest(c, http.ErrorDeposit)
		return
	}

	if err := h.svc.User.UpdatePoint(point, value); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, point)
}

func (h *UserHandler) Delete(c *gin.Context) {
	if !checkAdmin(c, h.svc) {
		return
	}

	obj := &struct {
		UUID string `json:"uuid" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&obj); err != nil || obj.UUID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.svc.User.Remove(obj.UUID); err != nil {
		http.InternalServerError(c, err)
		return
	}

	if err := h.ldap.DelUser(obj.UUID); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}
