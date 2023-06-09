package user

import (
	"github.com/gin-gonic/gin"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/dto"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/repositories"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandlerUser struct {
	ctrl ControllerUser
}

func NewUserRequestHandler(db *gorm.DB) RequestHandlerUser {
	userRepo := repositories.NewUserRepository(db)
	uc := UsecaseUser{userRepo}
	ctrl := ControllerUser{uc}

	return RequestHandlerUser{ctrl}
}

func (rh RequestHandlerUser) CreateUser(c *gin.Context) {
	request := UserParam{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := rh.ctrl.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (rh RequestHandlerUser) GetUserById(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	userID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	response, err := rh.ctrl.GetUserById(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (rh RequestHandlerUser) UpdateUserById(c *gin.Context) {
	id := c.Param("id")
	request := UserParam{}

	// Bind JSON
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	// Parse id to uint
	userID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	response, err := rh.ctrl.UpdateUserById(uint(userID), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (rh RequestHandlerUser) DeleteUserById(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	userID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	err = rh.ctrl.DeleteUserById(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Delete User Data Successfully",
	})
}
