package retailer

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

import (
	"suregem/src/config"
	"suregem/src/pkg/httpclient"
	"suregem/src/services"
)

type Resolver struct {
	cfg *config.Config
	// client      *httpclient.Client
	authService services.AuthServices
}

func NewResolver(cfg *config.Config) *Resolver {
	return &Resolver{
		cfg: cfg,
		// client:      httpclient.New(cfg.APIBaseURL),
		authService: services.NewAuthService(httpclient.New(cfg.APIBaseURL)),
	}
}
