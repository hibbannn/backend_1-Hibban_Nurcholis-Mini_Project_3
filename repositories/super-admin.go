package repositories

import (
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/entities"
	"gorm.io/gorm"
)

type SuperAdminRepository interface {
	// Admin
	CreateSuperAdmin(superAdmin *entities.Actor) (*entities.Actor, error)
	ApproveAdminRegistration(id uint) error
	GetAllAdmins(username string, page, pageSize int) ([]*entities.Actor, error)
	LoginSuperAdmin(username string) (*entities.Actor, error)
	RejectAdminRegistration(id uint) error
	UpdateAdminActiveStatus(id uint, isActive bool) error
	UpdateDeadactivedAdmin(id uint) error

	// User
	CreateUser(user *entities.User) (*entities.User, error)
	DeleteUserById(id uint, user *entities.User) error
	GetAllUsers(firstName, lastName, email string, page, pageSize int) ([]*entities.User, error)
	GetUserById(id uint) (*entities.User, error)

	// Approval
	GetApprovalRequests() ([]*entities.Actor, error)
}

type SuperAdmin struct {
	db *gorm.DB
}

func NewSuperAdmin(db *gorm.DB) SuperAdmin {
	return SuperAdmin{
		db: db,
	}
}

func (repo SuperAdmin) CreateSuperAdmin(superAdmin *entities.Actor) (*entities.Actor, error) {
	err := repo.db.Create(superAdmin).Error
	if err != nil {
		return nil, err
	}

	return superAdmin, nil
}

func (repo SuperAdmin) LoginSuperAdmin(username string) (*entities.Actor, error) {
	superAdmin := &entities.Actor{}

	err := repo.db.Where("username = ?", username).First(superAdmin).Error
	if err != nil {
		return nil, err
	}

	return superAdmin, nil
}

func (repo SuperAdmin) CreateUser(user *entities.User) (*entities.User, error) {
	err := repo.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo SuperAdmin) GetUserById(id uint) (*entities.User, error) {
	user := &entities.User{}

	err := repo.db.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo SuperAdmin) DeleteUserById(id uint) error {
	return repo.db.Delete(&entities.User{}, id).Error
}

func (repo SuperAdmin) GetAllUsers(firstName, lastName, email string, page, pageSize int) ([]*entities.User, error) {
	var users []*entities.User

	query := repo.db.Model(&entities.User{})

	if firstName != "" {
		query = query.Where("first_name LIKE ?", "%"+firstName+"%")
	} else if lastName != "" {
		query = query.Where("last_name LIKE ?", "%"+lastName+"%")
	} else if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	// Pagination
	offset := (page - 1) * pageSize

	err := query.Offset(offset).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo SuperAdmin) ApproveAdminRegistration(id uint) error {
	admin := &entities.Actor{}

	err := repo.db.Where("id = ?", id).First(admin).Error
	if err != nil {
		return err
	}

	admin.Verified = true
	return repo.db.Save(admin).Error
}

func (repo SuperAdmin) RejectAdminRegistration(id uint) error {
	admin := &entities.Actor{}

	err := repo.db.Where("id = ?", id).First(admin).Error
	if err != nil {
		return err
	}

	return repo.db.Delete(admin).Error
}

func (repo SuperAdmin) UpdateAdminActiveStatus(id uint, isActive bool) error {
	admin := &entities.Actor{}

	err := repo.db.Where("id = ?", id).First(admin).Error
	if err != nil {
		return err
	}

	admin.Active = isActive
	return repo.db.Save(admin).Error
}

func (repo SuperAdmin) GetApprovalRequests() ([]*entities.Actor, error) {
	var admins []*entities.Actor

	err := repo.db.Model(&entities.Actor{}).
		Select("id, username, role_id, is_verified, created_at, updated_at").
		Where("role_id = ? AND is_verified = ?", 2, false).
		Find(&admins).Error
	if err != nil {
		return nil, err
	}

	return admins, nil
}

func (repo SuperAdmin) GetAllAdmins(username string, page, pageSize int) ([]*entities.Actor, error) {
	var admins []*entities.Actor

	query := repo.db.Model(&entities.Actor{})

	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}

	query.Where("role_id = ? AND is_verified = ?", 2, true)

	// Pagination
	offset := (page - 1) * pageSize

	err := query.Offset(offset).Limit(pageSize).Find(&admins).Error
	if err != nil {
		return nil, err
	}

	return admins, nil

}

func (repo SuperAdmin) UpdateDeadactivedAdmin(id uint) error {
	// Check admin data by id
	err := repo.db.Where("id = ?", id).First(&entities.Actor{}).Error
	if err != nil {
		return err
	}

	admin := &entities.Actor{}
	err = repo.db.First(admin, id).Error
	if err != nil {
		return err
	}

	// Update Active
	if admin.Active == true {
		err = repo.db.Model(&entities.Actor{}).Where("id = ?", id).Update("active", false).Error
		if err != nil {
			return err
		}
	}

	return nil
}
