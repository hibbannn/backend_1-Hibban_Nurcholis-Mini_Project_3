package admin

import (
	"github.com/golang/mock/gomock"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/entities"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/mocks"
	"testing"
	"time"
)

func TestUsecaseAdmin_LoginAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAdminRepositoryInterface(ctrl)
	//mockHelpers := mock_helpers.NewMockHelpersInterface(ctrl)

	usecaseAdmin := UsecaseAdmin{
		AdminRepo: mockRepo,
		//Helpers:   mockHelpers, // Gunakan objek mockHelpers di use case

	}

	// Mock input and output
	username := "admin"
	password := "password"
	admin := &entities.Actor{
		Id:        1,
		Username:  "username",
		Password:  "hashedPassword",
		RoleId:    2,
		Verified:  false,
		Active:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tokenString := "token"

	// Mock expected behavior
	mockRepo.EXPECT().LoginAdmin(username).Return(admin, nil)
	//mockHelpers.EXPECT().ComparePass(gomock.Any(), gomock.Any()).Return(true, nil)
	//mockHelpers.EXPECT().GenerateToken(admin.Id, username).Return(tokenString, nil)

	// Call the use case method
	resultAdmin, resultToken, err := usecaseAdmin.LoginAdmin(username, password)
	if err != nil {
		t.Errorf("LoginAdmin returned an unexpected error: %v", err)
		return
	}

	// Verify the result
	if resultAdmin != admin {
		t.Errorf("LoginAdmin returned unexpected admin. Expected: %v, got: %v", admin, resultAdmin)
	}
	if resultToken != tokenString {
		t.Errorf("LoginAdmin returned unexpected token. Expected: %s, got: %s", tokenString, resultToken)
	}
}
