package log

type LogConfig struct {
	LogPath       []string
	LogFileFormat string
	Prefix        string
	IsUseStdout   bool
	AutoIndentStr string
	Level         int
}
