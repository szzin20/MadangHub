package service

import (
	"errors"
	"mhub/utils/helper"
	"regexp"
	user "mhub/features/user/entity"
)

type userUsecase struct {
	userRepository user.UserDataInterface
}

// Register implements user.UseCaseInterface.
func (uc *userUsecase) Register(data user.Main) (row int, err error) {
    if data.Email == "" || data.Password == "" {
        return 0, errors.New("error, email or password can't be empty")
    }

    emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    match, _ := regexp.MatchString(emailRegex, data.Email)
    if !match {
        return 0, errors.New("error. invalid email format")
    }

    erruser, _ := uc.userRepository.Register(data)
    return erruser, nil
}


// Login implements user.UseCaseInterface.
func (uc *userUsecase) Login(email string, username string, password string) (user.Main, string, error) {
	if email == "" || password == "" {
		return user.Main{}, "", errors.New("error, email or password can't be empty")
	}

	logindata, token, err := uc.userRepository.Login(email, username, password)

	if err != nil {
		// Handle the error from the repository
		return user.Main{}, "", err
	}

	if helper.CheckPasswordHash(logindata.Password, password) {
		if err != nil {
			return user.Main{}, "", err
		}

		return logindata, token, nil
	}

	return user.Main{}, "", errors.New("Login Failed")
}

func New(UserUcase user.UserDataInterface) user.UseCaseInterface {
	return &userUsecase{
		userRepository: UserUcase,
	}
}
