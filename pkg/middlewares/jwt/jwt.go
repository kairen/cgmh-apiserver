package jwt

import (
	"time"

	http "inwinstack/cgmh/apiserver/pkg/httpwrapper"
	"inwinstack/cgmh/apiserver/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			http.JSON(c, http.StatusUnauthorized, nil, http.ErrorAuthHeader)
			c.Abort()
			return
		}

		claims, err := util.ParseToken(token)
		if err != nil || time.Now().Unix() > claims.ExpiresAt {
			http.JSON(c, http.StatusUnauthorized, nil, http.ErrorAuthToken)
			c.Abort()
			return
		}
		c.Next()
	}
}
