package handler

import (
	"inwinstack/cgmh/apiserver/pkg/dao"
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/util"
	"inwinstack/cgmh/apiserver/pkg/version"

	"github.com/gin-gonic/gin"
)

type GlobalHandler struct {
	dao *dao.DataAccess

	// subhandler
	Auth *AuthHandler
	User *UserHandler
	Form *FormHandler
}

func New(dao *dao.DataAccess) *GlobalHandler {
	h := &GlobalHandler{dao: dao}
	h.Auth = &AuthHandler{dao: dao}
	h.User = &UserHandler{dao: dao}
	h.Form = &FormHandler{dao: dao}
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

func (h *GlobalHandler) GetDAO() *dao.DataAccess {
	return h.dao
}

func getUserUUIDByJWT(c *gin.Context, dao *dao.DataAccess) (string, error) {
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

func getUserByJWT(c *gin.Context, dao *dao.DataAccess) (*models.User, error) {
	uuid, err := getUserUUIDByJWT(c, dao)
	if err != nil {
		return nil, err
	}

	user, err := dao.User.FindByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func isAdmin(c *gin.Context, dao *dao.DataAccess) bool {
	jwtUser, err := getUserByJWT(c, dao)
	if err != nil {
		return false
	}

	if !jwtUser.IsAdmin() {
		return false
	}
	return true
}
