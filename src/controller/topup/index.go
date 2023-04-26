package topup

import (
	topupSvc "github.com/ahfrd/gateway-apps-grpc/src/service/topup"
)

type TopUpController struct {
	topupSvc.TopUpService
}

func NewTopUpController(topupService *topupSvc.TopUpService) TopUpController {
	return TopUpController{TopUpService: *topupService}
}
