package repository

import (
	"context"

	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/emoney"

	"github.com/ahfrd/gateway-apps-grpc/src/model/request"
)

func NewEmoneyRepository(protoEmoney *proto.EmoneyServiceClient) EmoneyRepository {
	return &emoneyRepositoryImpl{
		ProtoEmoney: *protoEmoney,
	}
}

type emoneyRepositoryImpl struct {
	ProtoEmoney proto.EmoneyServiceClient
}

func (repository *emoneyRepositoryImpl) GetWalletProfile(bodyReq request.GeneralRequestBody) (*proto.GetWalletProfileResponse, error) {
	res, err := repository.ProtoEmoney.GetWalletProfile(context.Background(), &proto.GetWalletProfileRequest{
		PhoneNumber: bodyReq.PhoneNumber,
	})

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repository *emoneyRepositoryImpl) GetWalletInfo(bodyReq request.GeneralRequestBody) (*proto.GetWalletInfoResponse, error) {
	res, err := repository.ProtoEmoney.GetWalletInfo(context.Background(), &proto.GetWalletInfoRequest{
		PhoneNumber: bodyReq.PhoneNumber,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repository *emoneyRepositoryImpl) UpdateProfile(bodyReq request.UpdateProfileRequest) (*proto.UpdateProfileResponse, error) {
	res, err := repository.ProtoEmoney.UpdateProfile(context.Background(), &proto.UpdateProfileRequest{
		PhoneNumber: bodyReq.PhoneNumber,
		Name:        bodyReq.Name,
		Email:       bodyReq.Email,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repository *emoneyRepositoryImpl) InsertWalletInfo(bodyReq request.InsertWalletInfoRequest) (*proto.InsertWalletInfoResponse, error) {
	res, err := repository.ProtoEmoney.InsertWalletInfo(context.Background(), &proto.InsertWalletInfoRequest{
		PhoneNumber:  bodyReq.PhoneNumber,
		SecurityCode: bodyReq.SecurityCode,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (repository *emoneyRepositoryImpl) UpdatePremiumAccount(bodyReq request.GeneralRequestBody) (*proto.UpdateToPremiumAccountResponse, error) {
	res, err := repository.ProtoEmoney.UpdatePremiumAccount(context.Background(), &proto.UpdateToPremiumAccountRequest{
		PhoneNumber: bodyReq.PhoneNumber,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
