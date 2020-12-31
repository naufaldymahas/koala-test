package repository

import (
	"rest-test/src/entity"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func ProvideCustomerRepository(db *gorm.DB) CustomerRepository {
	return CustomerRepository{db: db}
}

func (cr *CustomerRepository) InsertCustomer(c *entity.Customer) error {
	if err := cr.db.Create(c).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CustomerRepository) FindByEmailOrPhoneNumber(e, p string) (entity.Customer, error) {
	var customer entity.Customer
	if err := cr.db.Where("email = ? OR phone_number = ?", e, p).Take(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (cr *CustomerRepository) FindByID(id string) (entity.Customer, error) {
	var customer entity.Customer
	if err := cr.db.Where("customer_id = ?", id).Take(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}
