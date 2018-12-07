package handler

import (
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/ldap"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/services"
	"inwinstack/cgmh/apiserver/pkg/util"
	"inwinstack/cgmh/apiserver/pkg/version"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type GlobalHandler struct {
	svc *service.DataAccess

	// Sub handlers
	Auth         *AuthHandler
	User         *UserHandler
	Level        *LevelHandler
	Form         *FormHandler
	PointHistory *PointHandler
}

func New(svc *service.DataAccess) *GlobalHandler {
	h := &GlobalHandler{svc: svc}
	h.Auth = &AuthHandler{svc: svc}
	h.User = &UserHandler{svc: svc}
	h.Level = &LevelHandler{svc: svc}
	h.Form = &FormHandler{svc: svc}
	h.PointHistory = &PointHandler{svc: svc}
	return h
}

func (h *GlobalHandler) GetService() *service.DataAccess {
	return h.svc
}

func (h *GlobalHandler) SetLDAP(ldap *ldap.LDAP) {
	h.Auth.ldap = ldap
	h.User.ldap = ldap
}

func (h *GlobalHandler) Version(c *gin.Context) {
	c.JSON(http.StatusSuccess, map[string]string{
		"version": version.GetVersion(),
	})
}

func (h *GlobalHandler) Healthz(c *gin.Context) {
	c.String(http.StatusSuccess, "ok")
}

func (h *GlobalHandler) MonitorURL(c *gin.Context) {
	c.JSON(http.StatusSuccess, map[string]string{
		"url": viper.GetString("global.monitorURL"),
	})
}

func getUserUUIDByJWT(c *gin.Context) (string, error) {
	token := c.Request.Header.Get("Authorization")
	claims, err := util.ParseToken(token)
	if err != nil {
		return "", err
	}

	uuid, err := util.Base64Decode(claims.UUID)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func getUserByJWT(c *gin.Context, svc *service.DataAccess) (*model.User, error) {
	uuid, err := getUserUUIDByJWT(c)
	if err != nil {
		return nil, err
	}

	user, err := svc.User.FindByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func isAdmin(c *gin.Context, svc *service.DataAccess) bool {
	jwtUser, err := getUserByJWT(c, svc)
	if err != nil {
		return false
	}

	if !jwtUser.IsAdmin() {
		return false
	}
	return true
}

func checkAdmin(c *gin.Context, svc *service.DataAccess) bool {
	if !isAdmin(c, svc) {
		http.Forbidden(c, http.ErrorUserPermission)
		return false
	}
	return true
}

func checkUserUUID(c *gin.Context, svc *service.DataAccess, userUUID string) bool {
	if !isAdmin(c, svc) {
		uuid, err := getUserUUIDByJWT(c)
		if err != nil {
			http.InternalServerError(c, err)
			return false
		}

		if uuid != userUUID {
			http.Forbidden(c, http.ErrorUserPermission)
			return false
		}
	}
	return true
}
