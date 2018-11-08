package auth

import (
	"github.com/gin-gonic/gin"

	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/util"
)

func getUser(c *gin.Context) (*models.User, error) {
	token := c.Request.Header.Get("Authorization")
	claims, _ := util.ParseToken(token)
	email, err := util.Base64Decode(claims.Email)
	if err != nil {
		return nil, err
	}

	dao := &models.User{}
	user, err := dao.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := getUser(c)
		if err != nil {
			http.JSON(c, http.StatusForbidden, nil, http.ErrorUserPermission)
			c.Abort()
		}

		if !user.IsAdmin {
			http.JSON(c, http.StatusForbidden, nil, http.ErrorUserPermission)
			c.Abort()
		}
	}
}

func UserUUIDQueryRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := getUser(c)
		if err != nil {
			http.JSON(c, http.StatusForbidden, nil, http.ErrorUserPermission)
			c.Abort()
		}

		if !user.IsAdmin {
			query := &models.Query{}
			c.ShouldBindQuery(query)
			if query.UserUUID != user.UUID {
				http.JSON(c, http.StatusForbidden, nil, http.ErrorUserPermission)
				c.Abort()
			}
		}
	}
}
