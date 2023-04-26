package auth

import (
	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/auth"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(ctx *gin.Context, bodyReq request.LoginRequestBody) (*proto.LoginResponse, error)
	Register(ctx *gin.Context, c proto.AuthServiceClient, bodyReq request.RegisterRequestBody) (*proto.RegisterResponse, error)
}
