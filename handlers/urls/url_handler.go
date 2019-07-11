package handlers

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	dao "github.com/ildomm/linx_challenge/daos/urls"
	. "github.com/ildomm/linx_challenge/restapi/operations/urls"
	"net/http"
)

type Redirect struct {
	HTTPRequest *http.Request
	Url string
}

func (r Redirect) WriteResponse(hx http.ResponseWriter, rx runtime.Producer) {
	http.Redirect(hx, r.HTTPRequest, r.Url, 302)
}

func RedirectURLHandlerResponder(params RedirectURLParams) middleware.Responder {
	url := dao.ById(params.ID)
	if url == nil {
		return NewRedirectURLNotFound()
	}
	dao.IncreaseHits(url.ID)

	return &Redirect{params.HTTPRequest, url.URL}
}

func DeleteURLHandlerResponder(params DeleteURLParams) middleware.Responder {
	url := dao.ById(params.ID)
	if url == nil {
		return NewDeleteURLNotFound()
	}

	dao.Delete(params.ID)

	return nil
}
