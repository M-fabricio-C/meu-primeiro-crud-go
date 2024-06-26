package main

import (
	"log"

	"github.com/M-fabricio-C/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/M-fabricio-C/meu-primeiro-crud-go/src/controller"
	"github.com/M-fabricio-C/meu-primeiro-crud-go/src/controller/routes"
	"github.com/M-fabricio-C/meu-primeiro-crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	//Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}