package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ildomm/linx_challenge/daos/urls"
	"github.com/ildomm/linx_challenge/daos/users"
	"github.com/ildomm/linx_challenge/models"
	. "github.com/ildomm/linx_challenge/restapi/operations/users"
)

func CreateUserHandlerResponder(params CreateUserParams) middleware.Responder {
	if users.IdentifyExists(params.ID.ID) {
		return NewCreateUserDefault(409)
	}

	user := new(models.User)
	user.ID = params.ID.ID

	users.Insert(user.ID)

	return NewCreateUserOK().WithPayload(user)
}

func DeleteUserHandlerResponder(params DeleteUserParams) middleware.Responder {
	userId := users.UserId(params.UserID)
	if userId == 0 {
		return NewDeleteUserNotFound()
	}

	urls.DeleteByUser(userId)
	users.Delete(userId)

	return nil
}

func CreateURLHandlerResponder(params CreateURLParams) middleware.Responder {
	if urls.UrlExists(params.URL.URL) {
		return NewCreateURLDefault(409)
	}

	userId := users.UserId(params.UserID)
	if userId == 0 {
		return NewCreateURLNotFound()
	}

	url := new(models.URL)
	url.URL = params.URL.URL
	url, _ = urls.Insert(*url, userId)

	return NewCreateURLOK().WithPayload(url)
}