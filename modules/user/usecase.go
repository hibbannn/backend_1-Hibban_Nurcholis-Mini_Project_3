package user

import (
	"gitlab.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/entities"
	"gitlab.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/repositories"
	"time"
)

type UsecaseUserInterface interface {
	CreateUser(user UserParam) (entities.User, error)
	GetUserById(id uint) (entities.User, error)
	UpdateUserById(id uint, user UserParam) (entities.User, error)
	DeleteUserById(id uint) error
}

type UsecaseUser struct {
	userRepo repositories.UserRepositoryInterface
}

func NewUsecaseUser(userRepo repositories.UserRepositoryInterface) UsecaseUser {
	return UsecaseUser{
		userRepo: userRepo,
	}
}

func (uc UsecaseUser) CreateUser(user UserParam) (entities.User, error) {
	newUser := entities.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Avatar:    user.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := uc.userRepo.CreateUser(&newUser)
	if err != nil {
		return newUser, err
	}

	return *createdUser, nil
}

func (uc UsecaseUser) GetUserById(id uint) (entities.User, error) {
	user, err := uc.userRepo.GetUserById(id)
	if err != nil {
		return entities.User{}, err
	}

	return *user, nil
}

func (uc UsecaseUser) UpdateUserById(id uint, user UserParam) (entities.User, error) {
	existingUser, err := uc.userRepo.GetUserById(id)
	if err != nil {
		return entities.User{}, err
	}

	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Email = user.Email
	existingUser.Avatar = user.Avatar
	existingUser.UpdatedAt = time.Now()

	updatedUser, err := uc.userRepo.UpdateUserById(id, existingUser)
	if err != nil {
		return entities.User{}, err
	}

	return *updatedUser, nil
}

func (uc UsecaseUser) DeleteUserById(id uint) error {
	err := uc.userRepo.DeleteUserById(id)
	if err != nil {
		return err
	}

	return nil
}
