package main

import (
	"fmt"
	"log"

	routes "github.com/ahfrd/gateway-apps-grpc/src/routes"

	"github.com/ahfrd/gateway-apps-grpc/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		fmt.Println("asdas")
		log.Fatal("Failed at config", err)
	}

	router := gin.Default()
	authServiceClient := *routes.RoutesAuth(router, &c)
	routes.RoutesEmoney(router, &c, &authServiceClient)
	routes.RoutesTopUp(router, &c, &authServiceClient)
	fmt.Println(authServiceClient)
	router.Run(c.Port)

}
