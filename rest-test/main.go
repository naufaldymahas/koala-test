package main

import (
	"fmt"
	"log"
	"os"
	"rest-test/src/config"
	"rest-test/src/helper"
	"rest-test/src/router"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	db := config.InitDB()
	log.Println("Connected to database")

	e := echo.New()
	e.Validator = &helper.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Recover())
	router.AuthRouter(e, db)
	router.OrderRouter(e, db)

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	e.Logger.Fatal(e.Start(PORT))
}
