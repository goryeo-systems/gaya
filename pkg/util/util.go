package util

import (
	"log/slog"
	"os"
)

// Check panics if err is not nil
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// GetLogger returns a new JSON lines logger
func GetLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
