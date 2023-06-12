package user

import (
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/dto"
)

// UserParam represents the user parameter in the request
type UserParam struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

// SuccessCreateUser represents the success response for user creation
type SuccessCreateUser struct {
	dto.Response
	Data UserParam `json:"data"`
}

// SuccessUpdateUser represents the success response for user update
type SuccessUpdateUser struct {
	dto.Response
	Data UserParam `json:"data"`
}
