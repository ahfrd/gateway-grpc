package repository

import (
	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/topup"
)

type TopUpRepository interface {
	Form(bodyReq *request.FormRequest) (*proto.GeneralResponse, error)
	Inquiry(bodyReq *request.InquiryRequest) (*proto.GeneralResponse, error)
	Payment(bodyReq *request.PaymentRequest) (*proto.GeneralResponse, error)
}
