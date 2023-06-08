package admin

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/middleware"
	"gitlab.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/repositories"
	"gorm.io/gorm"
)

type AdminRouter struct {
	adminHandler *AdminRequestHandler
}

func NewAdminRouter(db *gorm.DB) *AdminRouter {
	adminRepo := repositories.NewAdminRepository(db)
	adminUsecase := UsecaseAdmin{
		adminRepo: adminRepo,
	}
	adminController := ControllerAdmin{
		uc: adminUsecase,
	}
	adminHandler := NewAdminRequestHandler(&adminController)

	return &AdminRouter{
		adminHandler: adminHandler,
	}
}

func (r *AdminRouter) Handle(router *gin.Engine) {
	basePath := "/admin"

	admin := router.Group(basePath)
	{
		admin.POST("/login", r.adminHandler.LoginAdmin)
		admin.POST("/register", r.adminHandler.RegisterAdmin)

		admin.Use(middleware.Authentication())
		{
			admin.POST("/create-customer", r.adminHandler.CreateCustomer)
			admin.DELETE("/delete-customer/:id", r.adminHandler.DeleteCustomerById)
			admin.GET("/customers", r.adminHandler.GetAllCustomers)
			admin.GET("/fetch-customers", r.adminHandler.SaveCustomersFromAPI)
		}
	}
}
