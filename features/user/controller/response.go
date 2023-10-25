package controller

import (
	user "mhub/features/user/entity"
	"time"
)

type UserLoginResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

func MainResponse(dataMain user.Main) UserResponse {

	userResponse := UserResponse{
		Id:        dataMain.Id,
		Name:      dataMain.Name,
		Email:     dataMain.Email,
		Address:   dataMain.Address,
		Phone:     dataMain.Phone,
		CreatedAt: dataMain.CreatedAt,
	}

	return userResponse
}

func LoginResponse(id, email, token string) UserLoginResponse {
	return UserLoginResponse{
		Id:    id,
		Email: email,
		Token: token,
	}
}
