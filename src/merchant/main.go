package main

import (
	"suregem/src/internal/merchant"
	merchantSchema "suregem/src/schemas/merchant"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	srv := handler.NewDefaultServer(
		merchantSchema.NewExecutableSchema(
			merchantSchema.Config{
				Resolvers: &merchant.Resolver{},
			},
		),
	)

	// GraphQL endpoint
	r.POST("/query", gin.WrapH(srv))
	r.GET("/query", gin.WrapH(srv))

	// Playground
	r.GET("/", gin.WrapH(playground.Handler("Merchant GraphQL", "/query")))

	r.Run(":8081")
}
