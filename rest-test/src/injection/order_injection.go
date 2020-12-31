package injection

import (
	"rest-test/src/controller"
	"rest-test/src/repository"
	"rest-test/src/service"

	"gorm.io/gorm"
)

func OrderInjection(db *gorm.DB) controller.OrderController {
	r := repository.ProvideOrderRepository(db)
	s := service.ProvideOrderService(r)
	return controller.ProvideOrderController(s)
}
