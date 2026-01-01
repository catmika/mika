package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/catmika/when-light/internal/config"
	"github.com/catmika/when-light/internal/handler"
	"github.com/catmika/when-light/internal/middleware"
)

func main() {
	config := config.Load()

	logger := config.Logger
	port := config.Port

	h := handler.New(logger)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      middleware.Logging(logger)(h.Routes()),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("starting server", slog.String("port", port))

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
