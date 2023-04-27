package service

import (
	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	"github.com/ahfrd/gateway-apps-grpc/src/model/response"
	"github.com/gin-gonic/gin"
)

type TopUpService interface {
	Form(ctx *gin.Context, bodyReq request.FormRequest) (*response.GeneralResponse, error)
	Inquiry(ctx *gin.Context, bodyReq request.InquiryRequest) (*response.GeneralResponse, error)
	Payment(ctx *gin.Context, bodyReq request.PaymentRequest) (*response.GeneralResponse, error)
}
