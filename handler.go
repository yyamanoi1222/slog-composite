package slogcomposite

import (
	"log/slog"
	"context"
)

type handler struct {
	handlers []slog.Handler
}

func newHandler(handlers []slog.Handler) *handler {
	return &handler{handlers: handlers}
}

func (h *handler) Enabled(ctx context.Context, level slog.Level) bool {
	enabled := true
	for _, handler := range h.handlers {
		enabled = enabled && handler.Enabled(ctx, level)
	}
	return enabled
}

func (h *handler) Handle(ctx context.Context, record slog.Record) error {
	for _, handler := range h.handlers {
		if err := handler.Handle(ctx, record); err != nil {
			return err
		}
	}
	return nil
}

func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandler := newHandler(append([]slog.Handler{}, h.handlers...))
	for i, handler := range newHandler.handlers {
		newHandler.handlers[i] = handler.WithAttrs(attrs)
	}
	return newHandler
}

func (h *handler) WithGroup(name string) slog.Handler {
	newHandler := newHandler(append([]slog.Handler{}, h.handlers...))
	for i, handler := range h.handlers {
		newHandler.handlers[i] = handler.WithGroup(name)
	}
	return newHandler
}

