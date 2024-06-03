package controller

import (
	"fmt"
	"log"

	"github.com/M-fabricio-C/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/M-fabricio-C/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
			log.Printf("Error trying to marshal object, error=%s\n", err.Error())
			restErr := rest_err.NewBadRequestError("Some fields are incorrect")
			
		c.JSON(restErr.Code, restErr)
		return
	}
	
	fmt.Println(userRequest)
}