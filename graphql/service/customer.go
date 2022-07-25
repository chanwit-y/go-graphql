package service

import (
	"go-graphql/graphql/entity"
	"go-graphql/graphql/model"
	"go-graphql/graphql/repository"
)

type (
	Service interface {
		GetCustomer(id int) (model.CustomerRes, error)
		GetCustomers() ([]model.CustomerRes, error)
		CreateCustomer(customer *model.CustomerRes) (model.CustomerRes, error)
	}
	service struct {
		repository repository.Repository
	}
)

func NewCustomerService(repository repository.Repository) Service {
	return service{repository}
}

func (c service) GetCustomer(id int) (model.CustomerRes, error) {
	customer, err := c.repository.GetByID(id)
	if err != nil {
		return model.CustomerRes{}, err
	}
	return model.CustomerRes{
		CustomerID:  customer.CustomerID,
		Name:        customer.Name,
		DateOfBirth: customer.DateOfBirth,
		City:        customer.City,
		ZipCode:     customer.ZipCode,
		Status:      customer.Status,
	}, nil
}

func (c service) GetCustomers() ([]model.CustomerRes, error) {
	customers, err := c.repository.GetAll()
	if err != nil {
		return nil, err
	}
	var customerRes []model.CustomerRes
	for _, v := range customers {
		customerRes = append(customerRes, model.CustomerRes{
			CustomerID:  v.CustomerID,
			Name:        v.Name,
			DateOfBirth: v.DateOfBirth,
			City:        v.City,
			ZipCode:     v.ZipCode,
			Status:      v.Status,
		})
	}
	return customerRes, nil
}

func (c service) CreateCustomer(customer *model.CustomerRes) (model.CustomerRes, error) {
	newCus := entity.Customer{
		CustomerID:  customer.CustomerID,
		Name:        customer.Name,
		DateOfBirth: customer.DateOfBirth,
		City:        customer.City,
		ZipCode:     customer.ZipCode,
		Status:      customer.Status,
	}
	res, err := c.repository.CreateCustomer(&newCus)
	if err != nil {
		return model.CustomerRes{}, err
	}
	return model.CustomerRes{
		CustomerID:  res.CustomerID,
		Name:        res.Name,
		DateOfBirth: res.DateOfBirth,
		City:        res.City,
		ZipCode:     res.ZipCode,
		Status:      res.Status,
	}, nil
}
