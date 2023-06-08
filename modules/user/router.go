package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterUser struct {
	UserRequestHandler RequestHandlerUser
}

func NewUserRouter(db *gorm.DB) RouterUser {
	return RouterUser{UserRequestHandler: NewUserRequestHandler(db)}
}

func (r RouterUser) Handle(router *gin.Engine) {
	basePath := "/user"

	user := router.Group(basePath)
	user.POST("/create", r.UserRequestHandler.CreateUser)
	user.GET("/:id", r.UserRequestHandler.GetUserById)
	user.PUT("/:id", r.UserRequestHandler.UpdateUserById)
	user.DELETE("/:id", r.UserRequestHandler.DeleteUserById)
}
