package log

import "log"

type LoggerContext struct {
	LogConfig *LogConfig
	Logger    *log.Logger
}
