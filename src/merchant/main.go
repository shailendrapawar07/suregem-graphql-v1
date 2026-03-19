package main

import (
	"fmt"
	"suregem/src/config"
	"suregem/src/internal/merchant"
	merchantSchema "suregem/src/schemas/merchant"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {

	//load configs
	cfg := config.Load()

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

	fmt.Println("GraphQL server: MERCHENT")             // 4002
	fmt.Println("Target server URL:=>", cfg.APIBaseURL) // http://localhost:8080/api/v1
	fmt.Println("Is Production:=>", cfg.IsProd())

	r.Run(":" + cfg.Merchant.Port)
}
