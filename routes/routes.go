package routes

import (
	"mhub/middlewares"
	"mhub/controllers"

	"github.com/labstack/echo/v4"
)

func InitmyRoutes() *echo.Echo {
	e := echo.New()
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

	return e
}
