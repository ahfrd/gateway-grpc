package service

import (
	"context"
	"net/http"

	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/auth"
	"github.com/ahfrd/gateway-apps-grpc/src/repository"
	"github.com/gin-gonic/gin"
)

func NewAuthService(authRepository *repository.AuthRepository) AuthService {
	return &authServiceImpl{
		AuthRepositry: *authRepository,
	}
}

type authServiceImpl struct {
	AuthRepositry repository.AuthRepository
}

func (service *authServiceImpl) Login(ctx *gin.Context, bodyReq request.LoginRequestBody) (*proto.LoginResponse, error) {
	res, err := service.AuthRepositry.Login(bodyReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}
	return res, nil
}

func (service *authServiceImpl) Register(ctx *gin.Context, c proto.AuthServiceClient, bodyReq request.RegisterRequestBody) (*proto.RegisterResponse, error) {
	res, err := c.Register(context.Background(), &proto.RegisterRequest{
		PhoneNumber: bodyReq.PhoneNumber,
		Password:    bodyReq.Password,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
