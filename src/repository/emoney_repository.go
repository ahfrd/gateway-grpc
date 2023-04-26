package repository

import (
	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/emoney"
)

type EmoneyRepository interface {
	GetWalletProfile(bodyReq request.GeneralRequestBody) (*proto.GetWalletProfileResponse, error)
	GetWalletInfo(bodyReq request.GeneralRequestBody) (*proto.GetWalletInfoResponse, error)
	UpdateProfile(bodyReq request.UpdateProfileRequest) (*proto.UpdateProfileResponse, error)
	InsertWalletInfo(bodyReq request.InsertWalletInfoRequest) (*proto.InsertWalletInfoResponse, error)
	UpdatePremiumAccount(bodyReq request.GeneralRequestBody) (*proto.UpdateToPremiumAccountResponse, error)
}
