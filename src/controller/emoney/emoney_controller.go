package emoney

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	"github.com/ahfrd/gateway-apps-grpc/src/utils"
	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
)

func (o *EmoneyController) GetWalletProfile(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.GeneralRequestBody
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

	response, err := o.EmoneyService.GetWalletProfile(ctx, bodyReq)
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

func (o *EmoneyController) GetWalletInfo(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.GeneralRequestBody
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

	response, err := o.EmoneyService.GetWalletInfo(ctx, bodyReq)
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

func (o *EmoneyController) UpdateProfile(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.UpdateProfileRequest
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

	response, err := o.EmoneyService.UpdateProfile(ctx, bodyReq)
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
func (o *EmoneyController) InsertWalletInfo(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.InsertWalletInfoRequest
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

	response, err := o.EmoneyService.InsertWalletInfo(ctx, bodyReq)
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
func (o *EmoneyController) UpdatePremiumAccount(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.GeneralRequestBody
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

	response, err := o.EmoneyService.UpdatePremiumAccount(ctx, bodyReq)
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
