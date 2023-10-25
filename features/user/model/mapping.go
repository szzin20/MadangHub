package model

import "mhub/features/user/entity"



func MapMainToModel(mainData entity.Main) User {
	return User{
		Name:     mainData.Name,
		Email:    mainData.Email,
		Password: mainData.Password,
		Address:  mainData.Address,
		Phone:    mainData.Phone,
		Role:     mainData.Role,
	}
}

func MapModelToMain(mainData User) entity.Main {
	return entity.Main{
		Id:        mainData.Id,
		Name:      mainData.Name,
		Email:     mainData.Email,
		Password:  mainData.Password,
		Address:   mainData.Address,
		Phone:     mainData.Phone,
		Role:      mainData.Role,
		CreatedAt: mainData.CreatedAt,
		UpdatedAt: mainData.UpdatedAt,
	}
}
