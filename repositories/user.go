package repositories

import (
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/entities"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	// Create
	CreateUser(user *entities.User) (*entities.User, error)

	// Retrieve
	GetUserById(id uint) (*entities.User, error)

	// Update
	UpdateUserById(id uint, user *entities.User) (*entities.User, error)

	// Delete
	DeleteUserById(id uint) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (repo UserRepository) CreateUser(user *entities.User) (*entities.User, error) {
	err := repo.db.Model(&entities.User{}).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepository) GetUserById(id uint) (*entities.User, error) {
	user := &entities.User{}

	err := repo.db.Model(&entities.User{}).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepository) UpdateUserById(id uint, user *entities.User) (*entities.User, error) {
	err := repo.db.Model(&entities.User{}).Where("id = ?", id).Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepository) DeleteUserById(id uint) error {
	err := repo.db.Model(&entities.User{}).Where("id = ?", id).Delete(&entities.User{}).Error
	if err != nil {
		return err
	}

	return nil
}
