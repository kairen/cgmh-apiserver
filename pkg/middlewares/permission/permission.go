package permission

import (
	"github.com/gin-gonic/gin"

	"inwinstack/cgmh/apiserver/pkg/dao"
	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/util"
)

func getUser(c *gin.Context, dao *dao.DataAccess) (*models.User, error) {
	token := c.Request.Header.Get("Authorization")
	claims, _ := util.ParseToken(token)
	uuid, err := util.Base64Decode(claims.UUID)
	if err != nil {
		return nil, err
	}

	user, err := dao.User.FindByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Admin(dao *dao.DataAccess) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := getUser(c, dao)
		if err != nil {
			http.Forbidden(c, http.ErrorUserPermission)
			c.Abort()
		}

		if !user.IsAdmin() {
			http.Forbidden(c, http.ErrorUserPermission)
			c.Abort()
		}
	}
}

func UserQueryAndParam(dao *dao.DataAccess) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := getUser(c, dao)
		if err != nil {
			http.Forbidden(c, http.ErrorUserPermission)
			c.Abort()
		}

		if !user.IsAdmin() {
			uuid := c.Param("uuid")
			query := &models.Query{}
			c.ShouldBindQuery(query)
			if query.UserUUID != "" || uuid != "" {
				if query.UserUUID == user.UUID || uuid == user.UUID {
					return
				}
			}
			http.Forbidden(c, http.ErrorUserPermission)
			c.Abort()
		}
	}
}
