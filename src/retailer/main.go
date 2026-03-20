package main

import (
	"fmt"
	"suregem/src/config"
	"suregem/src/internal/retailer"
	"suregem/src/middleware"
	retailerSchema "suregem/src/schemas/retailer"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	//load config
	cfg := config.Load()

	r := gin.Default()

	// Attach middleware here
	r.Use(middleware.ContextMiddleware())

	srv := handler.NewDefaultServer(
		retailerSchema.NewExecutableSchema(
			retailerSchema.Config{
				Resolvers: retailer.NewResolver(cfg),
			},
		),
	)

	r.POST("/query", gin.WrapH(srv))
	r.GET("/query", gin.WrapH(srv))

	r.GET("/", gin.WrapH(playground.Handler("Retailer GraphQL", "/query")))

	fmt.Println("GraphQL server: RETAILER")                 // 4002
	fmt.Println("Target server URL:=====>", cfg.APIBaseURL) // http://localhost:8080/api/v1
	fmt.Println("Is Production:=====>", cfg.IsProd())

	r.Run(":" + cfg.Merchant.Port)
}
