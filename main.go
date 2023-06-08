package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db2 "gitlab.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/db"
	"gitlab.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/modules/admin"
	"gitlab.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/modules/user"
)

func main() {
	router := gin.Default()

	// open connection to db
	db, err := db2.ConnectToDatabase()
	if err != nil {
		fmt.Println("failed to connect to database:", err)
		return
	}

	userRouter := user.NewUserRouter(db)
	userRouter.Handle(router)
	// Customer Handler
	//customerRouter := customer.NewRouter(db)
	//customerRouter.Register(router)

	// Admin Handler
	adminRouter := admin.NewAdminRouter(db)
	adminRouter.Handle(router)

	// Superadmin Handler
	//superadminRouter := superadmin.NewRouter(db)
	//superadminRouter.Register(router)

	err = router.Run(":8081")
	if err != nil {
		fmt.Println("error running server:", err)
		return
	}
}
