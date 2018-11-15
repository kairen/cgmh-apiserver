package handler

import (
	"time"

	"inwinstack/cgmh/apiserver/pkg/dao"
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/util"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	dao *dao.DataAccess
}

func (h *AuthHandler) Login(c *gin.Context) {
	login := &models.Login{}
	err := c.ShouldBindJSON(&login)
	if err != nil || login.Email == "" || login.Password == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	decode, err := util.Base64Decode(login.Password)
	if err != nil {
		http.BadRequest(c, http.ErrorUserLogin)
		return
	}

	secret := util.MD5Encode(decode)
	if !h.dao.User.VerifyAccount(login.Email, secret) {
		http.BadRequest(c, http.ErrorUserLogin)
		return
	}

	user, err := h.dao.User.FindByEmail(login.Email)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}

	if !user.Active {
		http.BadRequest(c, http.ErrorUserNotActive)
		return
	}

	token, err := util.GenerateToken(user.Email, user.UUID, 1*time.Hour)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, gin.H{"token": token})
}

func (h *AuthHandler) Register(c *gin.Context) {
	register := &models.Register{}
	err := c.ShouldBindJSON(&register)
	if err != nil || register.Email == "" || register.Password == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if h.dao.User.IsExistByEmail(register.Email) {
		http.BadRequest(c, http.ErrorUserRegister)
		return
	}

	decode, err := util.Base64Decode(register.Password)
	if err != nil {
		http.BadRequest(c, http.ErrorUserLogin)
		return
	}

	user := register.ToUser()
	user.Role = models.RoleUser
	secret := util.MD5Encode(decode)
	if err := h.dao.User.Register(user, secret); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}
