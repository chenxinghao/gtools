package log

type LogConfig struct {
	LogFilePaths  []string
	Prefix        string
	IsUseStdout   bool
	AutoIndentStr string
	Level         int
}
