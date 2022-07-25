package main

import (
	"fmt"
	"go-graphql/graphql/repository"
	"go-graphql/graphql/service"
	"go-graphql/pkg/database"
)

func main() {
	db := database.GetDB()
	repo := repository.NewCustomerRepository(db)
	srv := service.NewCustomerService(repo)
	cur, _ := srv.GetCustomer(1)

	fmt.Printf("%v", cur)
}
