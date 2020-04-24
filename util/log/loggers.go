package log

import (
	"github.com/robfig/cron"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var Log *Loggers

type LoggerContextMap struct {
	sync.RWMutex
	m map[string]*LoggerContext
}

type Loggers struct {
	LogContext  LoggerContextMap
	CLoseFiles  []*os.File
	PermitLevel int
	CornPtr     *cron.Cron
}

func init() {
	Log = &Loggers{}
	Log.LogContext = LoggerContextMap{m: make(map[string]*LoggerContext, 0)}
}

func (l *Loggers) RegisterLogContext(name string, Context *LoggerContext) {
	l.LogContext.Lock()
	defer l.LogContext.Unlock()
	if name != "" {
		l.LogContext.m[name] = Context
	} else {
		l.LogContext.m["Unkwown"] = Context
	}

	l.CornPtr = cron.New()
	spec := "0 0 0 1/1 * ?"
	l.CornPtr.AddFunc(spec, func() {
		l.refreshLoggerContext(name)
	})
	l.CornPtr.Start()

}
func (l *Loggers) getLogContextByName(name string) *LoggerContext {
	l.LogContext.RLock()
	defer l.LogContext.RUnlock()
	if context, ok := l.LogContext.m[name]; ok {
		return context
	} else {
		return nil
	}
}

func (l *Loggers) CreateLoggerContext(config *LogConfig) (*LoggerContext, error) {
	context := &LoggerContext{}
	context.LogConfig = config
	writers := make([]io.Writer, 0)
	var err error
	var f *os.File
	for _, logPath := range config.LogPath {
		logPath += time.Now().Format(config.LogFileFormat)
		f, err = os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		writers = append(writers, f)
		l.CLoseFiles = append(l.CLoseFiles, f)
	}
	if config.IsUseStdout {
		writers = append(writers, os.Stdout)
	}
	multiWriter := io.MultiWriter(writers...)
	context.Logger = log.New(multiWriter, config.Prefix, 0)
	return context, err
}

func (l *Loggers) Save(name string, indent int, contents ...interface{}) {
	LoggerP := l.getLogContextByName(name)
	if LoggerP != nil {
		for _, content := range contents {
			LoggerP.Logger.Println(content)
		}
	} else {
		LoggerP := l.getLogContextByName("Error")
		LoggerP.Logger.Println("LoggerContext is not exist,name: " + name)
	}
}

func (l *Loggers) refreshLoggerContext(name string) {
	l.LogContext.Lock()
	defer l.LogContext.Unlock()
	if context, ok := l.LogContext.m[name]; ok {
		logConfig := *context.LogConfig
		LoggerContext, _ := l.CreateLoggerContext(&logConfig)
		l.LogContext.m[name] = LoggerContext
	}

}

func (l *Loggers) ClearAll() {
	l.CornPtr.Stop()
}
