package emoney

import (
	"net/http"

	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/emoney"
	"github.com/ahfrd/gateway-apps-grpc/src/repository"
	"github.com/gin-gonic/gin"
)

func NewEmoneyService(emoneyRepository *repository.EmoneyRepository) EmoneyService {
	return &emoneyServiceImpl{
		EmoneyRepository: *emoneyRepository,
	}
}

type emoneyServiceImpl struct {
	EmoneyRepository repository.EmoneyRepository
}

func (service *emoneyServiceImpl) GetWalletProfile(ctx *gin.Context, bodyReq request.GeneralRequestBody) (*proto.GetWalletProfileResponse, error) {
	res, err := service.EmoneyRepository.GetWalletProfile(bodyReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil, err
	}
	return res, nil
}

func (service *emoneyServiceImpl) GetWalletInfo(ctx *gin.Context, bodyReq request.GeneralRequestBody) (*proto.GetWalletInfoResponse, error) {
	res, err := service.EmoneyRepository.GetWalletInfo(bodyReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil, err
	}
	return res, nil
}

func (service *emoneyServiceImpl) UpdateProfile(ctx *gin.Context, bodyReq request.UpdateProfileRequest) (*proto.UpdateProfileResponse, error) {
	res, err := service.EmoneyRepository.UpdateProfile(bodyReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil, err
	}
	return res, nil
}

func (service *emoneyServiceImpl) InsertWalletInfo(ctx *gin.Context, bodyReq request.InsertWalletInfoRequest) (*proto.InsertWalletInfoResponse, error) {
	res, err := service.EmoneyRepository.InsertWalletInfo(bodyReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil, err
	}
	return res, nil
}
func (service *emoneyServiceImpl) UpdatePremiumAccount(ctx *gin.Context, bodyReq request.GeneralRequestBody) (*proto.UpdateToPremiumAccountResponse, error) {
	res, err := service.EmoneyRepository.UpdatePremiumAccount(bodyReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return nil, err
	}
	return res, nil
}
