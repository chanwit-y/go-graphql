package database

import (
	"go-graphql/graphql/entity"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	dsn := "sqlserver://sa:abcABC123@localhost?database=graphql_demo"
	db, err := gorm.Open(sqlserver.Open(dsn))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Customer{})

	return db
}
