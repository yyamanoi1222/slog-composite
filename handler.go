package slogcomposite

import (
	"context"
	"log/slog"
)

type handler struct {
	handlers []slog.Handler
}

func newHandler(handlers []slog.Handler) *handler {
	return &handler{handlers: handlers}
}

func (h *handler) Enabled(ctx context.Context, level slog.Level) bool {
	// Always return true here.
	// determine Enabled for each Handler in the Handle method
	return true
}

func (h *handler) Handle(ctx context.Context, record slog.Record) error {
	for _, handler := range h.handlers {
		if !handler.Enabled(ctx, record.Level) {
			continue
		}

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
