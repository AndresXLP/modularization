package providers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"modularization/layered-proyect/src/controllers"
	"modularization/layered-proyect/src/router"
	"modularization/layered-proyect/src/router/groups"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *gin.Engine {
		return gin.Default()
	})

	_ = Container.Provide(router.NewRouter)

	_ = Container.Provide(groups.NewUserGroup)

	_ = Container.Provide(controllers.NewUserHandler)

	return Container
}
