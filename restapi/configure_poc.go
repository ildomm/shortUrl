// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/ildomm/linx_challenge/restapi/operations"
	"github.com/ildomm/linx_challenge/restapi/operations/stats"
	"github.com/ildomm/linx_challenge/restapi/operations/urls"
	"github.com/ildomm/linx_challenge/restapi/operations/users"
)

//go:generate swagger generate server --target ..\..\linx_challenge --name Poc --spec ..\spec\swagger.json

func configureFlags(api *operations.PocAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PocAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.UsersCreateURLHandler == nil {
		api.UsersCreateURLHandler = users.CreateURLHandlerFunc(func(params users.CreateURLParams) middleware.Responder {
			return middleware.NotImplemented("operation users.CreateURL has not yet been implemented")
		})
	}
	if api.UsersCreateUserHandler == nil {
		api.UsersCreateUserHandler = users.CreateUserHandlerFunc(func(params users.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation users.CreateUser has not yet been implemented")
		})
	}
	if api.UrlsDeleteURLHandler == nil {
		api.UrlsDeleteURLHandler = urls.DeleteURLHandlerFunc(func(params urls.DeleteURLParams) middleware.Responder {
			return middleware.NotImplemented("operation urls.DeleteURL has not yet been implemented")
		})
	}
	if api.UsersDeleteUserHandler == nil {
		api.UsersDeleteUserHandler = users.DeleteUserHandlerFunc(func(params users.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation users.DeleteUser has not yet been implemented")
		})
	}
	if api.StatsGetSystemStatsHandler == nil {
		api.StatsGetSystemStatsHandler = stats.GetSystemStatsHandlerFunc(func(params stats.GetSystemStatsParams) middleware.Responder {
			return middleware.NotImplemented("operation stats.GetSystemStats has not yet been implemented")
		})
	}
	if api.StatsGetURLStatsHandler == nil {
		api.StatsGetURLStatsHandler = stats.GetURLStatsHandlerFunc(func(params stats.GetURLStatsParams) middleware.Responder {
			return middleware.NotImplemented("operation stats.GetURLStats has not yet been implemented")
		})
	}
	if api.StatsGetUserStatsHandler == nil {
		api.StatsGetUserStatsHandler = stats.GetUserStatsHandlerFunc(func(params stats.GetUserStatsParams) middleware.Responder {
			return middleware.NotImplemented("operation stats.GetUserStats has not yet been implemented")
		})
	}
	if api.OptionsAllowHandler == nil {
		api.OptionsAllowHandler = operations.OptionsAllowHandlerFunc(func(params operations.OptionsAllowParams) middleware.Responder {
			return middleware.NotImplemented("operation .OptionsAllow has not yet been implemented")
		})
	}
	if api.UrlsRedirectURLHandler == nil {
		api.UrlsRedirectURLHandler = urls.RedirectURLHandlerFunc(func(params urls.RedirectURLParams) middleware.Responder {
			return middleware.NotImplemented("operation urls.RedirectURL has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
