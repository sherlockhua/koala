package logs

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Field struct {
	fields logrus.Fields
}

func NewField(fields logrus.Fields) *Field {
	return &Field{fields: fields}
}

func (f *Field) Tracef(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.logger.WithContext(ctx).WithFields(f.fields).Tracef(format, args...)
		return
	}
	loggerImp.logger.WithContext(ctx).Tracef(format, args...)
}

func (f *Field) Debugf(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.logger.WithContext(ctx).WithFields(f.fields).Debugf(format, args...)
		return
	}
	loggerImp.logger.WithContext(ctx).Debugf(format, args...)
}

func (f *Field) Infof(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.logger.WithContext(ctx).WithFields(f.fields).Infof(format, args...)
		return
	}
	loggerImp.logger.WithContext(ctx).Infof(format, args...)
}

func (f *Field) Warnf(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.errorLogger.WithContext(ctx).WithFields(f.fields).Warnf(format, args...)
		return
	}
	loggerImp.errorLogger.WithContext(ctx).Warnf(format, args...)
}

func (f *Field) Panicf(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.errorLogger.WithContext(ctx).WithFields(f.fields).Panicf(format, args...)
		return
	}
	loggerImp.errorLogger.WithContext(ctx).Panicf(format, args...)
}

func (f *Field) Fatalf(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.errorLogger.WithContext(ctx).WithFields(f.fields).Fatalf(format, args...)
		return
	}
	loggerImp.errorLogger.WithContext(ctx).Fatalf(format, args...)
}

func (f *Field) Errorf(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.errorLogger.WithContext(ctx).WithFields(f.fields).Errorf(format, args...)
		return
	}
	loggerImp.errorLogger.WithContext(ctx).Errorf(format, args...)
}

func (f *Field) Accessf(ctx context.Context, format string, args ...interface{}) {
	if f.fields != nil {
		loggerImp.accessLogger.WithContext(ctx).WithFields(f.fields).Infof(format, args...)
		return
	}
	loggerImp.accessLogger.WithContext(ctx).Infof(format, args...)
}

func (f *Field) WithFields(fields logrus.Fields) Logger {
	for k, v := range fields {
		f.fields[k] = v
	}
	return f
}
