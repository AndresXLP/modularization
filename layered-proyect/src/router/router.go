package router

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"modularization/layered-proyect/src/controllers"
	"modularization/layered-proyect/src/router/groups"
)

type Router struct {
	server    *gin.Engine
	userGroup groups.UserGroup
}

func NewRouter(server *gin.Engine, group groups.UserGroup) *Router {
	return &Router{server, group}
}

func (r *Router) Init() {
	r.server.Use(gin.Recovery())

	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(r.server)

	r.server.GET("/ping", controllers.Ping)

	r.userGroup.Resource(r.server)
}
