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
	return nil
}

func CreateURLHandlerResponder(params CreateURLParams) middleware.Responder {
	urlElement := new(models.URL)
	if models.UrlExists(params.UserID, params.URL.URL) {
		// Load details
	} else {
		urlElement.URL = params.URL.URL
		urlElement.GenerateShort()
		urlElement.SaveUrlKey(params.UserID)
		urlElement.SaveUrlDetail(params.UserID)
	}

	urlElement.URL = params.URL.URL
	urlElement.ID = models.UrlDBId(params.UserID, params.URL.URL)
	urlElement.LoadUrlDetail(params.UserID)

	return NewCreateURLOK().WithPayload(urlElement)
}