package repository

import (
	"context"

	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/topup"
)

func NewTopUpRepository() TopUpRepository {
	return &topUpRepositoryImpl{}
}

type topUpRepositoryImpl struct {
	ProtoTopUp proto.TopUpServiceClient
}

func (repository *topUpRepositoryImpl) Form(bodyReq *request.FormRequest) (*proto.GeneralResponse, error) {
	payload := &proto.FormRequest{
		PhoneNumber: bodyReq.PhoneNumber,
	}
	res, err := repository.ProtoTopUp.Form(context.Background(), payload)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (repository *topUpRepositoryImpl) Inquiry(bodyReq *request.InquiryRequest) (*proto.GeneralResponse, error) {
	payload := &proto.InquiryRequest{
		Method:       bodyReq.Method,
		PhoneNumber:  bodyReq.PhoneNumb,
		NominalTopUp: bodyReq.Nominal,
		FeeTopUp:     bodyReq.Fee,
		CardInfo: &proto.DebitCardRequest{
			NomorKartu:  bodyReq.Card.CardNumb,
			ValidUntil:  bodyReq.Card.Valid,
			SecurityNum: bodyReq.Card.SecurityNum,
		},
		CodeNumber: bodyReq.CodeNumb,
		BankCode:   bodyReq.BankCode,
	}
	res, err := repository.ProtoTopUp.Inquiry(context.Background(), payload)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (repository *topUpRepositoryImpl) Payment(bodyReq *request.PaymentRequest) (*proto.GeneralResponse, error) {
	payload := &proto.PaymentRequest{
		Method:       bodyReq.Method,
		PhoneNumber:  bodyReq.PhoneNumb,
		NominalTopUp: bodyReq.Nominal,
		FeeTopUp:     bodyReq.Fee,
		CardInfo: &proto.DebitCardRequest{
			NomorKartu:  bodyReq.Card.CardNumb,
			ValidUntil:  bodyReq.Card.Valid,
			SecurityNum: bodyReq.Card.SecurityNum,
		},
		CodeNumber: bodyReq.CodeNumb,
	}
	res, err := repository.ProtoTopUp.Payment(context.Background(), payload)
	if err != nil {
		return nil, err
	}
	return res, nil
}
