package controller

import (
	"net/http"

	"github.com/M-fabricio-C/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/M-fabricio-C/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/M-fabricio-C/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/M-fabricio-C/meu-primeiro-crud-go/src/model"
	"github.com/M-fabricio-C/meu-primeiro-crud-go/src/view"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
        zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
		    zap.String("journey", "createUser"),)
		errRest := validation.ValidateUserError(err)
			
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	
	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}
	
	logger.Info("User created successfully",
	    zap.String("journey", "createUser"),)
	
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domain,
	))
	
}