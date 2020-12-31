package router

import (
	"log"
	"os"
	"rest-test/src/injection"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func OrderRouter(e *echo.Echo, db *gorm.DB) {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	SECRET := os.Getenv("SALT_JWT")

	oc := injection.OrderInjection(db)

	order := e.Group("/order")
	order.Use(middleware.JWT([]byte(SECRET)))
	order.POST("", oc.CreateOrder)
}
