package logs

import (
	"context"

	"github.com/sirupsen/logrus"
)

type field struct {
	fields map[string]interface{}
}

type Fields map[string]interface{}

func NewField(fields Fields) *field {
	return &field{fields: fields}
}

func (f *field) Tracef(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.logger.WithContext(ctx).WithFields(logrus.Fields(f.fields)).Tracef(format, args...)
		return
	}
	loggerImp.logger.WithContext(ctx).Tracef(format, args...)
}

func (f *field) Debugf(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.logger.WithContext(ctx).WithFields(f.fields).Debugf(format, args...)
		return
	}
	loggerImp.logger.WithContext(ctx).Debugf(format, args...)
}

func (f *field) Infof(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.logger.WithContext(ctx).WithFields(f.fields).Infof(format, args...)
		return
	}
	loggerImp.logger.WithContext(ctx).Infof(format, args...)
}

func (f *field) Warnf(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.errorLogger.WithContext(ctx).WithFields(f.fields).Warnf(format, args...)
		return
	}
	loggerImp.errorLogger.WithContext(ctx).Warnf(format, args...)
}

func (f *field) Panicf(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.errorLogger.WithContext(ctx).WithFields(f.fields).Panicf(format, args...)
		return
	}
	loggerImp.errorLogger.WithContext(ctx).Panicf(format, args...)
}

func (f *field) Fatalf(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.errorLogger.WithContext(ctx).WithFields(f.fields).Fatalf(format, args...)
		return
	}
	loggerImp.errorLogger.WithContext(ctx).Fatalf(format, args...)
}

func (f *field) Errorf(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.errorLogger.WithContext(ctx).WithFields(f.fields).Errorf(format, args...)
		return
	}
	loggerImp.errorLogger.WithContext(ctx).Errorf(format, args...)
}

func (f *field) Accessf(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.accessLogger.WithContext(ctx).WithFields(f.fields).Infof(format, args...)
		return
	}
	loggerImp.accessLogger.WithContext(ctx).Infof(format, args...)
}

func (f *field) WithFields(fields Fields) Logger {
	for k, v := range fields {
		f.fields[k] = v
	}
	return f
}
