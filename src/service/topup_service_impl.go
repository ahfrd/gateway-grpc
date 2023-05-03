package service

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ahfrd/gateway-apps-grpc/src/helpers"
	"github.com/ahfrd/gateway-apps-grpc/src/model/entity"
	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	"github.com/ahfrd/gateway-apps-grpc/src/model/response"
	"github.com/ahfrd/gateway-apps-grpc/src/repository"

	"github.com/gin-gonic/gin"
)

func NewTopUpService(topUpRepository *repository.TopUpRepository, emoneyRepository *repository.EmoneyRepository) TopUpService {
	return &topupServiceImpl{
		EmoneyRepository: *emoneyRepository,
		TopUpRepository:  *topUpRepository,
	}
}

type topupServiceImpl struct {
	EmoneyRepository repository.EmoneyRepository
	TopUpRepository  repository.TopUpRepository
}

func (service *topupServiceImpl) Form(ctx *gin.Context, bodyReq *request.FormRequest) (*response.GeneralResponse, error) {
	res, err := service.TopUpRepository.Form(bodyReq)
	resData := response.GeneralResponse{}
	if err != nil {
		return nil, err
	}

	dataValue := []entity.FormStructEntity{}
	json.Unmarshal(res.GetResponseData().Value, &dataValue)
	fmt.Println(dataValue)
	resData.Code = helpers.RCSuksesCode
	resData.Msg = res.ResponseMessage
	resData.Data = dataValue
	return &resData, nil
}

func (service *topupServiceImpl) Inquiry(ctx *gin.Context, bodyReq *request.InquiryRequest) (*response.GeneralResponse, error) {
	res, err := service.TopUpRepository.Inquiry(bodyReq)
	resData := response.GeneralResponse{}
	if err != nil {
		return nil, err
	}
	dataValue := entity.FormStructEntity{}
	json.Unmarshal(res.GetResponseData().Value, &dataValue)
	resData.Code = helpers.RCSuksesCode
	resData.Msg = res.ResponseMessage
	resData.Data = dataValue
	return &resData, nil
}
func (service *topupServiceImpl) Payment(ctx *gin.Context, bodyReq *request.PaymentRequest) (*response.GeneralResponse, error) {
	generalRequest := &request.GeneralRequestBody{
		PhoneNumber: bodyReq.PhoneNumb,
	}
	selectWalletInfo, err := service.EmoneyRepository.GetWalletInfo(generalRequest)
	if err != nil {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  err.Error(),
		}, nil
	}
	if selectWalletInfo.RessponseCode != 202 {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  selectWalletInfo.ResponseMessage,
		}, nil
	}
	ballance := selectWalletInfo.ResponseData.Ballance
	fee, _ := strconv.Atoi(bodyReq.Fee)
	nominalTopUp, _ := strconv.Atoi(bodyReq.Nominal)
	intBallance, _ := strconv.Atoi(ballance)
	total := fee + nominalTopUp
	if intBallance < total {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  "Maaf saldo anda tidak cukup",
		}, nil
	}
	res, err := service.TopUpRepository.Payment(bodyReq)
	resData := response.GeneralResponse{}
	if err != nil {
		return nil, err
	}
	dataValue := entity.FormStructEntity{}
	json.Unmarshal(res.GetResponseData().Value, &dataValue)
	resData.Code = helpers.RCSuksesCode
	resData.Msg = res.ResponseMessage
	resData.Data = dataValue
	return &resData, nil
}
