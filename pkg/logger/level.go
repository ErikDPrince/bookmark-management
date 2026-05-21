package logger

import (
	"os"

	"github.com/rs/zerolog"
)

const EnvLogLevel = "LOG_LEVEL"

// SetLogLevel set the global log level based on the environment variable LOG_LEVEL
func SetLogLevel() {
	level, err := zerolog.ParseLevel(os.Getenv(EnvLogLevel))
	if err != nil || level == zerolog.NoLevel {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

}
