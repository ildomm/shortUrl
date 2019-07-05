package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	."github.com/ildomm/linx_challenge/restapi/operations"
)

func OptionsHandlerResponder(params OptionsAllowParams) middleware.Responder {
	return NewOptionsAllowOK()
}
