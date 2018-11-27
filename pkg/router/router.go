package router

import (
	_ "inwinstack/cgmh/apiserver/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"inwinstack/cgmh/apiserver/pkg/dao"
	"inwinstack/cgmh/apiserver/pkg/handlers"
	"inwinstack/cgmh/apiserver/pkg/middlewares/jwt"
	"inwinstack/cgmh/apiserver/pkg/middlewares/permission"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Router struct {
	engine  *gin.Engine
	handler *handler.GlobalHandler
}

func New(dao *dao.DataAccess) *Router {
	gin.DisableConsoleColor()
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	return &Router{
		engine:  engine,
		handler: handler.New(dao),
	}
}

func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}

func (r *Router) SetCORS(config cors.Config) {
	r.engine.Use(cors.New(config))
}

func (r *Router) LinkSwaggerAPI(swagger bool) {
	if swagger {
		r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func (r *Router) LinkHandlers() {
	dao := r.handler.GetDAO()

	r.engine.GET("/version", r.handler.GetVersion)
	r.engine.GET("/healthz", r.handler.GetHealthz)
	r.engine.POST("/auth/login", r.handler.Auth.Login)
	r.engine.POST("/auth/register", r.handler.Auth.Register)
	r.engine.PUT("/auth/reset", r.handler.Auth.Reset)

	apiv1 := r.engine.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/user/:uuid", r.handler.User.Get)
		apiv1.GET("/form", r.handler.Form.List)
		apiv1.GET("/form/:id", r.handler.Form.Get)
		apiv1.POST("/form", r.handler.Form.Create)
	}

	// Require admin user for common API
	admin := r.engine.Group("")
	admin.Use(jwt.JWT())
	admin.Use(permission.Admin(dao))
	{
		admin.PUT("/auth/forcereset", r.handler.Auth.ForceReset)
	}

	// Require admin user for V1 API
	adminv1 := apiv1.Group("")
	adminv1.Use(permission.Admin(dao))
	{
		adminv1.GET("/user", r.handler.User.List)
		adminv1.PUT("/user", r.handler.User.Update)
		adminv1.DELETE("/user", r.handler.User.Delete)
		adminv1.DELETE("/form", r.handler.Form.Delete)
		adminv1.PUT("/form", r.handler.Form.Update)
	}
}
