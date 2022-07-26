package schema

import "github.com/graphql-go/graphql"

var (
	Customer = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Customer",
			Fields: graphql.Fields{
				"customer_id":   &graphql.Field{Type: graphql.Int},
				"name":          &graphql.Field{Type: graphql.String},
				"date_of_birth": &graphql.Field{Type: graphql.String},
				"city":          &graphql.Field{Type: graphql.String},
				"zip_code":      &graphql.Field{Type: graphql.String},
				"status":        &graphql.Field{Type: graphql.Int},
			},
		},
	)
)
