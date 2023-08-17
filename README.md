# slog-composite

slog.Logger that can have multiple handlers

# Install

```
$ go get github.com/yyamanoi1222/slog-composite
```

# Usage

```go
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
}
```
