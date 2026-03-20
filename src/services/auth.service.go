package services

import (
	"suregem/src/pkg/httpclient"
)

type AuthServices interface {
}

type authServiceImpl struct {
	client *httpclient.Client
}

func NewAuthService(client *httpclient.Client) AuthServices {
	return &authServiceImpl{
		client: client,
	}
}
