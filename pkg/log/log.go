package log

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/sirupsen/logrus"
)

type LogursLogger struct {
	logger *logrus.Logger
}

func NewLogger() *LogursLogger {

	lg := logrus.New()
	lg.SetFormatter(&logrus.JSONFormatter{})

	return &LogursLogger{lg}
}
func (l *LogursLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 {
		return nil
	}
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "")
	}

	logursLevel, _ := logrus.ParseLevel(level.String())

	var logEntry *logrus.Entry

	for i := 0; i < len(keyvals); i += 2 {
		if logEntry == nil {
			logEntry = l.logger.WithField(keyvals[i].(string), keyvals[i+1])
		} else {
			logEntry = logEntry.WithField(keyvals[i].(string), keyvals[i+1])
		}
	}

	logEntry.Logf(logursLevel, "")
	return nil
}
