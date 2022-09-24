package repository

import (
	"order-go/db"
	"order-go/mocks"
	"order-go/model"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestRepository_CreateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		Database func(u model.User) db.Database
	}
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Create user successfully",
			fields: fields{
				Database: func(u model.User) db.Database {
					mockDB := mocks.NewMockDatabase(mockCtrl)
					mockDB.EXPECT().Create(&u).Return(int64(1)).Times(1)
					return mockDB
				},
			},
			args: args{
				user: model.User{
					Email:    "test@test.com",
					Password: "password",
					Role:     "ADMIN",
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Create user failed Email missing",
			fields: fields{
				Database: func(u model.User) db.Database {
					mockDB := mocks.NewMockDatabase(mockCtrl)
					return mockDB
				},
			},
			args: args{
				user: model.User{
					Password: "password",
					Role:     "ADMIN",
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Create user failed Invalid email",
			fields: fields{
				Database: func(u model.User) db.Database {
					mockDB := mocks.NewMockDatabase(mockCtrl)
					return mockDB
				},
			},
			args: args{
				user: model.User{
					Email:    "email",
					Password: "password",
					Role:     "ADMIN",
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Create user failed Password missing",
			fields: fields{
				Database: func(u model.User) db.Database {
					mockDB := mocks.NewMockDatabase(mockCtrl)
					return mockDB
				},
			},
			args: args{
				user: model.User{
					Email: "test@email.com",
					Role:  "ADMIN",
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Create user failed Role missing",
			fields: fields{
				Database: func(u model.User) db.Database {
					mockDB := mocks.NewMockDatabase(mockCtrl)
					return mockDB
				},
			},
			args: args{
				user: model.User{
					Email:    "test@email.com",
					Password: "password",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repository{
				Database: tt.fields.Database(tt.args.user),
			}
			got, err := repo.CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repository.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
