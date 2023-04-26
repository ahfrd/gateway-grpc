package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	authSvc "github.com/ahfrd/gateway-apps-grpc/src/service"

	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/auth"
	"github.com/ahfrd/gateway-apps-grpc/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthController struct {
	AuthService authSvc.AuthService
}

func NewAuthController(authService *authSvc.AuthService) AuthController {
	return AuthController{AuthService: *authService}
}

// func NewAuthController(authService *svc.AuthService) AuthController {
// 	return AuthController{AuthService: *authService}
// }

func (o *AuthController) Login(ctx *gin.Context, c proto.AuthServiceClient) {
	requestId := uuid.New()
	var bodyReq request.LoginRequestBody
	if err := ctx.BindJSON(&bodyReq); err != nil {
		utils.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		utils.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStart := utils.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := o.AuthService.Login(ctx, bodyReq)
	if err != nil {
		utils.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		utils.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := utils.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}

func (o *AuthController) Register(ctx *gin.Context, c proto.AuthServiceClient) {
	requestId := uuid.New()
	var bodyReq request.RegisterRequestBody
	if err := ctx.BindJSON(&bodyReq); err != nil {
		utils.LogError(ctx, err.Error(), requestId.String())
		return
	}
	requestData, err := json.Marshal(bodyReq)
	logStart := utils.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)
	if err != nil {
		utils.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	response, err := o.AuthService.Register(ctx, c, bodyReq)
	if err != nil {

		utils.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	responseData, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		utils.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := utils.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}
