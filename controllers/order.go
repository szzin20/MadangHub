package controller

import (
	"mhub/app/config"
	"mhub/middlewares"
	"mhub/models"
	"mhub/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetAllOrders digunakan untuk mendapatkan semua data pesanan.
func GetAllOrders(c echo.Context) error {
	_, role := middlewares.ExtractToken(c)

	if role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	var orders []models.Order
	if err := config.DB.Preload("Items.Foods").Find(&orders).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve orders"})
	}

	simplifiedOrders := responses.BuildOrderResponses(orders)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully retrieved all orders",
		"orders":  simplifiedOrders,
	})
}

// Endpoint untuk membuat pesanan
func CreateOrder(c echo.Context) error {
	// Bind data pesanan dari JSON
	order := models.Order{}
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid order data",
			"error":   err.Error(),
		})
	}

	// Menerima daftar makanan dalam pesanan
	foodsInOrder := []models.Food{}
	totalPrice := 0.0

	for _, item := range order.Items {
		food := models.Food{}
		if err := config.DB.First(&food, item.FoodID).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid food ID"})
		}
		foodsInOrder = append(foodsInOrder, food)
		totalPrice += food.Price * float64(item.Quantity)
	}

	// Hitung total harga dan set data pesanan
	order.TotalPrice = totalPrice

	// Simpan pesanan dalam basis data
	if err := config.DB.Create(&order).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create order",
			"error":   err.Error(),
		})
	}

	// Sederhanakan respons dan atur ulang data pesanan
	simplifiedOrder := struct {
		ID        uint   `json:"id"`
		UserID    uint   `json:"user_id"`
		Address   string `json:"address"`
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
		Items     []struct {
			FoodID   uint `json:"food_id"`
			Quantity int  `json:"quantity"`
		} `json:"items"`
		TotalPrice float64 `json:"total_price"`
	}{
		ID:        order.ID,
		UserID:    order.UserID,
		Address:   order.Address,
		Longitude: order.Longitude,
		Latitude:  order.Latitude,
		Items: make([]struct {
			FoodID   uint `json:"food_id"`
			Quantity int  `json:"quantity"`
		}, len(order.Items)),
		TotalPrice: order.TotalPrice,
	}

	for i, item := range order.Items {
		simplifiedOrder.Items[i].FoodID = item.FoodID
		simplifiedOrder.Items[i].Quantity = item.Quantity
	}

	// Mengirim respons HTTP berhasil setelah order berhasil dibuat
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Order successfully created",
		"order":   simplifiedOrder,
	})
}

func GetOrderByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid order ID"})
	}

	userId, role := middlewares.ExtractToken(c)
	if role != "admin" && userId != uint(id) {
		return c.JSON(http.StatusForbidden, map[string]string{"message": "Access denied"})
	}

	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Order not found"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success: get order by ID",
		"order":   order,
	})
}
