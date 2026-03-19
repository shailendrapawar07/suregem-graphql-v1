package main

import (
	"suregem/src/internal/retailer"
	retailerSchema "suregem/src/schemas/retailer"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	srv := handler.NewDefaultServer(
		retailerSchema.NewExecutableSchema(
			retailerSchema.Config{
				Resolvers: &retailer.Resolver{},
			},
		),
	)

	r.POST("/query", gin.WrapH(srv))
	r.GET("/query", gin.WrapH(srv))

	r.GET("/", gin.WrapH(playground.Handler("Retailer GraphQL", "/query")))

	r.Run(":8082")
}
