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
	if !h.dao.Auth.VerifyAccount(login.Email, secret) {
		http.BadRequest(c, http.ErrorUserLogin)
		return
	}

	user, err := h.dao.User.FindByEmail(login.Email)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}

	if !user.Active {
		http.Forbidden(c, http.ErrorUserNotActive)
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

	if !register.ValidateEmail() {
		http.BadRequest(c, http.ErrorUserEmailInvalid)
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
	if err := h.dao.Auth.Register(user, secret); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}

func (h *AuthHandler) Reset(c *gin.Context) {
	reset := &models.Reset{}
	err := c.ShouldBindJSON(&reset)
	if err != nil || reset.Email == "" || reset.OldPassword == "" || reset.NewPassword == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	oldBase, err := util.Base64Decode(reset.OldPassword)
	if err != nil {
		http.BadRequest(c, http.ErrorUserLogin)
		return
	}

	oldSecret := util.MD5Encode(oldBase)
	if !h.dao.Auth.VerifyAccount(reset.Email, oldSecret) {
		http.BadRequest(c, http.ErrorUserResetError)
		return
	}

	newBase, err := util.Base64Decode(reset.NewPassword)
	if err != nil {
		http.BadRequest(c, http.ErrorUserLogin)
		return
	}

	secret := util.MD5Encode(newBase)
	if err := h.dao.Auth.Reset(reset.Email, secret); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}

func (h *AuthHandler) ForceReset(c *gin.Context) {
	reset := &models.ForceReset{}
	err := c.ShouldBindJSON(&reset)
	if err != nil || reset.Email == "" {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if !h.dao.User.IsExistByEmail(reset.Email) {
		http.BadRequest(c, http.ErrorUserNotFound)
		return
	}

	hex, err := util.RandomHex(8)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}

	encode := util.Base64Encode(hex)
	secret := util.MD5Encode(hex)
	if err := h.dao.Auth.Reset(reset.Email, secret); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, gin.H{"secret": encode})
}
