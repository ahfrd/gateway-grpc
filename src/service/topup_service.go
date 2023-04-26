package service

import (
	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	"github.com/ahfrd/gateway-apps-grpc/src/model/response"
	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/topup"
	"github.com/gin-gonic/gin"
)

type TopUpService interface {
	Form(ctx *gin.Context, c proto.TopUpServiceClient, bodyReq request.FormRequest) (*response.GeneralResponse, error)
	Inquiry(ctx *gin.Context, c proto.TopUpServiceClient, bodyReq request.InquiryRequest) (*response.GeneralResponse, error)
}
