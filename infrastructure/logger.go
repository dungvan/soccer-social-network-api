package infrastructure

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	// OutputStdout check config(logger.output).
	OutputStdout = "stdout"
	// OutputFile check config(logger.output).
	OutputFile = "file"

	// FormatText check config(logger.format).
	FormatText = "text"
	// FormatJSON check config(logger.format).
	FormatJSON = "json"
)

// Logger struct.
type Logger struct {
	Log     *logrus.Logger
	Logfile *os.File
}

var logType string

// NewLoggerWithType returns new Logger for logType.
func NewLoggerWithType(t string) *Logger {
	logType = t
	return NewLogger()
}

// NewLogger returns new Logger.
// repository: https://github.com/sirupsen/logrus
func NewLogger() *Logger {
	var err error
	var file *os.File

	if logType == "" {
		logType = "app"
	}

	// get config.
	output := GetConfigString(logType + ".logger.output")
	level := GetConfigString(logType + ".logger.level")
	format := GetConfigString(logType + ".logger.format")

	// new logrus.
	log := logrus.New()

	// set output.
	switch output {
	case OutputStdout: // output: stdout
		log.Out = os.Stdout
	case OutputFile: // output: file
		logfile := GetConfigString(logType + ".logger.file")
		file, err = os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		log.Out = file
	}

	// set level.
	log.Level, err = logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}

	// set formatter.
	switch format {
	case FormatText:
		log.Formatter = &logrus.TextFormatter{}
	case FormatJSON:
		log.Formatter = &logrus.JSONFormatter{}
	}
	return &Logger{Log: log, Logfile: file}
}

// CloseLogger close logfile
func CloseLogger(logfile *os.File) {
	// close file.
	if logfile != nil {
		err := logfile.Close()
		if err != nil {
			panic(err)
		}
	}
}
