package log

import (
	"github.com/robfig/cron"
	"log"
	"os"
)

type LoggerContext struct {
	LogConfig *LogConfig
	Logger    *log.Logger
	FilePtr   *os.File
	CornPtr   *cron.Cron
}
