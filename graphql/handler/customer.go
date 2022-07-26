package handler

import (
	"go-graphql/graphql/service"

	"github.com/gofiber/fiber/v2"
)

type (
	handler struct {
		srv service.Service
	}
)

func NewCustomerHandler(srv service.Service) handler {
	return handler{srv}
}

func (c handler) GetCustomer(f *fiber.Ctx) error {
	id, err := f.ParamsInt("id", 1)
	if err != nil {
		return err
	}

	customer, err := c.srv.GetCustomer(id)
	if err != nil {
		return err
	}

	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  customer,
		"error": nil,
	})
}

func (c handler) GetCustomers(f *fiber.Ctx) error {
	customers, err := c.srv.GetCustomers()
	if err != nil {
		return err
	}
	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  customers,
		"error": nil,
	})
}
