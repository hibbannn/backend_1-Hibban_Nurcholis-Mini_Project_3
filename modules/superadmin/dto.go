package superadmin

import "github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/dto"

type UserParam struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type LoginSuperadmin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SuperAdminParam struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SuccessCreateSuperadmin struct {
	dto.Response
}

type SuccessCreateUser struct {
	dto.Response
	Data UserParam `json:"data"`
}

type SuccessUpdateRegisterAdmin struct {
	dto.Response
}

type SuccessGetApprovalRequest struct {
	dto.Response
	Data interface{} `json:"data"`
}

type SuccessLoginSuperadmin struct {
	dto.Response
	Username string `json:"username"`
	Token    string `json:"token"`
}

type SuccessGetAllUsers struct {
	dto.Response
	Data interface{} `json:"data"`
}

type SuccessGetAllAdmins struct {
	dto.Response
	Data interface{} `json:"data"`
}
