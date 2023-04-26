package emoney

import (
	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/emoney"

	"github.com/gin-gonic/gin"
)

type EmoneyService interface {
	GetWalletProfile(ctx *gin.Context, bodyReq request.GeneralRequestBody) (*proto.GetWalletProfileResponse, error)
	GetWalletInfo(ctx *gin.Context, bodyReq request.GeneralRequestBody) (*proto.GetWalletInfoResponse, error)
	UpdateProfile(ctx *gin.Context, bodyReq request.UpdateProfileRequest) (*proto.UpdateProfileResponse, error)
	InsertWalletInfo(ctx *gin.Context, bodyReq request.InsertWalletInfoRequest) (*proto.InsertWalletInfoResponse, error)
	UpdatePremiumAccount(ctx *gin.Context, bodyReq request.GeneralRequestBody) (*proto.UpdateToPremiumAccountResponse, error)
}
