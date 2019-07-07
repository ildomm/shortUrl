package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ildomm/linx_challenge/models"
	. "github.com/ildomm/linx_challenge/restapi/operations/users"
)

func CreateUserHandlerResponder(params CreateUserParams) middleware.Responder {
	if models.UserDBId(params.ID.ID) {
		return NewCreateUserDefault(409)
	}

	user := new(models.User)
	user.ID = params.ID.ID
	user.Save()

	return NewCreateUserOK().WithPayload(user)
}

func DeleteUserHandlerResponder(params DeleteUserParams) middleware.Responder {
	return nil
}

func CreateURLHandlerResponder(params CreateURLParams) middleware.Responder {
	return nil
}