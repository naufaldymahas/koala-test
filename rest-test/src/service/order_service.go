package service

import (
	"fmt"
	"time"

	"rest-test/src/entity"
	"rest-test/src/repository"
)

type OrderService struct {
	or repository.OrderRepository
}

func ProvideOrderService(or repository.OrderRepository) OrderService {
	return OrderService{or: or}
}

func (os *OrderService) CreateOrder(order *entity.Order) error {
	m := int(time.Now().Month())
	var lastValMonth int
	os.or.DB.Raw("SELECT last_value FROM order_month_seq").Scan(&lastValMonth)
	if m != lastValMonth {
		os.or.DB.Exec("SELECT setval('order_month_seq', ?, true)", m)
		os.or.DB.Exec("ALTER SEQUENCE order_ref_seq RESTART WITH 123")
	}

	y := time.Now().Year()
	var orderNumber string
	os.or.DB.Raw(fmt.Sprintf("SELECT 'PO-' || (SELECT nextval('order_ref_seq')) || '/' || to_char(%d, 'FMRN') || '/' || %d", m, y)).Scan(&orderNumber)
	order.OrderNumber = orderNumber

	return os.or.CreateOrder(order)
}
