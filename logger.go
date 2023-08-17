package slogcomposite

import (
	"log/slog"
)

func New(handlers ...slog.Handler) *slog.Logger {
	if len(handlers) == 0 {
		return slog.New(nil)
	}

	return slog.New(newHandler(handlers))
}
