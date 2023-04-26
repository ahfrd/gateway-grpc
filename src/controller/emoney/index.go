package emoney

import (
	emoneySvc "github.com/ahfrd/gateway-apps-grpc/src/service/emoney"
)

type EmoneyController struct {
	emoneySvc.EmoneyService
}

func NewEmoneyController(emoneyService *emoneySvc.EmoneyService) EmoneyController {
	return EmoneyController{EmoneyService: *emoneyService}
}
