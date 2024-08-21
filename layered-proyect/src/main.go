package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"modularization/layered-proyect/src/config"
	"modularization/layered-proyect/src/providers"
	"modularization/layered-proyect/src/router"
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
