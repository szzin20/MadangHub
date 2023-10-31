package responses

import (
	"mhub/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}

// UserList mengembalikan respons daftar pengguna
func UserList(c echo.Context, users []models.User) error {
	structuredUsers := make([]UserResponse, len(users))
	for i, user := range users {
		structuredUsers[i] = UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Address:  user.Address,
			Role:     user.Role,
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all users",
		"users":   structuredUsers,
	})
}

func UserResponses(c echo.Context, id uint, username, email, address, role string) error {
    response := UserResponse{
        ID:       id,
        Username: username,
        Email:    email,
        Address:  address,
        Role:     role,
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Success: get user by ID",
        "user":    response,
    })
}
