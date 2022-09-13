package logger

import (
	"fmt"

	"go.uber.org/zap"
)

// NewLogger initializes a new logger interface.
func NewLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction(zap.Development())
	if err != nil {
		return nil, fmt.Errorf("zap initialize failed: %w", err)
	}

	return logger, nil
}
