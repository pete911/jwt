package cmd

import (
	"io"
	"log/slog"
	"os"
)

func initLog(silent bool) {

	if silent {
		slog.SetDefault(newLogger(io.Discard))
		return
	}
	slog.SetDefault(newLogger(os.Stderr))
}

func newLogger(out io.Writer) *slog.Logger {
	return slog.New(slog.NewTextHandler(out, nil))
}
