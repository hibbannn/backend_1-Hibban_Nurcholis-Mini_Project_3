package superadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/middleware"
)

type RouterSuperadmin struct {
	SuperadminRequestHandler *RequestHandlerSuperadmin
}

func NewRouterSuperadmin(ctrl ControllerSuperadminInterface) *RouterSuperadmin {
	return &RouterSuperadmin{
		SuperadminRequestHandler: NewRequestHandlerSuperadmin(ctrl),
	}
}

func (r *RouterSuperadmin) Handle(router *gin.Engine) {
	basePath := "/superadmin"

	superadmin := router.Group(basePath)
	superadmin.POST("/register", r.SuperadminRequestHandler.CreateSuperadmin)
	superadmin.POST("/login", r.SuperadminRequestHandler.LoginSuperadmin)

	// About User
	superadmin.Use(middleware.Authentication())
	superadmin.GET("/users", r.SuperadminRequestHandler.GetAllUsers)
	superadmin.POST("/create-user", r.SuperadminRequestHandler.CreateUser)
	superadmin.DELETE("/delete-user/:id", r.SuperadminRequestHandler.DeleteUserByID)

	// About Admin
	superadmin.Use(middleware.Authentication())
	superadmin.GET("/admins", r.SuperadminRequestHandler.GetAllAdmins)
	superadmin.POST("/:id/approved", r.SuperadminRequestHandler.ApproveAdminRegistration)
	superadmin.POST("/:id/rejected", r.SuperadminRequestHandler.RejectAdminRegistration)
	superadmin.POST("/:id/actived", r.SuperadminRequestHandler.UpdateAdminActiveStatus)
	superadmin.POST("/:id/deadactived", r.SuperadminRequestHandler.UpdateAdminInactiveStatus)
	superadmin.GET("/approval-request", r.SuperadminRequestHandler.GetApprovalRequests)
}
