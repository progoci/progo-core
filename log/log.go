package log

import (
	log "github.com/sirupsen/logrus"
)

// Logger represents a logger.
type Logger interface {
	WithFields(fields log.Fields) *log.Entry
	Printf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}
