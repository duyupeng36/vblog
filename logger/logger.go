package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"sync"
)

var (
	once   sync.Once
	logger zerolog.Logger
)

func Logger() *zerolog.Logger {
	return &logger
}

func init() {

	once.Do(func() {
		writer := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05 -0700"}

		// 级别格式
		writer.FormatLevel = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		}

		logger = zerolog.New(writer).With().Timestamp().Caller().Logger().Level(zerolog.InfoLevel)
	})
}
