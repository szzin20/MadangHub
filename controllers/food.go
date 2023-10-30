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

func RegisterFoods(c echo.Context) error {
	food := models.Food{}
	if err := c.Bind(&food); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid food data",
			"error":   err.Error(),
		})
	}

	_, role := middlewares.ExtractToken(c)

	if role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	var existingFood models.Food
	err := config.DB.Where("title = ?", food.Title).First(&existingFood).Error
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": constants.ErrFoodAlreadyExists,
		})
	}

	err = config.DB.Create(&food).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": constants.ErrFailedToCreate,
			"error":   err.Error(),
		})
	}

	// Mengirim respons HTTP berhasil setelah pengguna berhasil didaftarkan
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create new food",
		"food":    food,
	})
}

// GetAllfoods digunakan untuk mendapatkan semua data pengguna.
func GetAllFoods(c echo.Context) error {
	var foods []models.Food
	if err := config.DB.Find(&foods).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve foods"})
	}

	return responses.FoodsList(c, foods)
}

// GetUserByID digunakan untuk mendapatkan data pengguna berdasarkan ID.
func GetFoodByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid food ID"})
	}

	var food models.Food
	if err := config.DB.First(&food, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Foods not found"})
	}

	return responses.FoodByID(c, food)
}

// Fungsi UpdateFoodByID digunakan untuk memperbarui data pengguna berdasarkan ID.
func UpdateFoodByID(c echo.Context) error {
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	_, role := middlewares.ExtractToken(c)

	var existingFood models.Food
	if err := config.DB.First(&existingFood, Id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "food not found")
	}

	if role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	food := new(models.Food)
	if err := c.Bind(food); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	existingFood.Title = food.Title
	existingFood.Description = food.Description
	existingFood.Price = food.Price
	existingFood.Status = food.Status

	if err := config.DB.Save(&existingFood).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status" : "Success",
		"message": "Success update food",
	})
}

// Menghapus data user berdasarkan ID
func DeleteFood(c echo.Context) error {
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Food ID")
	}

	_, role := middlewares.ExtractToken(c)

	var Food models.Food
	if err := config.DB.First(&Food, Id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Food Not Found")
	}

	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, "Access denied")
	}

	if err := config.DB.Delete(&Food).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status" : "Success",
		"message": "Success delete food",
	})
}
