package groups

import (
	"github.com/gin-gonic/gin"
	"modularization/hexagonal-proyect/src/cmd/api/handler"
)

type UserGroup interface {
	Resource(e *gin.Engine)
}

type userRoute struct {
	user handler.UserHandler
}

func NewUserGroup(user handler.UserHandler) UserGroup {
	return &userRoute{user}
}

func (group userRoute) Resource(e *gin.Engine) {
	userGroup := e.Group("/user")
	userGroup.GET("", group.user.GetUser)
}
