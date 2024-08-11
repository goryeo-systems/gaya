package util

import (
	"fmt"
	"log/slog"
	"math/big"
	"os"
)

// Check panics if err is not nil.
func Check(err error) {
	if err != nil {
		LogError(err)
		panic(err)
	}
}

func getLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

var Log = getLogger() //nolint:gochecknoglobals

// LogError logs an error if it is not nil.
func LogError(err error) {
	if err != nil {
		Log.Error("error", "error", err)
	}
}

// StringToBigFloat converts a string to a big.Float.
func StringToBigFloat(s string) (*big.Float, error) {
	if v, ok := new(big.Float).SetString(s); ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to convert string to big.Float: %s", s)
}
