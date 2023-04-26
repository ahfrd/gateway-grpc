package routes

import (
	"fmt"

	"github.com/ahfrd/gateway-apps-grpc/config"
	controller "github.com/ahfrd/gateway-apps-grpc/src/controller/emoney"
	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/emoney"
	repository "github.com/ahfrd/gateway-apps-grpc/src/repository"
	svc "github.com/ahfrd/gateway-apps-grpc/src/service/emoney"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClientEmoney struct {
	Client proto.EmoneyServiceClient
}
type ControllerEmoney struct {
	Controller controller.EmoneyController
}

func InitServiceClientEmoney(c *config.Config) proto.EmoneyServiceClient {
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("Could not connect : ", err)
	}
	return proto.NewEmoneyServiceClient(cc)
}
func InitControllerEmoney() *ControllerEmoney {
	var emoneyProto proto.EmoneyServiceClient
	emoneyRepository := repository.NewEmoneyRepository(&emoneyProto)
	emoneySvc := svc.NewEmoneyService(&emoneyRepository)
	emoneyController := controller.NewEmoneyController(&emoneySvc)
	controller := &ControllerEmoney{
		Controller: emoneyController,
	}
	return controller
}
func RoutesEmoney(r *gin.Engine, c *config.Config, authSvc *ServiceClient) {

	svc := &ServiceClientEmoney{
		Client: InitServiceClientEmoney(c),
	}

	routes := r.Group("/emoney")
	routes.POST("/wallet-profile", svc.UpdateProfile)
	routes.POST("/wallet-info", svc.GetWalletInfo)
	routes.POST("/update-profile", svc.GetWalletProfile)
	routes.POST("/update-premium", svc.UpdatePremium)
	routes.POST("/insert-wallet-info", svc.InsertWalletInfo)
}

func (ServiceClientEmoney *ServiceClientEmoney) UpdateProfile(ctx *gin.Context) {
	routing := InitControllerEmoney()
	routing.Controller.UpdateProfile(ctx)
}

func (ServiceClientEmoney *ServiceClientEmoney) GetWalletInfo(ctx *gin.Context) {
	routing := InitControllerEmoney()
	routing.Controller.GetWalletInfo(ctx)
}

func (ServiceClientEmoney *ServiceClientEmoney) GetWalletProfile(ctx *gin.Context) {
	routing := InitControllerEmoney()
	routing.Controller.GetWalletProfile(ctx)
}
func (ServiceClientEmoney *ServiceClientEmoney) UpdatePremium(ctx *gin.Context) {
	routing := InitControllerEmoney()
	routing.Controller.UpdatePremiumAccount(ctx)
}
func (ServiceClientEmoney *ServiceClientEmoney) InsertWalletInfo(ctx *gin.Context) {
	routing := InitControllerEmoney()
	routing.Controller.InsertWalletInfo(ctx)
}
