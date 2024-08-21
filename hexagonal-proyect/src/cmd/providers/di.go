package providers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"modularization/hexagonal-proyect/src/cmd/api/handler"
	"modularization/hexagonal-proyect/src/cmd/api/router"
	"modularization/hexagonal-proyect/src/cmd/api/router/groups"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *gin.Engine {
		return gin.Default()
	})

	_ = Container.Provide(router.NewRouter)

	_ = Container.Provide(groups.NewUserGroup)

	_ = Container.Provide(handler.NewUserHandler)

	return Container
}
