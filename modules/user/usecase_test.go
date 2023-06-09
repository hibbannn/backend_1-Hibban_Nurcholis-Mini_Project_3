package user

import (
	"github.com/golang/mock/gomock"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/entities"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(mockCtrl)

	usecase := NewUsecaseUser(mockRepo)

	userParam := UserParam{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Avatar:    "avatar.jpg",
	}
	mockRepo.EXPECT().CreateUser(gomock.Any()).Return(&entities.User{
		Id:        1,
		FirstName: userParam.FirstName,
		LastName:  userParam.LastName,
		Email:     userParam.Email,
		Avatar:    userParam.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil)

	createdUser, err := usecase.CreateUser(userParam)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, uint(1), createdUser.Id)
	assert.Equal(t, "John", createdUser.FirstName)
	assert.Equal(t, "Doe", createdUser.LastName)
	assert.Equal(t, "john@example.com", createdUser.Email)
	assert.Equal(t, "avatar.jpg", createdUser.Avatar)
}

func TestGetUserById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(mockCtrl)
	usecase := NewUsecaseUser(mockRepo)

	expectedUser := &entities.User{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Avatar:    "avatar.jpg",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.EXPECT().GetUserById(expectedUser.Id).Return(expectedUser, nil)

	user, err := usecase.GetUserById(expectedUser.Id)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.Id, user.Id)
	assert.Equal(t, expectedUser.FirstName, user.FirstName)
	assert.Equal(t, expectedUser.LastName, user.LastName)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.Equal(t, expectedUser.Avatar, user.Avatar)
}

func TestUpdateUserById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(mockCtrl)
	usecase := NewUsecaseUser(mockRepo)

	userParam := UserParam{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Avatar:    "avatar.jpg",
	}

	existingUser := &entities.User{
		Id:        1,
		FirstName: "Existing",
		LastName:  "User",
		Email:     "existing@example.com",
		Avatar:    "existing-avatar.jpg",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.EXPECT().GetUserById(existingUser.Id).Return(existingUser, nil)
	mockRepo.EXPECT().UpdateUserById(existingUser.Id, gomock.Any()).Return(existingUser, nil)

	updatedUser, err := usecase.UpdateUserById(existingUser.Id, userParam)

	assert.NoError(t, err)
	assert.NotNil(t, updatedUser)
	assert.Equal(t, existingUser.Id, updatedUser.Id)
	assert.Equal(t, userParam.FirstName, updatedUser.FirstName)
	assert.Equal(t, userParam.LastName, updatedUser.LastName)
	assert.Equal(t, userParam.Email, updatedUser.Email)
	assert.Equal(t, userParam.Avatar, updatedUser.Avatar)
}

func TestDeleteUserById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(mockCtrl)
	usecase := NewUsecaseUser(mockRepo)

	userID := uint(1)

	mockRepo.EXPECT().DeleteUserById(userID).Return(nil)

	err := usecase.DeleteUserById(userID)

	assert.NoError(t, err)
}
