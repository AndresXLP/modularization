package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler interface {
	GetUser(c *gin.Context)
}

type userHandler struct{}

func NewUserHandler() UserHandler {
	return &userHandler{}
}

func (uh *userHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User Found Succesfully",
	})
}
