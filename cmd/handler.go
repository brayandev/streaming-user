package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	user "github.com/streaming-user/streaming-user"
	"go.uber.org/zap"
)

func createUserHandler(svc user.Service, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		router := "create-user"
		usr := &user.User{}

		pErr := parseJSON(r.Body, &usr)
		if pErr != nil {
			writeError(w, pErr)
			user.LogError(ctx, logger, router, "cannot parse content", pErr)
		}
		fmt.Println(usr)

		result, err := svc.InsertUser(ctx, usr)
		if err != nil {
			writeError(w, err)
			user.LogError(ctx, logger, router, "cannot save user on database", err, zap.Any("result", result))
			return
		}
		lastID, lErr := result.LastInsertId()
		if lErr != nil {
			user.LogError(ctx, logger, router, "error on get last insert id", lErr, zap.Int64("last-id", lastID))
			return
		}
		responseWriter(w, http.StatusCreated, nil, lastID)
	}
}

// versionHandler returns application version.
func versionHandler(w http.ResponseWriter, r *http.Request) {
	contentType := "application/json"
	w.Header().Set("Content-Type", fmt.Sprintf("%s; charset=utf-8", contentType))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user.JSON()))
}

func responseWriter(w http.ResponseWriter, code int, content versionable, response interface{}) error {
	if content == nil {
		w.WriteHeader(code)
		return nil
	}

	jsonContent, err := json.Marshal(content)
	if err != nil {
		return err
	}

	resp, err := json.Marshal(response)
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
	w.Write(resp)

	return nil
}

func writeError(w http.ResponseWriter, err error) {
	switch tErr := err.(type) {
	case *user.Error:
		responseWriter(w, getErrorHTTPCode(tErr), tErr, nil)
	default:
		responseWriter(w, http.StatusInternalServerError, user.NewUnknownError(err.Error()), nil)
	}
}

func getErrorHTTPCode(err *user.Error) int {
	switch err.ErrType {
	case user.ErrorInvalidContent:
		return http.StatusBadRequest

	default:
		return http.StatusInternalServerError
	}
}

func parseJSON(reader io.ReadCloser, out interface{}) error {
	err := json.NewDecoder(reader).Decode(out)
	if err != nil {
		return user.NewInvalidContentError(fmt.Sprintf("could not parse body content, error: %s", err.Error()))
	}
	return nil
}

type versionable interface {
	Version() string
}
