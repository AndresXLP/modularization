package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"modularization/hexagonal-proyect/src/cmd/api/router"
	"modularization/hexagonal-proyect/src/cmd/providers"
	"modularization/hexagonal-proyect/src/config"
)

func main() {
	port := config.Environments().AppPort

	container := providers.BuildContainer()
	if err := container.Invoke(func(server *gin.Engine, router *router.Router) {
		router.Init()
		server.Run(fmt.Sprintf(":%d", port))
	}); err != nil {
		log.Fatal(err)
	}
}
