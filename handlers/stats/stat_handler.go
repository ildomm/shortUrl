package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ildomm/linx_challenge/restapi/operations/stats"
)

func StatsGetSystemStatsHandlerResponder(params stats.GetSystemStatsParams) middleware.Responder {
	return nil
}

func StatsGetURLStatsHandlerResponder(params stats.GetURLStatsParams) middleware.Responder {
	return nil
}

func StatsGetUserStatsHandlerResponder(params stats.GetUserStatsParams) middleware.Responder {
	return nil
}