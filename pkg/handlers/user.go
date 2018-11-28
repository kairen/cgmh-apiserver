package handler

import (
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"
	service "inwinstack/cgmh/apiserver/pkg/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *service.DataAccess
}

func (h *UserHandler) Get(c *gin.Context) {
	if !isAdmin(c, h.svc) {
		uuid, err := getUserUUIDByJWT(c)
		if err != nil {
			http.InternalServerError(c, err)
			return
		}

		if uuid != c.Param("uuid") {
			http.Forbidden(c, http.ErrorUserPermission)
			return
		}
	}

	user, err := h.svc.User.FindByUUID(c.Param("uuid"))
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, user)
}

func (h *UserHandler) List(c *gin.Context) {
	if !isAdmin(c, h.svc) {
		http.Forbidden(c, http.ErrorUserPermission)
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
	user := &models.User{}
	if err := c.ShouldBindJSON(&user); err != nil || user.UUID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if !isAdmin(c, h.svc) {
		uuid, err := getUserUUIDByJWT(c)
		if err != nil {
			http.InternalServerError(c, err)
			return
		}

		if uuid != user.UUID {
			http.Forbidden(c, http.ErrorUserPermission)
			return
		}
	}

	if err := h.svc.User.Update(user); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	if !isAdmin(c, h.svc) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

	obj := &struct {
		UUID string `json:"uuid" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&obj); err != nil || obj.UUID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.svc.User.RemoveByUUID(obj.UUID); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}

func (h *UserHandler) UpdateRole(c *gin.Context) {
	if !isAdmin(c, h.svc) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

	role := &models.UserRole{}
	if err := c.ShouldBindJSON(&role); err != nil || role.UserUUID == "" {
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
	if !isAdmin(c, h.svc) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

	stat := &models.UserStatus{}
	if err := c.ShouldBindJSON(&stat); err != nil || stat.UserUUID == "" {
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
	if !isAdmin(c, h.svc) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

	userLevel := &models.UserLevel{}
	if err := c.ShouldBindJSON(&userLevel); err != nil || userLevel.UserUUID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if !h.svc.Level.IsExistByName(userLevel.Name) {
		http.BadRequest(c, http.ErrorResourceNotFound)
		return
	}

	if err := h.svc.User.UpdateLevel(userLevel); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, userLevel)
}
