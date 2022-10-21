// Package ci_example defines the logger.
//
// package is using a global logger with some default parameters.
// It is disabled by default and the level can be increased using
// an environment variable:
//
//	MY_LOG=trace
//	MY_LOG=info
package ci_example

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// EnvLogLevel is the name of the environment variable to change the logging
// level.
const EnvLogLevel = "MY_LOG"

const defaultLogLevel = zerolog.WarnLevel

func init() {
	lvl := os.Getenv(EnvLogLevel)

	var level zerolog.Level

	switch lvl {
	case "error":
		level = zerolog.ErrorLevel
	case "warn":
		level = zerolog.WarnLevel
	case "info":
		level = zerolog.InfoLevel
	case "debug":
		level = zerolog.DebugLevel
	case "trace":
		level = zerolog.TraceLevel
	case "disabled":
		level = zerolog.Disabled
	default:
		level = defaultLogLevel
	}

	Logger = Logger.Level(level)
}

var logout = zerolog.ConsoleWriter{
	Out:        os.Stdout,
	TimeFormat: time.RFC3339,
}

// Logger is a globally available logger instance. By default, it only prints
// error level messages but it can be changed through a environment variable.
var Logger = zerolog.New(logout).Level(defaultLogLevel).
	With().Timestamp().Logger()
