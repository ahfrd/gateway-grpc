package routes

import (
	"fmt"

	"github.com/ahfrd/gateway-apps-grpc/config"
	controller "github.com/ahfrd/gateway-apps-grpc/src/controller/topup"
	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/topup"
	svc "github.com/ahfrd/gateway-apps-grpc/src/service/topup"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClientTopUp struct {
	Client proto.TopUpServiceClient
}
type ControllerTopUp struct {
	Controller controller.TopUpController
}

func InitServiceClientTopUp(c *config.Config) proto.TopUpServiceClient {
	fmt.Println(c.TopUpSvcUrl)
	fmt.Println(",,,,")
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial(c.TopUpSvcUrl, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("Could not connect : ", err)
	}
	return proto.NewTopUpServiceClient(cc)
}
func InitControllerTopUp() *ControllerTopUp {
	topUpSvc := svc.NewTopUpService()
	topUpController := controller.NewTopUpController(&topUpSvc)
	controller := &ControllerTopUp{
		Controller: topUpController,
	}
	return controller
}
func RoutesTopUp(r *gin.Engine, c *config.Config, authSvc *ServiceClient) {

	svc := &ServiceClientTopUp{
		Client: InitServiceClientTopUp(c),
	}

	routes := r.Group("/topup")
	routes.POST("/form", svc.Form)

}

func (ServiceClientTopUp *ServiceClientTopUp) Form(ctx *gin.Context) {
	routing := InitControllerTopUp()
	routing.Controller.Form(ctx, ServiceClientTopUp.Client)
}
