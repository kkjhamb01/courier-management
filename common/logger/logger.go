package logger

import "gitlab.artin.ai/backend/courier-management/common/logger/tag"

type logger interface {
	info(s string, tags ...tag.Tag)
	infof(s string, args ...interface{})
	debug(s string, tags ...tag.Tag)
	debugf(s string, args ...interface{})
	warning(s string, tags ...tag.Tag)
	warningf(s string, args ...interface{})
	error(s string, err error, tags ...tag.Tag)
	errorf(s string, err error, args ...interface{})
	fatal(s string, tags ...tag.Tag)
	fatalf(s string, args ...interface{})
}

var loggerImpl logger

func Info(s string, tags ...tag.Tag) {
	loggerImpl.info(s, tags...)
}

func Infof(s string, args ...interface{}) {
	loggerImpl.infof(s, args...)
}

func Debug(s string, tags ...tag.Tag) {
	loggerImpl.debug(s, tags...)
}

func Debugf(s string, args ...interface{}) {
	loggerImpl.debugf(s, args...)
}

func Warning(s string, tags ...tag.Tag) {
	loggerImpl.warning(s, tags...)
}

func Warningf(s string, args ...interface{}) {
	loggerImpl.warningf(s, args...)
}

func Error(s string, err error, tags ...tag.Tag) {
	loggerImpl.error(s, err, tags...)
}

func Errorf(s string, err error, args ...interface{}) {
	loggerImpl.errorf(s, err, args...)
}

func Fatal(s string, tags ...tag.Tag) {
	loggerImpl.fatal(s, tags...)
}

func Fatalf(s string, args ...interface{}) {
	loggerImpl.fatalf(s, args...)
}

func InitLogger() {
	loggerImpl = createZeroLog()
}
