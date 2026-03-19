package retailer

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

import (
	"suregem/src/config"
	"suregem/src/pkg/httpclient"
)

type Resolver struct {
	cfg    *config.Config
	client *httpclient.Client
}

func NewResolver(cfg *config.Config) *Resolver {
	return &Resolver{
		cfg:    cfg,
		client: httpclient.New(cfg.APIBaseURL),
	}
}
