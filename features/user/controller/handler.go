package controller

// import (
// 	user "mhub/features/user/entity"
// 	"mhub/utils/helper"
// 	"net/http"
// 	"strings"

// 	"github.com/labstack/echo/v4"
// )

// type UserController struct {
// 	userUseCase user.UseCaseInterface
// }

// func NewUserControllers(uc user.UseCaseInterface) *UserController {
// 	return &UserController{
// 		userUseCase: uc,
// 	}
// }

// func (uco *UserController) CreateUser(c echo.Context) error {
// 	// Bind data
// 	dataInput := UserRequest{}
// 	errBind := c.Bind(&dataInput)
// 	if errBind != nil {
// 		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
// 	}

// 	data := RequestMain(dataInput)

// 	errCreate := uco.userUseCase.Create(data)
// 	if errCreate != nil {
// 		if strings.Contains(errCreate.Error(), "validation") {
// 			return c.JSON(http.StatusBadRequest, helper.FailedResponse(errCreate.Error()))
// 		} else {
// 			return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed to create data"))
// 		}
// 	}

// 	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
// }