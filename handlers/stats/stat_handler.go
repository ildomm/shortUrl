package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ildomm/linx_challenge/daos/stats"
	"github.com/ildomm/linx_challenge/daos/users"
	. "github.com/ildomm/linx_challenge/restapi/operations/stats"
	urls "github.com/ildomm/linx_challenge/daos/urls"
)

func StatsGetSystemStatsHandlerResponder(params GetSystemStatsParams) middleware.Responder {
	tops, _ := stats.Top10SummaryHits(nil)

	return NewGetSystemStatsOK().WithPayload(tops)
}

func StatsGetURLStatsHandlerResponder(params GetURLStatsParams) middleware.Responder {
	url := urls.ById(params.ID)
	if url == nil {
		return NewGetURLStatsDefault(404)
	}

	return NewGetURLStatsOK().WithPayload(url)
}

func StatsGetUserStatsHandlerResponder(params GetUserStatsParams) middleware.Responder {
	userId := users.UserId(params.UserID)
	if userId == 0 {
		return NewGetUserStatsDefault(404)
	}

	tops, _ := stats.Top10SummaryHits(&userId)

	return NewGetUserStatsOK().WithPayload(tops)
}