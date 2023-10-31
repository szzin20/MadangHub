package responses

import (
	"mhub/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderResponse struct {
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`
	Items  []struct {
		FoodID   uint `json:"food_id"`
		Quantity int  `json:"quantity"`
	} `json:"items"`
	TotalPrice float64 `json:"total_price"`
}

// BuildOrderResponses takes a slice of order models and builds the simplified order responses.
func BuildOrderResponses(orders []models.Order) []OrderResponse {
	simplifiedOrders := make([]OrderResponse, len(orders))

	for i, order := range orders {
		simplifiedOrders[i].ID = order.ID
		simplifiedOrders[i].UserID = order.UserID
		simplifiedOrders[i].Items = make([]struct {
			FoodID   uint `json:"food_id"`
			Quantity int  `json:"quantity"`
		}, len(order.Items))
		simplifiedOrders[i].TotalPrice = order.TotalPrice

		for j, item := range order.Items {
			simplifiedOrders[i].Items[j].FoodID = item.FoodID
			simplifiedOrders[i].Items[j].Quantity = item.Quantity
		}
	}

	return simplifiedOrders
}

// SuccessOrderResponse sends a successful order response.
func SuccessOrderResponse(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": message,
		"data":    data,
	})
}
