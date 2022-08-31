package main

import (
	"go-graphql/graphql/handler"
	"go-graphql/graphql/repository"
	"go-graphql/graphql/resolver"
	"go-graphql/graphql/schema"
	"go-graphql/graphql/service"
	"go-graphql/pkg/database"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/graphql-go/graphql"
	gqlHandler "github.com/graphql-go/handler"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func main() {
	db := database.GetDB()
	repo := repository.NewCustomerRepository(db)
	srv := service.NewCustomerService(repo)
	hdl := handler.NewCustomerHandler(srv)

	cResolver := resolver.NewCustomerResolver(srv)
	cSchema := schema.NewCustomerSchema(cResolver)
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    cSchema.Query(),
		Mutation: cSchema.Mutation(),
	})

	gh := gqlHandler.New(&gqlHandler.Config{
		Schema:   &graphqlSchema,
		GraphiQL: true,
		Pretty:   true,
	})
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${method} ${path}",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Bangkok",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept,Authorization",
	}))

	app.Get("/customer/:id", hdl.GetCustomer)
	app.Get("/customers", hdl.GetCustomers)

	app.Get("/graph", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gh.ServeHTTP(w, r)
		})(c.Context())
		return nil
	})
	app.Post("/graph", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gh.ServeHTTP(w, r)
		})(c.Context())
		return nil
	})

	errAppListen := app.Listen(":3000")
	if errAppListen != nil {
		panic(err)
	}
}
