package util

import (
	"fmt"
	"log/slog"
	"math/big"
	"os"
)

// Check panics if err is not nil
func Check(err error) {
	if err != nil {
		LogError(err)
		panic(err)
	}
}

// GetLogger returns a new JSON lines logger
func GetLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

var log = GetLogger()

// LogError logs an error if it is not nil
func LogError(err error) {
	if err != nil {
		log.Error("error", "error", err)
	}
}

// StringToBigFloat converts a string to a big.Float
func StringToBigFloat(s string) (*big.Float, error) {
	if v, ok := new(big.Float).SetString(s); ok {
		return v, nil
	} else {
		return nil, fmt.Errorf("failed to convert string to big.Float: %s", s)
	}
}
