package resolver

import (
	"fmt"
	"go-graphql/graphql/model"
	"go-graphql/graphql/service"

	"github.com/graphql-go/graphql"
)

type (
	customerResolver struct {
		CustomerService service.Service
	}
	CustomerResolver interface {
		GetCustomer(params graphql.ResolveParams) (any, error)
		GetCustomers(params graphql.ResolveParams) (any, error)
		CreateCustomer(params graphql.ResolveParams) (any, error)
	}
)

func NewCustomerResolver(cus service.Service) CustomerResolver {
	return customerResolver{CustomerService: cus}
}

func (c customerResolver) GetCustomer(params graphql.ResolveParams) (any, error) {
	var (
		id int
		ok bool
	)
	if id, ok = params.Args["id"].(int); !ok || id == 0 {
		return nil, fmt.Errorf("id is not integr or zero")
	}
	cutomer, err := c.CustomerService.GetCustomer(id)
	if err != nil {
		return nil, err
	}
	return cutomer, nil
}

func (c customerResolver) GetCustomers(params graphql.ResolveParams) (any, error) {
	customers, err := c.CustomerService.GetCustomers()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c customerResolver) CreateCustomer(params graphql.ResolveParams) (any, error) {
	var (
		cusID       int
		cusName     string
		dateOfBirth string
		city        string
		zipCode     string
		status      int
		ok          bool
	)
	if cusID, ok = params.Args["CustomerID"].(int); !ok || cusID == 0 {
		return nil, fmt.Errorf("id is not int or 0")
	}
	if cusName, ok = params.Args["Name"].(string); !ok || cusName == "" {
		return nil, fmt.Errorf("name is not stirng or nil")
	}
	if dateOfBirth, ok = params.Args["DateOfBirth"].(string); !ok || dateOfBirth == "" {
		return nil, fmt.Errorf("date of birth is not stirng or nil")
	}
	if city, ok = params.Args["City"].(string); !ok || city == "" {
		return nil, fmt.Errorf("city is not string or nil")
	}
	if zipCode, ok = params.Args["ZipCode"].(string); !ok || zipCode == "" {
		return nil, fmt.Errorf("zip code is not string or nil")
	}
	if status, ok = params.Args["Status"].(int); !ok || status == 0 {
		return nil, fmt.Errorf("status is not int or 0")
	}

	newCus := model.CustomerRes{
		CustomerID:  cusID,
		Name:        cusName,
		DateOfBirth: dateOfBirth,
		City:        city,
		ZipCode:     zipCode,
		Status:      status,
	}
	res, err := c.CustomerService.CreateCustomer(&newCus)
	if err != nil {
		return nil, err
	}
	return res, nil
}
