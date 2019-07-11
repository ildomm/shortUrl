package utils

import (
	"encoding/json"
	"github.com/go-openapi/runtime/middleware"
	"github.com/ildomm/linx_challenge/models"
	"github.com/ildomm/linx_challenge/restapi/operations/users"
	"net/http"
	"syreclabs.com/go/faker"
)

func GenerateRandomHttpRequest() *http.Request {
	hTTPRequest := new(http.Request)
	hTTPRequest.RemoteAddr = faker.Internet().IpV4Address()

	return hTTPRequest
}

func GenerateRandomUserParams(_id *string) *users.CreateUserParams {
	params := new(users.CreateUserParams)
	params.HTTPRequest = GenerateRandomHttpRequest()

	id := new(models.User)
	if _id == nil {
		id.ID = faker.Bitcoin().Address()
	} else {
		id.ID = *_id
	}

	params.ID = id

	return params
}

func GenerateRandomUrlParams(userId string) *users.CreateURLParams {
	params := new(users.CreateURLParams)
	params.HTTPRequest = GenerateRandomHttpRequest()

	url := new(users.CreateURLBody)
	url.URL = faker.Bitcoin().Address()

	params.URL = *url
	params.UserID = userId

	return params
}

func GenerateRandomUserDeleteParams(_id *string) *users.DeleteUserParams {
	params := new(users.DeleteUserParams)
	params.HTTPRequest = GenerateRandomHttpRequest()

	if _id == nil {
		params.UserID = faker.Bitcoin().Address()
	} else {
		params.UserID = *_id
	}

	return params
}

func GetUsersCreationPayload(middlewareResult middleware.Responder) *models.User {
	out, err := json.Marshal(middlewareResult)
	if err != nil {
		panic(err)
	}
	payload := new(users.CreateUserOK)
	if err := json.Unmarshal(out, &payload); err != nil {
		panic(err)
	}

	return payload.Payload
}

func GetUrlsCreationPayload(middlewareResult middleware.Responder) *models.URL {
	out, err := json.Marshal(middlewareResult)
	if err != nil {
		panic(err)
	}
	payload := new(users.CreateURLOK)
	if err := json.Unmarshal(out, &payload); err != nil {
		panic(err)
	}

	return payload.Payload
}
