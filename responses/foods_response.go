package responses

import (
	"mhub/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FoodResponse struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"`
}

// FoodsResponse adalah respons untuk daftar makanan
type FoodsResponse struct {
	Message string         `json:"message"`
	Foods   []FoodResponse `json:"foods"`
}

// FoodsList mengembalikan respons daftar makanan tanpa CreatedAt, UpdatedAt, dan DeletedAt
func FoodsList(c echo.Context, foods []models.Food) error {
	foodResponses := make([]FoodResponse, len(foods))
	for i, food := range foods {
		foodResponses[i] = FoodResponse{
			ID:          food.ID,
			Title:       food.Title,
			Description: food.Description,
			Price:       food.Price,
			Status:      food.Status,
		}
	}

	response := FoodsResponse{
		Message: "Success: get all foods",
		Foods:   foodResponses,
	}

	return c.JSON(http.StatusOK, response)
}

func FoodByID(c echo.Context, food models.Food) error {
    response := FoodResponse{
        ID:          food.ID,
        Title:       food.Title,
        Description: food.Description,
        Price:       food.Price,
        Status:      food.Status,
    }

    return c.JSON(http.StatusOK, response)
}