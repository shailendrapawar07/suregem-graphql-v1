package services

import (
	"context"
	"errors"
	"net/http"
	"suregem/src/pkg/httpclient"
	"suregem/src/schemas/retailer"
)

type AuthServices interface {
	Login(ctx context.Context, input retailer.LoginInput) (interface{}, *http.Response, error)
}

type authServiceImpl struct {
	client *httpclient.Client
}

func NewAuthService(client *httpclient.Client) AuthServices {
	return &authServiceImpl{
		client: client,
	}
}

func (s *authServiceImpl) Login(ctx context.Context, input retailer.LoginInput) (interface{}, *http.Response, error) {

	var responseBody map[string]interface{}

	res, err := s.client.Post("/user/login", &retailer.LoginInput{
		Email:    input.Email,
		Password: input.Password,
	}, &responseBody)

	// 1: handle network  level error
	if err != nil {
		return nil, res, errors.New(err.Error())
	}

	// fmt.Println("result in resolver====>")
	// fmt.Println(responseBody)

	return responseBody, res, nil
}
