package repository

import (
	"go-graphql/graphql/entity"

	"gorm.io/gorm"
)

type (
	Repository interface {
		GetAll() ([]entity.Customer, error)
		GetByID(id int) (*entity.Customer, error)
		CreateCustomer(customer *entity.Customer) (entity.Customer, error)
	}
	repository struct {
		db *gorm.DB
	}
)

func NewCustomerRepository(db *gorm.DB) Repository {
	return repository{db}
}

func (r repository) GetAll() ([]entity.Customer, error) {
	var customers []entity.Customer
	err := r.db.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r repository) GetByID(id int) (*entity.Customer, error) {
	var customer entity.Customer
	err := r.db.Where("id = ?", id).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r repository) CreateCustomer(customer *entity.Customer) (entity.Customer, error) {
	err := r.db.Create(customer).Error
	if err != nil {
		return entity.Customer{}, err
	}
	return *customer, nil
}
