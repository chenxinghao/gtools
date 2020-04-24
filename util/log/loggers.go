package log

import (
	FileUtils "github.com/chenxinghao/gtools/util/file"
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
	PermitLevel int
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

	Context.CornPtr = cron.New()
	spec := "0 0 0 1/1 * ?"
	Context.CornPtr.AddFunc(spec, func() {
		l.refreshLoggerContext(name)
	})
	Context.CornPtr.Start()

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
	info := &FileUtils.Info{}
	for _, logPath := range config.LogPath {
		logPath += info.GetSystemFilePathDelimiter() + time.Now().Format(config.LogFileFormat)
		f, err = os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		writers = append(writers, f)
		context.FilePtr = f
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
	if context, ok := l.LogContext.m[name]; ok {
		logConfig := *context.LogConfig
		LoggerContext, _ := l.CreateLoggerContext(&logConfig)
		l.RegisterLogContext(name, LoggerContext)
		context.FilePtr.Close()
		context.CornPtr.Stop()
	}

}

func (l *Loggers) ClearAll() {
	l.LogContext.Lock()
	defer l.LogContext.Unlock()
	for _, v := range l.LogContext.m {
		v.CornPtr.Stop()
		v.FilePtr.Close()
	}
}
