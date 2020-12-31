package controller

import (
	"log"
	"rest-test/src/entity"
	"rest-test/src/payload"
	"rest-test/src/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	os service.OrderService
}

func ProvideOrderController(os service.OrderService) OrderController {
	return OrderController{os: os}
}

func (oc *OrderController) CreateOrder(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	var order entity.Order
	if err := c.Bind(&order); err != nil {
		log.Println(err)
		return c.JSON(400, payload.ResponseWithMessage(false, err.Error()))
	}

	if len(order.OrderDetails) == 0 {
		return c.JSON(400, payload.ResponseWithMessage(false, "Input order_details, min=1"))
	}

	if err := c.Validate(&order); err != nil {
		log.Println(err)
		return c.JSON(400, payload.ResponseWithMessage(false, err.Error()))
	}

	order.CustomerID = claims["sub"].(string)

	if err := oc.os.CreateOrder(&order); err != nil {
		log.Println(err)
		return c.JSON(400, payload.ResponseWithMessage(false, err.Error()))
	}

	return c.JSON(200, payload.ResponseWithData(true, order))
}
