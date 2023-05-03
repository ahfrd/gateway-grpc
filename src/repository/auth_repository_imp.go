package repository

import (
	"context"
	"fmt"

	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/auth"

	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
)

func NewAuthRepository() AuthRepository {
	return &authRepositoryImpl{}
}

type authRepositoryImpl struct {
	ProtoClient proto.AuthServiceClient
}

func (repository *authRepositoryImpl) Login(bodyReq *request.LoginRequestBody) (*proto.LoginResponse, error) {
	payload := &proto.LoginRequest{
		PhoneNumber: bodyReq.PhoneNumber,
		Password:    bodyReq.Password,
	}
	res, err := repository.ProtoClient.Login(context.Background(), payload)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repository authRepositoryImpl) Register(bodyReq *request.RegisterRequestBody) (*proto.RegisterResponse, error) {
	fmt.Println(bodyReq)
	payload := &proto.RegisterRequest{
		PhoneNumber: bodyReq.PhoneNumber,
		Password:    bodyReq.Password,
	}
	res, err := repository.ProtoClient.Register(context.Background(), payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}
