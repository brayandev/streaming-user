package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	user "github.com/streaming-user/streaming-user"
	"go.uber.org/zap"
)

func createUserHandler(svc user.Service, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseWriter(w, http.StatusCreated, nil)
	}
}

// versionHandler returns application version.
func versionHandler(w http.ResponseWriter, r *http.Request) {
	contentType := "application/json"
	w.Header().Set("Content-Type", fmt.Sprintf("%s; charset=utf-8", contentType))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user.JSON()))
}

func responseWriter(w http.ResponseWriter, code int, content versionable) error {
	if content == nil {
		w.WriteHeader(code)
		return nil
	}

	jsonContent, err := json.Marshal(content)
	if err != nil {
		return err
	}

	contentType := "application/json"
	if content.Version() != "" {
		contentType = fmt.Sprintf("application/%s+json", content.Version())
	}

	w.Header().Set("Content-Type", fmt.Sprintf("%s; charset=utf-8", contentType))
	w.WriteHeader(code)
	w.Write(jsonContent)

	return nil
}

type versionable interface {
	Version() string
}
