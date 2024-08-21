package interfaces

import "github.com/gin-gonic/gin"

// UserHandler Interfaces
type UserHandler interface {
	GetUser(c *gin.Context)
}
