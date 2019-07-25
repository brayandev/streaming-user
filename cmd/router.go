package main

import (
	"github.com/go-chi/chi"
	user "github.com/streaming-user/strm-user"
	"go.uber.org/zap"
)

func createRouter(service user.Service, logger *zap.Logger) chi.Router {
	router := chi.NewRouter()

	router.Use(contextMiddleware)
	router.With(accessLogMiddleware(logger)).Route("/user", func(router chi.Router) {
		router.Post("/create", createUserHandler(service, logger))
	})

	router.Get("/version", versionHandler)
	return router
}
