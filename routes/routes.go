package routes

import (
	controller "mhub/controllers"
	"mhub/middlewares"

	"github.com/labstack/echo/v4"
)

func InitmyRoutes() *echo.Echo {
	e := echo.New()

	e.POST("/recommend-food", func(c echo.Context) error {
        return controller.RecommendFood(c, controller.NewFoodUsecase())
    })

	middlewares.LogMiddleware(e)

	userGroup := e.Group("/users")
	userGroup.POST("/register", controller.RegisterUser)
	userGroup.POST("/login", controller.LoginUserController)
	userGroup.GET("", controller.GetAllUsers, middlewares.JWTMiddleware())
	userGroup.GET("/:id", controller.GetUserByID, middlewares.JWTMiddleware())
	userGroup.PUT("/:id", controller.UpdateUserByID, middlewares.JWTMiddleware())
	userGroup.DELETE("/:id", controller.DeleteUser, middlewares.JWTMiddleware())

	foodGroup := e.Group("/foods")
	foodGroup.POST("", controller.RegisterFoods, middlewares.JWTMiddleware())
	foodGroup.GET("", controller.GetAllFoods, middlewares.JWTMiddleware())
	foodGroup.GET("/:id", controller.GetFoodByID, middlewares.JWTMiddleware())
	foodGroup.PUT("/:id", controller.UpdateFoodByID, middlewares.JWTMiddleware())
	foodGroup.DELETE("/:id", controller.DeleteFood, middlewares.JWTMiddleware())

	orderGroup := e.Group("/orders")
	orderGroup.POST("", controller.CreateOrder, middlewares.JWTMiddleware())
	orderGroup.GET("", controller.GetAllOrders, middlewares.JWTMiddleware())
	orderGroup.GET("/:id", controller.GetOrderByID, middlewares.JWTMiddleware())



	return e
}
