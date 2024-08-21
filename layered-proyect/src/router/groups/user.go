package groups

import (
	"github.com/gin-gonic/gin"
	"modularization/layered-proyect/src/interfaces"
)

type UserGroup interface {
	Resource(e *gin.Engine)
}

type userRoute struct {
	user interfaces.UserHandler
}

func NewUserGroup(user interfaces.UserHandler) UserGroup {
	return &userRoute{user}
}

func (group userRoute) Resource(e *gin.Engine) {
	userGroup := e.Group("/user")
	userGroup.GET("", group.user.GetUser)
}
