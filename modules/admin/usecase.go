package admin

import (
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/entities"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/helpers"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/mocks"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/repositories"
	"time"
)

type UsecaseAdminInterface interface {
	LoginAdmin(username, password string) (entities.Actor, error)
	RegisterAdmin(admin AdminParam) (entities.Actor, error)
	DeleteUserByID(id uint) error
	CreateUser(user *entities.User) (*entities.User, error)
	GetAllUsers(firstName, lastName, email string, page, pageSize int) ([]*entities.User, error)
	SaveUsersFromAPI() error
}

type UsecaseAdmin struct {
	adminRepo repositories.AdminRepository
	AdminRepo *mocks.MockAdminRepositoryInterface
}

func (uc UsecaseAdmin) LoginAdmin(username, password string) (*entities.Actor, string, error) {
	admin, err := uc.adminRepo.LoginAdmin(username)
	if err != nil {
		return nil, "", err
	}

	// Verify hashed password
	comparePass := helpers.ComparePass([]byte(admin.Password), []byte(password))
	if !comparePass {
		return nil, "", err
	}

	// Generate JWT token
	tokenString, err := helpers.GenerateToken(admin.Id, username)
	if err != nil {
		return nil, "", err
	}

	return admin, tokenString, nil
}

func (uc UsecaseAdmin) RegisterAdmin(admin AdminParam) (*entities.Actor, error) {
	hashedPassword := helpers.HashPass(admin.Password)

	newAdmin := &entities.Actor{
		Username:  admin.Username,
		Password:  hashedPassword,
		RoleId:    2,
		Verified:  false,
		Active:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdAdmin, err := uc.adminRepo.RegisterAdmin(newAdmin)
	if err != nil {
		return nil, err
	}

	return createdAdmin, nil
}

func (uc UsecaseAdmin) CreateUser(user *entities.User) (entities.User, error) {
	newUser := &entities.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Avatar:    user.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := uc.adminRepo.CreateUser(newUser)
	if err != nil {
		return *newUser, err
	}

	return *createdUser, nil
}

func (uc UsecaseAdmin) DeleteUserByID(id uint) error {
	// Get Existing User Data
	existingData, err := uc.adminRepo.GetUserById(id)
	if err != nil {
		return err
	}

	return uc.adminRepo.DeleteUserById(id, existingData)
}

func (uc UsecaseAdmin) GetAllUsers(firstName, lastName, email string, page, pageSize int) ([]*entities.User, error) {
	users, err := uc.adminRepo.GetAllUsers(firstName, lastName, email, page, pageSize)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc UsecaseAdmin) SaveUsersFromAPI() error {
	url := "https://reqres.in/api/users?page=2"

	err := uc.adminRepo.SaveUsersFromAPI(url)
	if err != nil {
		return err
	}

	return nil
}
