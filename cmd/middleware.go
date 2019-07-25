package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	user "github.com/streaming-user/strm-user"
	"go.uber.org/zap"
)

type loggableResponseWriter struct {
	writer     http.ResponseWriter
	size       int
	statusCode int
}

func (lrw *loggableResponseWriter) Header() http.Header {
	return lrw.writer.Header()
}

func (lrw *loggableResponseWriter) Write(b []byte) (int, error) {
	lrw.size += len(b)
	return lrw.writer.Write(b)
}

func (lrw *loggableResponseWriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.writer.WriteHeader(statusCode)
}

func accessLogMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			lrw := &loggableResponseWriter{writer: w, statusCode: http.StatusOK}
			next.ServeHTTP(lrw, r)

			logger.Info("server access response",
				zap.String("transaction-id", user.GetTransactionID(r.Context())),
				zap.String("internal-id", user.GetInternalID(r.Context())),
				zap.Int("code", lrw.statusCode),
				zap.Duration("duration", time.Since(start)),
				zap.String("uri", r.RequestURI),
				zap.String("method", r.Method),
				zap.Int("size", lrw.size),
			)
		}
		return http.HandlerFunc(fn)
	}
}

func contextMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		transactionID := r.Header.Get("X-Transaction-Id")
		if transactionID == "" {
			transactionID = generateID()
		}
		internalID := generateID()
		w.Header().Add("X-Internal-Id", internalID)
		w.Header().Add("X-Transaction-Id", transactionID)

		next.ServeHTTP(w, r.WithContext(context.WithValue(context.WithValue(r.Context(), user.ContextKeyTransactionID, transactionID), user.ContextKeyInternalID, internalID)))
	}
	return http.HandlerFunc(fn)
}

func generateID() string {
	id, err := uuid.NewV4()
	if err != nil {
		id = uuid.FromStringOrNil(time.Now().Format(time.RFC3339Nano))
	}
	return "strm-user-" + id.String()
}
