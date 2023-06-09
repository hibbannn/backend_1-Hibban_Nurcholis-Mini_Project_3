package repositories

import (
	"encoding/json"
	"errors"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/entities"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"
)

type AdminRepositoryInterface interface {
	// Admin
	CreateApproval(adminId uint) (*entities.Approval, error)
	GetApprovalById(id uint) (*entities.Approval, error)
	LoginAdmin(username string) (*entities.Actor, error)
	RegisterAdmin(admin *entities.Actor) (*entities.Actor, error)

	// API Integration
	FetchUsersFromAPI() ([]*entities.User, error)
	SaveUsersFromAPI(url string) error

	// User
	CreateUser(user *entities.User) (*entities.User, error)
	DeleteUserById(id uint, user *entities.User) error
	GetAllUsers(firstName, lastName, email string, page, pageSize int) ([]*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	GetUserById(id uint) (*entities.User, error)
}

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return AdminRepository{
		db: db,
	}
}

func (repo AdminRepository) LoginAdmin(username string) (*entities.Actor, error) {
	admin := &entities.Actor{}

	err := repo.db.Model(&entities.Actor{}).Where("username = ? AND verified = ? AND active = ?", username, true, true).First(admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (repo AdminRepository) RegisterAdmin(admin *entities.Actor) (*entities.Actor, error) {
	err := repo.db.Model(&entities.Actor{}).Create(admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (repo AdminRepository) CreateUser(user *entities.User) (*entities.User, error) {
	err := repo.db.Model(&entities.User{}).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo AdminRepository) GetUserById(id uint) (*entities.User, error) {
	user := &entities.User{}

	err := repo.db.Model(&entities.User{}).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo AdminRepository) DeleteUserById(id uint, user *entities.User) error {
	err := repo.db.Model(&entities.User{}).Where("id = ?", id).Delete(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo AdminRepository) GetAllUsers(firstName, lastName, email string, page, pageSize int) ([]*entities.User, error) {
	var users []*entities.User

	query := repo.db.Model(&entities.User{})

	if firstName != "" {
		query = query.Where("first_name LIKE ?", "%"+firstName+"%")
	} else if lastName != "" {
		query = query.Where("last_name LIKE ?", "%"+lastName+"%")
	} else if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	offset := (page - 1) * pageSize

	err := query.Offset(offset).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo AdminRepository) GetUserByEmail(email string) (*entities.User, error) {
	user := &entities.User{}

	err := repo.db.Model(&entities.User{}).Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

type GetUsersResponse struct {
	Users []*entities.User `json:"data"`
}

func (repo AdminRepository) SaveUsersFromAPI(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			// Tangani error penutupan body di sini
		}
	}()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	userAPIResponse := new(GetUsersResponse)

	err = json.Unmarshal(body, userAPIResponse)
	if err != nil {
		return err
	}

	for _, user := range userAPIResponse.Users {
		_, err := repo.GetUserByEmail(user.Email)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				newUser := &entities.User{
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Email:     user.Email,
					Avatar:    user.Avatar,
				}
				_, err = repo.CreateUser(newUser)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			// User already exists, skip saving
		}
	}

	return nil
}

func (repo AdminRepository) CreateApproval(adminId uint) (*entities.Approval, error) {
	approval := &entities.Approval{
		AdminId:      adminId,
		SuperAdminId: 1,
		Status:       false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := repo.db.Model(&entities.Approval{}).Create(approval).Error
	if err != nil {
		return nil, err
	}

	return approval, nil
}

func (repo AdminRepository) GetApprovalById(id uint) (*entities.Approval, error) {
	approval := &entities.Approval{}

	err := repo.db.Model(&entities.Approval{}).Where("id = ?", id).First(approval).Error
	if err != nil {
		return nil, err
	}

	return approval, nil
}
