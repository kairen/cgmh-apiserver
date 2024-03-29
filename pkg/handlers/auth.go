package handler

import (
	"time"

	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/ldap"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/services"
	"inwinstack/cgmh/apiserver/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type AuthHandler struct {
	svc  *service.DataAccess
	ldap *ldap.LDAP
}

func (h *AuthHandler) Login(c *gin.Context) {
	login := &model.Login{}
	if err := c.ShouldBindJSON(&login); err != nil || !login.Validate() {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	decode, err := util.Base64Decode(login.Password)
	if err != nil {
		http.BadRequest(c, http.ErrorUserLogin)
		return
	}

	secret := util.MD5Encode(decode)
	if !h.svc.Auth.VerifyAccount(login.Email, secret) {
		http.BadRequest(c, http.ErrorUserLogin)
		return
	}

	user, err := h.svc.User.FindByEmail(login.Email)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}

	if !user.Active {
		http.Forbidden(c, http.ErrorUserNotActive)
		return
	}

	expireTime := time.Hour * time.Duration(viper.GetInt("global.jwtExpireHour"))
	token, err := util.GenerateToken(user.Email, user.UUID, expireTime)
	if err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, gin.H{"token": token})
}

func (h *AuthHandler) Register(c *gin.Context) {
	register := &model.Register{}
	if err := c.ShouldBindJSON(&register); err != nil || !register.Validate() {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	if !register.ValidateEmail() {
		http.BadRequest(c, http.ErrorUserEmailInvalid)
		return
	}

	if h.svc.User.IsExistByEmail(register.Email) {
		http.BadRequest(c, http.ErrorUserRegister)
		return
	}

	decode, err := util.Base64Decode(register.Password)
	if err != nil {
		http.BadRequest(c, http.ErrorUserLogin)
		return
	}

	user := register.ToUser()
	user.Role = model.RoleUser
	secret := util.MD5Encode(decode)
	if err := h.svc.Auth.Register(user, secret); err != nil {
		http.InternalServerError(c, err)
		return
	}

	if err := h.ldap.AddUser(user, decode); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}

func (h *AuthHandler) Reset(c *gin.Context) {
	reset := &model.Reset{}
	if err := c.ShouldBindJSON(&reset); err != nil || !reset.Validate() {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	user, err := h.svc.User.FindByEmail(reset.Email)
	if err != nil {
		http.BadRequest(c, http.ErrorUserNotFound)
		return
	}

	oldBase, err := util.Base64Decode(reset.OldPassword)
	if err != nil {
		http.BadRequest(c, http.ErrorUserLogin)
		return
	}

	oldSecret := util.MD5Encode(oldBase)
	if !h.svc.Auth.VerifyAccount(reset.Email, oldSecret) {
		http.BadRequest(c, http.ErrorUserResetError)
		return
	}

	newBase, err := util.Base64Decode(reset.NewPassword)
	if err != nil {
		http.BadRequest(c, http.ErrorUserLogin)
		return
	}

	secret := util.MD5Encode(newBase)
	if err := h.svc.Auth.Reset(user.UUID, secret); err != nil {
		http.InternalServerError(c, err)
		return
	}

	if err := h.ldap.ModifyPassword(user.UUID, oldBase, newBase); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, nil)
}

func (h *AuthHandler) ForceReset(c *gin.Context) {
	if !checkAdmin(c, h.svc) {
		return
	}

	reset := &model.ForceReset{}
	if err := c.ShouldBindJSON(&reset); err != nil || !reset.Validate() {
		http.BadRequest(c, http.ErrorPayloadField)
		return
	}

	user, err := h.svc.User.FindByEmail(reset.Email)
	if err != nil {
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
	if err := h.svc.Auth.Reset(user.UUID, secret); err != nil {
		http.InternalServerError(c, err)
		return
	}

	if err := h.ldap.ForceModifyPassword(user.UUID, hex); err != nil {
		http.InternalServerError(c, err)
		return
	}
	http.Success(c, gin.H{"secret": encode})
}
