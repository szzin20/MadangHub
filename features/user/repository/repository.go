package repository

import (
	"errors"
	"mhub/features/user/entity"
	"mhub/features/user/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// Login implements entity.UserDataInterface.
func (*userRepository) Login(email string, username string, password string) (entity.Main, string, error) {
	panic("unimplemented")
}

// Register implements entity.UserDataInterface.
func (*userRepository) Register(data entity.Main) (row int, err error) {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) entity.UserDataInterface {
	return &userRepository{
		db: db,
	}
}

// Create implements user.UserDataInterface.
func (u *userRepository) Create(user entity.Main) (string, error) {
	dataInput := model.MapMainToModel(user)

	tx := u.db.Create(&dataInput)
	if tx.Error != nil {
		return "", tx.Error
	}

	return dataInput.Id, nil
}

// CheckLogin implements user.UserDataInterface.
func (u *userRepository) CheckLogin(email string, password string) (entity.Main, error) {
	var data model.User

	tx := u.db.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return entity.Main{}, tx.Error
	}
	dataMain := model.MapModelToMain(data)
	return dataMain, nil
}

// GetById implements user.UserDataInterface.
func (u *userRepository) GetById(id string) (entity.Main, error) {
	var userData model.User

	result := u.db.Preload("Pickups").Where("id = ?", id).First(&userData)
	if result.Error != nil {
		return entity.Main{}, result.Error
	}

	var userById = model.MapModelToMain(userData)
	return userById, nil
}

// UpdateById implements user.UserDataInterface.
func (u *userRepository) UpdateById(id string, updated entity.Main) (data entity.Main, err error) {
	var userData model.User
	resultFind := u.db.First(&userData, id)
	if resultFind.Error != nil {
		return entity.Main{}, resultFind.Error
	}

	u.db.Model(&userData).Updates(model.MapMainToModel(updated))

	data = model.MapModelToMain(userData)
	return data, nil
}

// Delete implements user.UserDataInterface.
func (u *userRepository) DeleteById(id string) error {
	result := u.db.Where("id = ?", id).Delete(&model.User{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("failed to delete data, record not found")
	}

	return nil
}

// FindAllUsers implements user.UserDataInterface.
// func (u *userRepository) FindAllUsers() ([]entity.Main, error) {
// 	var user []model.User

// 	err := u.db.Preload("Pickups").Find(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	var allUser []entity.Main = model.ModelToMainMapping(user)

// 	return allUser, nil
// }
