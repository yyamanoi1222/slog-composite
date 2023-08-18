package main

import (
	"log/slog"
	"os"

	slogcomposite "github.com/yyamanoi1222/slog-composite"
)

func main() {
	lgr := slogcomposite.New(
		slog.NewJSONHandler(os.Stdout, nil),
		slog.NewTextHandler(os.Stdout, nil),
	)
	lgr.Info("Hello World!")

	lgrWGroup := lgr.WithGroup("group")
	lgrWGroup.Info("Hello World! with Group", slog.String("key", "value"))

	lgrWAttrs := lgr.With(slog.String("key", "value"))
	lgrWAttrs.Info("Hello World! with Attrs")

	lgrWithLevel := slogcomposite.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	lgrWithLevel.Debug("Hello World!")
}
