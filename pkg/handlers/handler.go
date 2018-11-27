package handler

import (
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/services"
	"inwinstack/cgmh/apiserver/pkg/util"
	"inwinstack/cgmh/apiserver/pkg/version"

	"github.com/gin-gonic/gin"
)

type GlobalHandler struct {
	svc *service.DataAccess

	// Sub handlers
	Auth  *AuthHandler
	User  *UserHandler
	Level *LevelHandler
	Form  *FormHandler
}

func New(svc *service.DataAccess) *GlobalHandler {
	h := &GlobalHandler{svc: svc}
	h.Auth = &AuthHandler{svc: svc}
	h.User = &UserHandler{svc: svc}
	h.Level = &LevelHandler{svc: svc}
	h.Form = &FormHandler{svc: svc}
	return h
}

func (h *GlobalHandler) GetVersion(c *gin.Context) {
	c.JSON(http.StatusSuccess, map[string]string{
		"version": version.GetVersion(),
	})
}

func (h *GlobalHandler) GetHealthz(c *gin.Context) {
	c.String(http.StatusSuccess, "ok")
}

func (h *GlobalHandler) GetService() *service.DataAccess {
	return h.svc
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

func getUserByJWT(c *gin.Context, svc *service.DataAccess) (*models.User, error) {
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
