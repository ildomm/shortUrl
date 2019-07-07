package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ildomm/linx_challenge/models"
	. "github.com/ildomm/linx_challenge/restapi/operations/users"
)

func CreateUserHandlerResponder(params CreateUserParams) middleware.Responder {
	if models.UserExists(params.ID.ID) {
		return NewCreateUserDefault(409)
	}

	user := new(models.User)
	user.ID = params.ID.ID
	user.SaveUser()

	return NewCreateUserOK().WithPayload(user)
}

func DeleteUserHandlerResponder(params DeleteUserParams) middleware.Responder {

	// Delete urls from user
	// Delete user

	return nil
}

func CreateURLHandlerResponder(params CreateURLParams) middleware.Responder {

	return NewCreateURLOK()
}