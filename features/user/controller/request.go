package controller

import "mhub/features/user/entity"

type UserRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RequestMain(dataRequest UserRequest) entity.Main {
	return entity.Main{
		Id:       dataRequest.ID,
		Name:     dataRequest.Name,
		Email:    dataRequest.Email,
		Password: dataRequest.Password,
		Address:  dataRequest.Address,
		Phone:    dataRequest.Phone,
	}
}
