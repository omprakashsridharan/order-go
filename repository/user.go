package repository

import (
	"errors"
	"order-go/db"
	"order-go/model"

	"github.com/go-playground/validator/v10"
)

type IUserRepository interface {
	// UpdateUser(id uint, user model.User) (model.User, error)
	// DeleteUser(id string) (int64, error)
	// SelectUserWIthId(id string) (model.User, error)
	CreateUser(user model.User) (int64, error)
}

type Repository struct {
	Database db.Database
}

func (repo Repository) CreateUser(user model.User) (int64, error) {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return 0, err.(validator.ValidationErrors)
	}

	result := repo.Database.Create(&user)
	if result == 0 {
		return 0, errors.New("user not created")
	}
	return result, nil
}

// func (repo Repository) SelectUserWIthId(id string) (model.User, error) {
// 	var user model.User
// 	result := repo.Database.First(&user, "id = ?", id)
// 	if result.RowsAffected == 0 {
// 		return model.User{}, errors.New("user data not found")
// 	}
// 	return user, nil
// }

// func (repo Repository) UpdateUser(id uint, user model.User) (model.User, error) {
// 	var updateUser model.User
// 	result := repo.Database.Model(&updateUser).Where("id = ?", id).Updates(user)
// 	if result.RowsAffected == 0 {
// 		return model.User{}, errors.New("user data not updated")
// 	}
// 	return updateUser, nil
// }

// func (repo Repository) DeleteUser(id string) (int64, error) {
// 	var deletedUser model.User
// 	result := repo.Database.Where("id = ?", id).Delete(&deletedUser)
// 	if result.RowsAffected == 0 {
// 		return 0, errors.New("user data not update")
// 	}
// 	return result.RowsAffected, nil
// }
