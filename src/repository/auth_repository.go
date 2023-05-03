package repository

import (
	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/auth"
)

type AuthRepository interface {
	Login(bodyReq *request.LoginRequestBody) (*proto.LoginResponse, error)
	Register(bodyReq *request.RegisterRequestBody) (*proto.RegisterResponse, error)
}
