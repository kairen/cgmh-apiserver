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
	if !isAdmin(c, h.dao) {
		uuid, err := getUserUUIDByJWT(c, h.dao)
		if err != nil {
			http.InternalServerError(c, err)
			return
		}

		if uuid != c.Param("uuid") {
			http.Forbidden(c, http.ErrorUserPermission)
			return
		}
	}

	user, err := h.dao.User.FindByUUID(c.Param("uuid"))
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, user)
}

func (h *UserHandler) List(c *gin.Context) {
	if !isAdmin(c, h.dao) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

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

	if !isAdmin(c, h.dao) {
		uuid, err := getUserUUIDByJWT(c, h.dao)
		if err != nil {
			http.InternalServerError(c, err)
			return
		}

		if uuid != user.UUID {
			http.Forbidden(c, http.ErrorUserPermission)
			return
		}
	}

	if err := h.dao.User.Update(user); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	if !isAdmin(c, h.dao) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

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

func (h *UserHandler) UpdateRole(c *gin.Context) {
	if !isAdmin(c, h.dao) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

	role := &models.UserRole{}
	err := c.ShouldBindJSON(&role)
	if err != nil || role.UserUUID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.dao.User.UpdateRole(role); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, role)
}

func (h *UserHandler) UpdateStatus(c *gin.Context) {
	if !isAdmin(c, h.dao) {
		http.Forbidden(c, http.ErrorUserPermission)
		return
	}

	stat := &models.UserStatus{}
	err := c.ShouldBindJSON(&stat)
	if err != nil || stat.UserUUID == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if err := h.dao.User.UpdateStatus(stat); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, stat)

}
