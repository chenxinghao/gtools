package log

import (
	"fmt"
	"testing"
)

func TestLoggers_Save(t *testing.T) {
	LConfig := &LogConfig{}
	filepath := make([]string, 0)
	LConfig.IsUseStdout = true
	LConfig.Prefix = "[test]"
	LConfig.LogPath = append(filepath, "C:\\cxh\\work\\gtools\\")
	LConfig.LogFileFormat = "2006-01-02_15_04_05.log"
	LConfig.AutoIndentStr = ">>>>"
	LConfig.Level = 10
	LCP, err := Log.CreateLoggerContext(LConfig)
	if err != nil {
		fmt.Println(err)
	}
	Log.RegisterLogContext("test", LCP)

	Log.Save("test", 0, "testing", LConfig)
	Log.ClearAll()

}
