package main

import (
	"fmt"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	user "github.com/streaming-user/streaming-user"
	"go.uber.org/zap"
)

func main() {
	cfg, cErr := newConfig()
	if cErr != nil {
		panic(cErr)
	}

	logger, lErr := user.ConfigLog(zap.NewAtomicLevelAt(cfg.LogLevel.Value)).Build()
	if lErr != nil {
		panic(lErr)
	}
	defer logger.Sync()

	db, dbErr := user.NewMySQL(cfg.DBDriver, cfg.DBSource, cfg.DBTimeout, logger)
	if dbErr != nil {
		logger.Error("failed to create a new connection on db", zap.Error(dbErr))
	}

	repository := user.NewRepository(db, cfg.DBTimeout)

	service := user.NewService(repository)

	router := createRouter(service, logger)

	sErr := gracehttp.Serve(&http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})
	if sErr != nil {
		logger.Error("failed on server start", zap.NamedError("error", sErr))
	}

}
