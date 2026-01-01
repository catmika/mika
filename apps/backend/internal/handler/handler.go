package handler

import "log/slog"

type Handler struct {
	Logger *slog.Logger
}

func New(logger *slog.Logger) *Handler {
	return &Handler{
		Logger: logger,
	}
}
