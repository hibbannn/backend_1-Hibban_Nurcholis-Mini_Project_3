package superadmin

import (
	"github.com/golang/mock/gomock"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/entities"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/mocks"
	"github.com/hibbannn/backend_1-Hibban_Nurcholis-Mini_Project_2/repositories"
	"reflect"
	"testing"
)

func TestUsecaseSuperadmin_ApprovedAdminRegister(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}
			if err := uc.ApprovedAdminRegister(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ApprovedAdminRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecaseSuperadmin_CreateSuperadmin(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		superadmin SuperAdminParam
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Case 1",
			fields: fields{superadminRepo: mockRepo},
			args: args{superadmin: SuperAdminParam{
				Username: "hibban",
				Password: "Nurcholis",
			}},
			want: &entities.Actor{
				Username: "hibban",
				Password: "Nurcholis",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}

			mockRepo.EXPECT().CreateSuperAdmin(gomock.Any()).Return(tt.want, nil)

			got, err := uc.CreateSuperadmin(tt.args.superadmin)
			got.Password = "123"
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSuperadmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSuperadmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_DeleteCustomerById(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		id uint
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "asdasd",
			fields:  fields{superadminRepo: mockRepo},
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}
			mockRepo.EXPECT().GetUserById(gomock.Any()).Return(&entities.User{}, nil)
			mockRepo.EXPECT().DeleteUserById(gomock.Any(), gomock.Any()).Return(nil).Times(1)

			if err := uc.DeleteUserByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecaseSuperadmin_GetAllAdmins(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		username string
		page     int
		pageSize int
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Case 1",
			fields: fields{superadminRepo: mockRepo},
			args: args{
				username: "dias",
				page:     2,
				pageSize: 6,
			},
			want:    []*entities.Actor{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}

			mockRepo.EXPECT().GetAllAdmins(tt.args.username, tt.args.page, tt.args.pageSize).Return(tt.want, nil)

			got, err := uc.GetAllAdmins(tt.args.username, tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllAdmins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllAdmins() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_GetAllCustomers(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		first_name string
		last_name  string
		email      string
		page       int
		pageSize   int
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Case 1",
			fields: fields{superadminRepo: mockRepo},
			args: args{
				first_name: "dias",
				last_name:  "pangestu",
				email:      "dias@gmail.com",
				page:       2,
				pageSize:   6,
			},
			want:    []*entities.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}

			mockRepo.EXPECT().GetAllUsers(tt.args.first_name, tt.args.last_name, tt.args.email, tt.args.page, tt.args.pageSize).Return(tt.want, nil)

			got, err := uc.GetAllUsers(tt.args.first_name, tt.args.last_name, tt.args.email, tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCustomers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_GetApprovalRequest(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		want    []*entities.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Case 1",
			fields:  fields{superadminRepo: mockRepo},
			want:    []*entities.Actor{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}

			mockRepo.EXPECT().GetApprovalRequests().Return(tt.want, nil)

			got, err := uc.GetApprovalRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApprovalRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApprovalRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_LoginSuperadmin(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		id       uint
		username string
		password string
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}

			mockRepo.EXPECT().LoginSuperAdmin(tt.args.username).Return(tt.want, nil)

			got, _, err := uc.LoginSuperadmin(tt.args.id, tt.args.username, tt.args.password)

			if (err != nil) != tt.wantErr {
				t.Errorf("LoginSuperadmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginSuperadmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecaseSuperadmin_RejectedAdminRegister(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		id uint
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Case 1",
			fields:  fields{superadminRepo: mockRepo},
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}

			mockRepo.EXPECT().RejectAdminRegistration(tt.args.id).Return(nil)

			if err := uc.RejectedAdminRegister(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("RejectedAdminRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//func TestUsecaseSuperadmin_UpdateActivedAdmin(t *testing.T) {
//	type fields struct {
//		superadminRepo repositories.SuperAdminRepository
//	}
//	type args struct {
//		id uint
//	}
//
//	ctrl := gomock.NewController(t)
//	mockRepo := mocks.NewMockSuperAdminRepository(ctrl)
//
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{
//			name:    "Test Case 1",
//			fields:  fields{superadminRepo: mockRepo},
//			args:    args{id: 1},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			uc := UsecaseSuperadmin{
//				superadminRepo: tt.fields.superadminRepo,
//			}
//
//			mockRepo.EXPECT().UpdateAdminActiveStatus(tt.args.id).Return(nil)
//
//			if err := uc.UpdateActivedAdmin(); (err != nil) != tt.wantErr {
//				t.Errorf("UpdateActivedAdmin() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

func TestUsecaseSuperadmin_UpdateDeadactivedAdmin(t *testing.T) {
	type fields struct {
		superadminRepo repositories.SuperAdminRepository
	}
	type args struct {
		id uint
	}

	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockSuperAdminRepository(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Case 1",
			fields:  fields{superadminRepo: mockRepo},
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := UsecaseSuperadmin{
				superadminRepo: tt.fields.superadminRepo,
			}

			mockRepo.EXPECT().UpdateDeadactivedAdmin(tt.args.id).Return(nil)

			if err := uc.UpdateDeadactivedAdmin(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDeadactivedAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
