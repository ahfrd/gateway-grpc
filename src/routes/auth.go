package routes

import (
	"fmt"

	"github.com/ahfrd/gateway-apps-grpc/config"
	controller "github.com/ahfrd/gateway-apps-grpc/src/controller/auth"
	proto "github.com/ahfrd/gateway-apps-grpc/src/proto/auth"
	repository "github.com/ahfrd/gateway-apps-grpc/src/repository"
	svc "github.com/ahfrd/gateway-apps-grpc/src/service/auth"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client proto.AuthServiceClient
}
type ControllerAuth struct {
	Controller controller.AuthController
}

func InitServiceClient(c *config.Config) proto.AuthServiceClient {
	fmt.Println(c.AuthSvcUrl)
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("Could not connect : ", err)
		fmt.Println(err)
	}
	return proto.NewAuthServiceClient(cc)
}

func RoutesAuth(r *gin.Engine, c *config.Config) *ServiceClient {
	serviceClient := &ServiceClient{
		Client: InitServiceClient(c),
	}
	routes := r.Group("/auth")
	routes.POST("/register", serviceClient.Register)
	routes.POST("/login", serviceClient.Login)
	return serviceClient
}

func InitControllerAuth() *ControllerAuth {
	authRepository := repository.NewAuthRepository()
	authSvc := svc.NewAuthService(&authRepository)
	authController := controller.NewAuthController(&authSvc)
	controller := &ControllerAuth{
		Controller: authController,
	}
	return controller
}

func (ServiceClient *ServiceClient) Register(ctx *gin.Context) {
	routing := InitControllerAuth()
	routing.Controller.Register(ctx, ServiceClient.Client)
}

func (ServiceClient *ServiceClient) Login(ctx *gin.Context) {
	routing := InitControllerAuth()
	routing.Controller.Login(ctx, ServiceClient.Client)
}
