package admin

import "github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/dto"

type AdminParam struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserParam struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type SuccessCreateAdmin struct {
	dto.Response
}

type SuccessCreateUser struct {
	dto.Response
	Data UserParam `json:"data"`
}

type SuccessLoginAdmin struct {
	dto.Response
	Username string `json:"username"`
	Token    string `json:"token"`
}

type SuccessGetAllUsers struct {
	dto.Response
	Data interface{} `json:"data"`
}

type SuccessFetchUsersFromAPI struct {
	dto.Response
}
