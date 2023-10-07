package usecase

import (
	"errors"
	"reflect"
	"strconv"
	"testing"

	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/mock"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestAdminHome(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	adminRepo := mock.NewMockAdminRepository(ctrl)
	adminUseCase := NewAdminUseCase(adminRepo)

	tests := []struct {
		name           string
		input          uint
		beforeTest     func(mock.MockAdminRepository, uint)
		expectedOutput models.AdminDetails
		expectedErr    error
	}{
		{
			name:  "successful admin home",
			input: 3,
			beforeTest: func(mar mock.MockAdminRepository, id uint) {
				mar.EXPECT().AdminDetails(id).Times(1).Return(models.AdminDetails{
					ID:    3,
					Name:  "aarathy",
					Email: "aarathy@gmail.com",
					PhNo:  "+919745503907",
				}, nil)
			},
			expectedOutput: models.AdminDetails{
				ID:    3,
				Name:  "aarathy",
				Email: "aarathy@gmail.com",
				PhNo:  "+919745503907",
			},
			expectedErr: nil,
		},
		{
			name:  "error in fetching admin details",
			input: 3,
			beforeTest: func(mar mock.MockAdminRepository, id uint) {
				mar.EXPECT().AdminDetails(id).Times(1).Return(models.AdminDetails{}, errors.New("error in fetching admin details"))
			},
			expectedOutput: models.AdminDetails{},
			expectedErr:    errors.New("error in fetching admin details"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(*adminRepo, tt.input)

			got, err := adminUseCase.AdminHome(tt.input)
			assert.Equal(t, err, tt.expectedErr)
			if !reflect.DeepEqual(got, tt.expectedOutput) {
				t.Errorf("adminUseCase_AdminHome()=%v want %v", got, tt.expectedOutput)
			}
		})
	}

}

func TestListUsers(t *testing.T) {
	ctrl := gomock.NewController(t)

	adminRepo := mock.NewMockAdminRepository(ctrl)
	adminUseCase := NewAdminUseCase(adminRepo)

	type args struct {
		pages  string
		counts string
	}
	tests := []struct {
		name       string
		input      args
		beforeTest func(mock.MockAdminRepository)
		want       []models.AdminUserResponse
		wantErr    error
	}{
		{
			name:  "listing users",
			input: args{pages: "2", counts: "1"},
			beforeTest: func(mar mock.MockAdminRepository) {
				mar.EXPECT().ListUsers(2, 1).Times(1).Return([]models.AdminUserResponse{{
					ID:     2,
					Name:   "aarathy",
					Email:  "aarathy@gmail.com",
					PhNo:   "+919745503907",
					Status: false,
				}}, nil)
			},
			want: []models.AdminUserResponse{{
				ID:     2,
				Name:   "aarathy",
				Email:  "aarathy@gmail.com",
				PhNo:   "+919745503907",
				Status: false,
			}},
			wantErr: nil,
		},
		{
			name:  "error in listing",
			input: args{pages: "2", counts: "1"},
			beforeTest: func(mar mock.MockAdminRepository) {
				mar.EXPECT().ListUsers(2, 1).Times(1).Return(nil, errors.New("error in fetching user list"))
			},
			want:    nil,
			wantErr: errors.New("error in fetching user list"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(*adminRepo)

			got, err := adminUseCase.ListUsers(tt.input.pages, tt.input.counts)

			assert.Equal(t, err, tt.wantErr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("adminUseCase_ListUsers()=%v want %v", got, tt.want)
			}
		})
	}
}

func TestBlockUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	adminRepo := mock.NewMockAdminRepository(ctrl)
	adminUseCase := NewAdminUseCase(adminRepo)

	
	tests := []struct {
		name       string
		input string
		beforeTest func(mock.MockAdminRepository, string)
		wantErr    error
	}{
		{
			name:  "block user",
			input: "1",
			beforeTest: func(mar mock.MockAdminRepository, id string) {
				userId,_:=strconv.Atoi(id)
				gomock.InOrder(
					mar.EXPECT().IsBlocked(uint(userId)).Times(1).Return(false),
					mar.EXPECT().BlockUser(uint(userId)).Times(1).Return(nil),
				)
			},
			wantErr: nil,
		},
		{
			name:  "already blocked user",
			input: "1",
			beforeTest: func(mar mock.MockAdminRepository, id string) {
				userId,_:=strconv.Atoi(id)
				gomock.InOrder(
					mar.EXPECT().IsBlocked(uint(userId)).Times(1).Return(true),
					mar.EXPECT().BlockUser(uint(userId)).Times(0),
				)
			},
			wantErr: errors.New("already blocked user"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(*adminRepo, tt.input)

			err := adminUseCase.BlockUser(tt.input)

			assert.Equal(t, err, tt.wantErr)
		})
	}
}
