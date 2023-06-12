package user

import "github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/dto"

type ControllerUserInterface interface {
	CreateUser(req UserParam) (interface{}, error)
	GetUserById(id uint) (interface{}, error)
	UpdateUserById(id uint, user UserParam) (interface{}, error)
	DeleteUserById(id uint) error
}

type ControllerUser struct {
	uc UsecaseUserInterface
}

func NewControllerUser(uc UsecaseUserInterface) ControllerUser {
	return ControllerUser{
		uc: uc,
	}
}

func (ctrl ControllerUser) CreateUser(req UserParam) (interface{}, error) {
	user, err := ctrl.uc.CreateUser(req)
	if err != nil {
		return dto.ErrorResponse{}, err
	}

	response := SuccessCreateUser{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Create User",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: UserParam{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Avatar:    user.Avatar,
		},
	}

	return response, nil
}

func (ctrl ControllerUser) GetUserById(id uint) (interface{}, error) {
	user, err := ctrl.uc.GetUserById(id)
	if err != nil {
		return dto.ErrorResponse{}, err
	}

	response := SuccessCreateUser{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Get User",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: UserParam{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Avatar:    user.Avatar,
		},
	}

	return response, nil
}

func (ctrl ControllerUser) UpdateUserById(id uint, user UserParam) (interface{}, error) {
	updatedUser, err := ctrl.uc.UpdateUserById(id, user)
	if err != nil {
		return dto.ErrorResponse{}, err
	}

	response := SuccessUpdateUser{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Update User",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: UserParam{
			FirstName: updatedUser.FirstName,
			LastName:  updatedUser.LastName,
			Email:     updatedUser.Email,
			Avatar:    updatedUser.Avatar,
		},
	}

	return response, nil
}

func (ctrl ControllerUser) DeleteUserById(id uint) error {
	err := ctrl.uc.DeleteUserById(id)
	if err != nil {
		return err
	}

	return nil
}
