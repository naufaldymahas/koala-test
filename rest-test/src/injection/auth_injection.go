package injection

import (
	"rest-test/src/controller"
	"rest-test/src/repository"

	"gorm.io/gorm"
)

func AuthInjection(db *gorm.DB) controller.AuthController {
	r := repository.ProvideCustomerRepository(db)
	return controller.ProvideAuthController(r)
}
