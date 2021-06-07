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
	logEntry := logrus.NewEntry(l.logger)

	for i := 0; i < len(keyvals); i += 2 {
		key := keyvals[i].(string)
		if key == "ts" {
			continue
		}
		logEntry = logEntry.WithField(key, keyvals[i+1])
	}

	logEntry.Logf(logursLevel, "")
	return nil
}
