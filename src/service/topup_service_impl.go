package service

import (
	"encoding/json"
	"fmt"

	"github.com/ahfrd/gateway-apps-grpc/src/helpers"
	"github.com/ahfrd/gateway-apps-grpc/src/model/entity"
	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	"github.com/ahfrd/gateway-apps-grpc/src/model/response"
	"github.com/ahfrd/gateway-apps-grpc/src/repository"

	"github.com/gin-gonic/gin"
)

func NewTopUpService(topUpRepository *repository.TopUpRepository) TopUpService {
	return &topupServiceImpl{
		TopUpRepository: *topUpRepository,
	}
}

type topupServiceImpl struct {
	TopUpRepository repository.TopUpRepository
}

func (service *topupServiceImpl) Form(ctx *gin.Context, bodyReq request.FormRequest) (*response.GeneralResponse, error) {
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

func (service *topupServiceImpl) Inquiry(ctx *gin.Context, bodyReq request.InquiryRequest) (*response.GeneralResponse, error) {
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
func (service *topupServiceImpl) Payment(ctx *gin.Context, bodyReq request.PaymentRequest) (*response.GeneralResponse, error) {
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
