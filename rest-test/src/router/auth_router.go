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

func AuthRouter(e *echo.Echo, db *gorm.DB) {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	SECRET := os.Getenv("SALT_JWT")

	ac := injection.AuthInjection(db)

	auth := e.Group("/auth")
	auth.POST("/register", ac.Register)
	auth.POST("/login", ac.Login)

	auth.Use(middleware.JWT([]byte(SECRET)))
	auth.PATCH("/refresh", ac.RefreshToken)
}
