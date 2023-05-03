package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	topupSvc "github.com/ahfrd/gateway-apps-grpc/src/service"

	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/topup"
	"github.com/ahfrd/gateway-apps-grpc/src/utils"
	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
)

type TopUpController struct {
	topupSvc.TopUpService
}

func NewTopUpController(topupService *topupSvc.TopUpService) TopUpController {
	return TopUpController{TopUpService: *topupService}
}
func (o *TopUpController) Form(ctx *gin.Context, c proto.TopUpServiceClient) {
	requestId := guuid.New()
	var bodyReq request.FormRequest
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

	response, err := o.TopUpService.Form(ctx, &bodyReq)
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
func (o *TopUpController) Inquiry(ctx *gin.Context, c proto.TopUpServiceClient) {
	requestId := guuid.New()
	var bodyReq request.InquiryRequest
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

	response, err := o.TopUpService.Inquiry(ctx, &bodyReq)
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
func (o *TopUpController) Payment(ctx *gin.Context, c proto.TopUpServiceClient) {
	requestId := guuid.New()
	var bodyReq request.InquiryRequest
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

	response, err := o.TopUpService.Inquiry(ctx, &bodyReq)
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
