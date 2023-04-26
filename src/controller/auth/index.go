package auth

import (
	authSvc "github.com/ahfrd/gateway-apps-grpc/src/service/auth"
)

type AuthController struct {
	AuthService authSvc.AuthService
}

func NewAuthController(authService *authSvc.AuthService) AuthController {
	return AuthController{AuthService: *authService}
}
