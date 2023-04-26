package topup

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ahfrd/gateway-apps-grpc/src/helpers"
	"github.com/ahfrd/gateway-apps-grpc/src/model/entity"
	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	"github.com/ahfrd/gateway-apps-grpc/src/model/response"

	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/topup"
	"github.com/gin-gonic/gin"
)

func NewTopUpService() TopUpService {
	return &topupServiceImpl{}
}

type topupServiceImpl struct {
}

func (service *topupServiceImpl) Form(ctx *gin.Context, c proto.TopUpServiceClient, bodyReq request.FormRequest) (*response.GeneralResponse, error) {
	res, err := c.Form(context.Background(), &proto.FormRequest{
		PhoneNumber: bodyReq.PhoneNumber,
	})
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

func (service *topupServiceImpl) Inquiry(ctx *gin.Context, c proto.TopUpServiceClient, bodyReq request.InquiryRequest) (*response.GeneralResponse, error) {
	res, err := c.Inquiry(context.Background(), &proto.InquiryRequest{
		PhoneNumber: bodyReq.PhoneNumb,
	})
	fmt.Println(",,msm")
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
func (service *topupServiceImpl) Payment(ctx *gin.Context, c proto.TopUpServiceClient, bodyReq request.FormRequest) (*response.GeneralResponse, error) {
	res, err := c.Payment(context.Background(), &proto.PaymentRequest{
		PhoneNumber: bodyReq.PhoneNumber,
	})
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
