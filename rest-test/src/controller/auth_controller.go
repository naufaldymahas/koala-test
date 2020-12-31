package controller

import (
	"encoding/base64"
	"log"
	"rest-test/src/entity"
	"rest-test/src/helper"
	"rest-test/src/payload"
	"rest-test/src/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	cr repository.CustomerRepository
}

func ProvideAuthController(cr repository.CustomerRepository) AuthController {
	return AuthController{cr: cr}
}

func (ac *AuthController) Register(c echo.Context) error {
	var request payload.RegisterRequest
	if err := c.Bind(&request); err != nil {
		log.Println(err)
		return c.JSON(400, payload.ResponseWithMessage(false, err.Error()))
	}

	if err := c.Validate(&request); err != nil {
		log.Println(err)
		return c.JSON(400, payload.ResponseWithMessage(false, err.Error()))
	}

	dob, err := time.Parse("2006-01-02", request.DOB)
	if err != nil {
		log.Println(err)
		return c.JSON(400, payload.ResponseWithMessage(false, "Date format must 'YYYY-MM-DD'"))
	}

	res, _ := ac.cr.FindByEmailOrPhoneNumber(request.Email, request.PhoneNumber)
	if res.Email != "" {
		return c.JSON(400, payload.ResponseWithMessage(false, "Email or Phone number already used!"))
	}

	salt := helper.GenerateRandomSalt(16)
	b64salt := base64.StdEncoding.EncodeToString(salt)

	customer := entity.Customer{
		CustomerName: request.CustomerName,
		Email:        request.Email,
		PhoneNumber:  request.PhoneNumber,
		DOB:          dob,
		Sex:          request.Sex,
		Salt:         b64salt,
		Password:     helper.HSHA256Handler(request.Password, salt),
	}

	if err := ac.cr.InsertCustomer(&customer); err != nil {
		log.Println(err)
		return c.JSON(400, payload.ResponseWithMessage(false, err.Error()))
	}

	return c.JSON(200, payload.ResponseWithMessage(true, "Customer created successfully"))
}

func (ac *AuthController) Login(c echo.Context) error {
	var request payload.LoginRequest
	if err := c.Bind(&request); err != nil {
		log.Println(err)
		return c.JSON(400, payload.ResponseWithMessage(false, err.Error()))
	}

	if err := c.Validate(&request); err != nil {
		log.Println(err)
		return c.JSON(400, payload.ResponseWithMessage(false, err.Error()))
	}

	customer, err := ac.cr.FindByEmailOrPhoneNumber(request.PhoneNumberOrEmail, request.PhoneNumberOrEmail)
	if err != nil {
		log.Println(err)
		return c.JSON(404, payload.ResponseWithMessage(false, "Customer not found"))
	}

	salt, _ := base64.StdEncoding.DecodeString(customer.Salt)
	if customer.Password != helper.HSHA256Handler(request.Password, salt) {
		return c.JSON(401, payload.ResponseWithMessage(false, "Wrong password"))
	}

	tokens := ac.generateTokens(customer.CustomerID)

	return c.JSON(200, payload.ResponseWithData(true, tokens))
}

func (ac *AuthController) RefreshToken(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	customer, err := ac.cr.FindByID(claims["sub"].(string))
	if err != nil {
		log.Println(err)
		return c.JSON(401, payload.ResponseWithMessage(false, "Customer not valid"))
	}

	tokens := ac.generateTokens(customer.CustomerID)

	return c.JSON(200, payload.ResponseWithData(true, tokens))
}

func (ac *AuthController) generateTokens(id string) payload.TokenResponse {
	t := helper.GenerateJWT(id, 0)
	r := helper.GenerateJWT(id, 1)
	return payload.TokenResponse{Token: t, RefreshToken: r}
}
