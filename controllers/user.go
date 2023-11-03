package controller

import (
	"mhub/app/config"
	"mhub/constants"
	"mhub/middlewares"
	"mhub/models"
	"mhub/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Fungsi RegisterUser digunakan untuk mendaftarkan pengguna baru.
func RegisterUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid user data",
			"error":   err.Error(),
		})
	}

	// Cek apakah pengguna sudah terdaftar berdasarkan alamat email
	var existingUser models.User
	err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": constants.ErrUserAlreadyExists,
		})
	}

	// Jika pengguna belum terdaftar, simpan data pendaftaran ke dalam basis data
	err = config.DB.Create(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": constants.ErrFailedToRegister,
			"error":   err.Error(),
		})
	}

	// Mengirim respons HTTP berhasil setelah pengguna berhasil didaftarkan
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"message": "Success create new user",
	})
}

// Fungsi LoginUserController digunakan untuk mengautentikasi pengguna dan memberikan token akses jika berhasil.
func LoginAdminController(c echo.Context) error {
	// Membuat instance pengguna dan mengikat data dari permintaan HTTP
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Fail to parse request body",
			"error":   err.Error(),
		})
	}

	// Mencari pengguna dalam basis data berdasarkan email dan kata sandi
	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": constants.ErrFailedToLogIn,
			"error":   err.Error(),
		})
	}

	token, err := middlewares.CreateToken(user.ID, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": constants.ErrTokenCreationFailed,
			"error":   err.Error(),
		})
	}
	UserResponse := models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success login",
		"user":    UserResponse,
	})
}

func LoginUserController(c echo.Context) error {
	// Membuat instance pengguna dan mengikat data dari permintaan HTTP
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Fail to parse request body",
			"error":   err.Error(),
		})
	}

	// Mencari pengguna dalam basis data berdasarkan email dan kata sandi
	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": constants.ErrFailedToLogIn,
			"error":   err.Error(),
		})
	}

	token, err := middlewares.CreateToken(user.ID, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": constants.ErrTokenCreationFailed,
			"error":   err.Error(),
		})
	}
	UserResponse := models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success login",
		"user":    UserResponse,
	})
}

// GetAllUsers digunakan untuk mendapatkan semua data pengguna.
func GetAllUsers(c echo.Context) error {
	_, role := middlewares.ExtractToken(c)

	if role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve users"})
	}

	return responses.UserList(c, users)
}

// GetUserByID digunakan untuk mendapatkan data pengguna berdasarkan ID.
func GetUserByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	userId, role := middlewares.ExtractToken(c)
	if role != "admin" && userId != uint(id) {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}

	return responses.UserResponses(c, user.ID, user.Username, user.Email, user.Address, user.Role)
}

// Fungsi UpdateUserByID digunakan untuk memperbarui data pengguna berdasarkan ID.
func UpdateUserByID(c echo.Context) error {
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	userId, role := middlewares.ExtractToken(c)

	var existingUser models.User
	if err := config.DB.First(&existingUser, Id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	if role != "admin" && userId != existingUser.ID {
		return echo.NewHTTPError(http.StatusForbidden, "Access denied")
	}

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	existingUser.Address = user.Address

	if err := config.DB.Save(&existingUser).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success",
		"message": "Success update user",
	})
}

// Menghapus data user berdasarkan ID
func DeleteUser(c echo.Context) error {
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid User ID")
	}

	userId, role := middlewares.ExtractToken(c)

	var User models.User
	if err := config.DB.First(&User, Id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User Not Found")
	}

	if role != "admin" && uint(userId) != User.ID {
		return echo.NewHTTPError(http.StatusForbidden, "Access denied")
	}

	if err := config.DB.Delete(&User).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success",
		"message": "Success delete user",
	})
}
