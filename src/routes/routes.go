package routes

import (
	controller "github.com/ahfrd/gateway-apps-grpc/src/controller"
	repository "github.com/ahfrd/gateway-apps-grpc/src/repository"
	svc "github.com/ahfrd/gateway-apps-grpc/src/service"
)

type ControllerEntity struct {
	ControllerAuth   controller.AuthController
	ControllerEmoney controller.EmoneyController
	ControllerTopUp  controller.TopUpController
}

func InitController() *ControllerEntity {
	//Repository
	authRepository := repository.NewAuthRepository()
	emoneyRepository := repository.NewEmoneyRepository()

	//Service
	authSvc := svc.NewAuthService(&authRepository)
	emoneySvc := svc.NewEmoneyService(&emoneyRepository)
	topUpSvc := svc.NewTopUpService()
	//Controller
	authController := controller.NewAuthController(&authSvc)
	emoneyController := controller.NewEmoneyController(&emoneySvc)
	topUpController := controller.NewTopUpController(&topUpSvc)

	//Init
	controller := &ControllerEntity{
		ControllerAuth:   authController,
		ControllerEmoney: emoneyController,
		ControllerTopUp:  topUpController,
	}
	return controller
}
