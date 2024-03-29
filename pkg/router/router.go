package router

import (
	_ "inwinstack/cgmh/apiserver/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"inwinstack/cgmh/apiserver/pkg/handlers"
	"inwinstack/cgmh/apiserver/pkg/middlewares/jwt"
	"inwinstack/cgmh/apiserver/pkg/services"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Router struct {
	engine  *gin.Engine
	handler *handler.GlobalHandler
}

func New(svc *service.DataAccess) *Router {
	gin.DisableConsoleColor()
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	return &Router{
		engine:  engine,
		handler: handler.New(svc),
	}
}

func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}

func (r *Router) GetHandler() *handler.GlobalHandler {
	return r.handler
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
	r.engine.GET("/version", r.handler.Version)
	r.engine.GET("/healthz", r.handler.Healthz)
	r.engine.GET("/monitorurl", r.handler.MonitorURL)
	r.engine.POST("/auth/login", r.handler.Auth.Login)
	r.engine.POST("/auth/register", r.handler.Auth.Register)
	r.engine.PUT("/auth/reset", r.handler.Auth.Reset)
	r.engine.Group("").Use(jwt.JWT()).PUT("/auth/forcereset", r.handler.Auth.ForceReset)

	// V1 API
	apiv1 := r.engine.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/user", r.handler.User.List)
		apiv1.GET("/user/:uuid", r.handler.User.Get)
		apiv1.PUT("/user", r.handler.User.Update)
		apiv1.DELETE("/user", r.handler.User.Delete)
		apiv1.PUT("/userrole", r.handler.User.UpdateRole)
		apiv1.PUT("/userstatus", r.handler.User.UpdateStatus)
		apiv1.PUT("/userlevel", r.handler.User.UpdateLevel)
		apiv1.PUT("/userpoint", r.handler.User.UpdatePoint)

		apiv1.GET("/form", r.handler.Form.List)
		apiv1.GET("/form/:id", r.handler.Form.Get)
		apiv1.POST("/form", r.handler.Form.Create)
		apiv1.PUT("/form", r.handler.Form.Update)
		apiv1.PUT("/formstatus", r.handler.Form.UpdateStatus)
		apiv1.DELETE("/form", r.handler.Form.Delete)

		apiv1.GET("/level", r.handler.Level.List)
		apiv1.GET("/level/:id", r.handler.Level.Get)
		apiv1.POST("/level", r.handler.Level.Create)
		apiv1.PUT("/level", r.handler.Level.Update)
		apiv1.DELETE("/level", r.handler.Level.Delete)
		apiv1.PUT("/leveldefault", r.handler.Level.UpdateDefault)

		apiv1.GET("/pointhistory", r.handler.PointHistory.List)
	}
}
