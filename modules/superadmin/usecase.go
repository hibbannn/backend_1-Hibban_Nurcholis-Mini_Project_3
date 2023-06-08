package superadmin

import (
	"gitlab.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/entities"
	"gitlab.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/helpers"
	"gitlab.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/repositories"
	"time"
)

type UsecaseSuperadminInterface interface {
	CreateSuperadmin(superadmin SuperAdminParam) (*entities.Actor, error)
	LoginSuperadmin(id uint, username, password string) (*entities.Actor, string, error)
	CreateUser(user UserParam) (entities.User, error)
	DeleteUserByID(id uint) error
	GetAllUsers(firstName, lastName, email string, page, pageSize int) ([]*entities.User, error)
	ApprovedAdminRegister(id uint) error
	RejectedAdminRegister(id uint) error
	UpdateActivedAdmin(id uint) error
	UpdateDeadactivedAdmin(id uint) error
	GetApprovalRequest() ([]*entities.User, error)
	GetAllAdmins(username string, page, pageSize int) ([]*entities.User, error)
}

type UsecaseSuperadmin struct {
	superadminRepo repositories.SuperAdminRepository
}

func (uc UsecaseSuperadmin) CreateSuperadmin(superadmin SuperAdminParam) (*entities.Actor, error) {
	hasPass := helpers.HashPass(superadmin.Password)

	newSuperadmin := &entities.Actor{
		Username:  superadmin.Username,
		Password:  hasPass,
		RoleId:    1,
		Verified:  true,
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createSuperadmin, err := uc.superadminRepo.CreateSuperAdmin(newSuperadmin)
	if err != nil {
		return nil, err
	}

	return createSuperadmin, nil
}

func (uc UsecaseSuperadmin) LoginSuperadmin(id uint, username, password string) (*entities.Actor, string, error) {
	superadmin, err := uc.superadminRepo.LoginSuperAdmin(username)
	if err != nil {
		return nil, "", err
	}

	// Verify hashed password
	comparePass := helpers.ComparePass([]byte(superadmin.Password), []byte(password))
	if !comparePass {
		return nil, "", err
	}

	// Generate JWT token
	tokenString, err := helpers.GenerateToken(id, username)
	if err != nil {
		return nil, "", err
	}

	return superadmin, tokenString, nil
}

func (uc UsecaseSuperadmin) CreateUser(user UserParam) (entities.User, error) {
	newUser := &entities.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Avatar:    user.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := uc.superadminRepo.CreateUser(newUser)
	if err != nil {
		return *newUser, err
	}

	return *newUser, nil
}

func (uc UsecaseSuperadmin) DeleteUserByID(id uint) error {
	// Get existing user data
	existingData, err := uc.superadminRepo.GetUserById(id)
	if err != nil {
		return err
	}

	return uc.superadminRepo.DeleteUserById(id, existingData)
}

func (uc UsecaseSuperadmin) GetAllUsers(firstName, lastName, email string, page, pageSize int) ([]*entities.User, error) {
	users, err := uc.superadminRepo.GetAllUsers(firstName, lastName, email, page, pageSize)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc UsecaseSuperadmin) ApprovedAdminRegister(id uint) error {
	err := uc.superadminRepo.ApproveAdminRegistration(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc UsecaseSuperadmin) RejectedAdminRegister(id uint) error {
	err := uc.superadminRepo.RejectAdminRegistration(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc UsecaseSuperadmin) UpdateActivedAdmin(id uint, isActive bool) error {
	err := uc.superadminRepo.UpdateAdminActiveStatus(id, isActive)
	if err != nil {
		return err
	}

	return nil
}

func (uc UsecaseSuperadmin) UpdateDeadactivedAdmin(id uint) error {
	err := uc.superadminRepo.UpdateDeadactivedAdmin(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc UsecaseSuperadmin) GetApprovalRequest() ([]*entities.Actor, error) {
	actors, err := uc.superadminRepo.GetApprovalRequests()
	if err != nil {
		return nil, err
	}

	updatedActors := make([]*entities.Actor, len(actors))
	for i, actor := range actors {
		updatedActors[i] = &entities.Actor{
			Id:        actor.Id,
			Username:  actor.Username,
			RoleId:    actor.RoleId,
			Verified:  actor.Verified,
			CreatedAt: actor.CreatedAt,
			UpdatedAt: actor.UpdatedAt,
		}
	}

	return updatedActors, nil
}

func (uc UsecaseSuperadmin) GetAllAdmins(username string, page, pageSize int) ([]*entities.Actor, error) {
	admins, err := uc.superadminRepo.GetAllAdmins(username, page, pageSize)
	if err != nil {
		return nil, err
	}

	return admins, nil
}
