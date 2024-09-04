package logger_test

import (
	"testing"
	"vblog/logger"
)

func TestLogger(t *testing.T) {
	logger.Logger().Info().Msgf("debug")
}
