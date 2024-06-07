package service

import (
	"github.com/M-fabricio-C/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/M-fabricio-C/meu-primeiro-crud-go/src/model"
)

func (*userDomainService) UpdateUser(
	userId string, 
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}