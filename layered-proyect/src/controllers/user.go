package controllers

import (
	"github.com/gin-gonic/gin"
	"modularization/layered-proyect/src/interfaces"
	"net/http"
)

type userHandler struct{}

func NewUserHandler() interfaces.UserHandler {
	return &userHandler{}
}

func (uh *userHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User Found Succesfully",
	})
}
