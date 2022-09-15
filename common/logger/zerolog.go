package logger

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

type zeroLog struct {
	logger zerolog.Logger
}

func (z zeroLog) info(s string, tags ...tag.Tag) {
	tagsToEvents(z.logger.Info(), tags...).Time("time", time.Now()).Msg(s)
}

func (z zeroLog) infof(s string, args ...interface{}) {
	z.logger.Info().Time("time", time.Now()).Msg(fmt.Sprintf(s, args...))
}

func (z zeroLog) debug(s string, tags ...tag.Tag) {
	tagsToEvents(z.logger.Debug(), tags...).Time("time", time.Now()).Msg(s)
}

func (z zeroLog) debugf(s string, args ...interface{}) {
	z.logger.Debug().Time("time", time.Now()).Msg(fmt.Sprintf(s, args...))
}

func (z zeroLog) warning(s string, tags ...tag.Tag) {
	tagsToEvents(z.logger.Warn(), tags...).Time("time", time.Now()).Msg(s)
}

func (z zeroLog) warningf(s string, args ...interface{}) {
	z.logger.Warn().Time("time", time.Now()).Msg(fmt.Sprintf(s, args...))
}

func (z zeroLog) error(s string, err error, tags ...tag.Tag) {
	// to prevent panic
	if err == nil {
		z.info(s, tags...)
		return
	}

	tags = append(tags, tag.Err("err", err))
	tagsToEvents(z.logger.Error(), tags...).Time("time", time.Now()).Msg(s)
}

func (z zeroLog) errorf(s string, err error, args ...interface{}) {
	// to prevent panic
	if err == nil {
		z.infof(s, args...)
		return
	}

	z.logger.Error().Time("time", time.Now()).Msg(fmt.Sprintf(s, args...))
}

func (z zeroLog) fatal(s string, tags ...tag.Tag) {
	tagsToEvents(z.logger.Fatal(), tags...).Time("time", time.Now()).Msg(s)
}

func (z zeroLog) fatalf(s string, args ...interface{}) {
	z.logger.Fatal().Time("time", time.Now()).Msg(fmt.Sprintf(s, args...))
}

func tagsToEvents(event *zerolog.Event, tags ...tag.Tag) *zerolog.Event {
	for _, t := range tags {
		switch t.Type() {
		case tag.TypeString:
			event = event.Str(t.Key, t.Val.(string))
		case tag.TypeInt64:
			event = event.Int64(t.Key, t.Val.(int64))
		case tag.TypeInt32:
			event = event.Int32(t.Key, t.Val.(int32))
		case tag.TypeInt:
			event = event.Int(t.Key, t.Val.(int))
		case tag.TypeFloat64:
			event = event.Float64(t.Key, t.Val.(float64))
		case tag.TypeFloat32:
			event = event.Float32(t.Key, t.Val.(float32))
		case tag.TypeBoolean:
			event = event.Bool(t.Key, t.Val.(bool))
		case tag.TypeError:
			event = event.Err(t.Val.(error))
		case tag.TypeTime:
			event = event.Time(t.Key, t.Val.(time.Time))
		case tag.TypeDuration:
			event = event.Dur(t.Key, t.Val.(time.Duration))
		default:
			event = event.Interface(t.Key, t.Val)
		}
	}

	return event
}

func createZeroLog() zeroLog {
	logDir := os.Getenv("COURIER_MANAGEMENT_WORKING_DIR")
	if logDir == "" {
		logDir0, _ := user.Current()
		logDir = logDir0.HomeDir
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logFile, err := os.OpenFile(filepath.Join(logDir, config.Log().FilePath), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		zlog.Error().Err(err).Msg("there was an error during creating/opening the log file")
	}

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	multiLevelWriter := zerolog.MultiLevelWriter(consoleWriter, logFile)
	zeroLogger := zerolog.New(multiLevelWriter).With().Logger()

	return zeroLog{
		logger: zeroLogger,
	}
}
