package logs

import (
	"context"
	"fmt"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sherlockhua/koala/config"
	"github.com/sirupsen/logrus"
)

var (
	loggerImp *LoggerImp
)

type Logger interface {
	Tracef(ctx context.Context, format string, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Accessf(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Panicf(ctx context.Context, format string, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})
	WithFields(fields logrus.Fields) Logger
}

type LoggerImp struct {
	logger       *logrus.Logger
	errorLogger  *logrus.Logger
	accessLogger *logrus.Logger
}

func NewLogger(conf *config.Config) (Logger, error) {

	loggerImp = &LoggerImp{
		logger:       logrus.New(),
		errorLogger:  logrus.New(),
		accessLogger: logrus.New(),
	}

	// Initialize the loggers
	err := Init(conf.Logger.AccessFileName, conf.Logger.LogLevel, loggerImp.accessLogger)
	if err != nil {
		return nil, err
	}

	err = Init(conf.Logger.Filename, conf.Logger.LogLevel, loggerImp.logger)
	if err != nil {
		return nil, err
	}

	err = Init(conf.Logger.ErrFileName, conf.Logger.LogLevel, loggerImp.errorLogger)
	if err != nil {
		return nil, err
	}

	return loggerImp, nil
}

func Init(filename string, logLevel string, logger *logrus.Logger) (err error) {
	/*file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		logrus.SetOutput(file)
	}*/

	logger.SetReportCaller(false)
	switch logLevel {
	case config.LogLevelDebug:
		logger.SetLevel(logrus.DebugLevel)
	case config.LogLevelInfo:
		logger.SetLevel(logrus.InfoLevel)
	case config.LogLevelWarn:
		logger.SetLevel(logrus.WarnLevel)
	case config.LogLevelError:
		logger.SetLevel(logrus.ErrorLevel)
	case config.LogLevelFatal:
		logger.SetLevel(logrus.FatalLevel)
	default:
		return fmt.Errorf("Unknown log level: %s", logLevel)

	}
	//logrus.SetFormatter(&MyFormatter{})

	writer, err := rotatelogs.New(
		//filename+".%Y%m%d%H%M",
		filename+".%Y%m%d%H",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Duration(7*24)*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(fmt.Sprintf("rotatelogs failed %v", err))
	}

	logger.AddHook(&MyHook{})
	logger.SetOutput(writer)

	return
}

// accessf logs a message at level Trace on the standard logger.
func Accessf(ctx context.Context, format string, args ...interface{}) {
	loggerImp.accessLogger.WithContext(ctx).Infof(format, args...)
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(ctx context.Context, format string, args ...interface{}) {
	loggerImp.logger.WithContext(ctx).Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(ctx context.Context, format string, args ...interface{}) {
	loggerImp.logger.WithContext(ctx).Debugf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(ctx context.Context, format string, args ...interface{}) {
	loggerImp.logger.WithContext(ctx).Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(ctx context.Context, format string, args ...interface{}) {
	loggerImp.errorLogger.WithContext(ctx).Warnf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(ctx context.Context, format string, args ...interface{}) {
	loggerImp.errorLogger.WithContext(ctx).Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(ctx context.Context, format string, args ...interface{}) {
	loggerImp.errorLogger.WithContext(ctx).Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	loggerImp.errorLogger.WithContext(ctx).Fatalf(format, args...)
}

func WithFields(fields logrus.Fields) Logger {
	return NewField(fields)
}

func (l *LoggerImp) Tracef(ctx context.Context, format string, args ...interface{}) {
	l.logger.WithContext(ctx).Tracef(format, args...)
}

func (l *LoggerImp) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.logger.WithContext(ctx).Debugf(format, args...)
}

func (l *LoggerImp) Infof(ctx context.Context, format string, args ...interface{}) {
	l.logger.WithContext(ctx).Infof(format, args...)
}

func (l *LoggerImp) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.errorLogger.WithContext(ctx).Warnf(format, args...)
}

func (l *LoggerImp) Panicf(ctx context.Context, format string, args ...interface{}) {
	l.errorLogger.WithContext(ctx).Panicf(format, args...)
}

func (l *LoggerImp) Fatalf(ctx context.Context, format string, args ...interface{}) {
	l.errorLogger.WithContext(ctx).Fatalf(format, args...)
}

func (l *LoggerImp) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.errorLogger.WithContext(ctx).Errorf(format, args...)
}

// accessf logs a message at level Trace on the standard logger.
func (l *LoggerImp) Accessf(ctx context.Context, format string, args ...interface{}) {
	l.accessLogger.WithContext(ctx).Infof(format, args...)
}

func (l *LoggerImp) WithFields(fields logrus.Fields) Logger {
	return NewField(fields)
}
