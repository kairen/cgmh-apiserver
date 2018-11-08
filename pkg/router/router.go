package router

import (
	_ "inwinstack/cgmh/apiserver/api"

	"github.com/gin-gonic/gin"

	"inwinstack/cgmh/apiserver/pkg/handlers"
	"inwinstack/cgmh/apiserver/pkg/handlers/v1"
	"inwinstack/cgmh/apiserver/pkg/middlewares/auth"
	"inwinstack/cgmh/apiserver/pkg/middlewares/jwt"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	gin.DisableConsoleColor()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/version", handler.GetVersion)
	r.GET("/healthz", handler.GetHealthz)
	r.POST("/auth/login", handler.Login)
	r.POST("/auth/register", handler.Register)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/form/:id", v1.GetForm)
		apiv1.POST("/form", v1.CreateForm)
	}

	user := apiv1.Group("")
	user.Use(auth.UserUUIDQueryRequired())
	{
		user.GET("/form", v1.ListForm)
	}

	admin := apiv1.Group("")
	admin.Use(auth.AdminRequired())
	{
		admin.GET("/user", v1.ListUser)
		admin.PUT("/user", v1.UpdateUser)
		admin.DELETE("/user", v1.DeleteUser)
		admin.DELETE("/form", v1.DeleteForm)
		admin.PUT("/form", v1.UpdateForm)
	}
	return r
}
