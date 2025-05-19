package logs

import (
	"context"
	"testing"

	"github.com/sherlockhua/koala/config"
)

func TestNewLogger_TraceMessageWithContext(t *testing.T) {
	// Create a mock config
	mockConfig := &config.Config{
		Logger: config.LoggerConfig{
			AccessFileName: "./logs/access.log",
			Filename:       "./logs/app.log",
			ErrFileName:    "./logs/error.log",
			LogLevel:       "debug",
		},
	}

	// Create a new logger
	logger := NewLogger(mockConfig)
	// Create a buffer to capture log output
	//var buf bytes.Buffer

	// Replace the output of the logger with our buffer
	//logger.(*LoggerImp).logger.SetOutput(&buf)

	// Create a context
	ctx := context.Background()
	ctx = WithLogId(ctx, "thisislslslslsslsllsl")
	// Log a trace message
	for i := 0; i < 1000; i++ {
		WithFields(Fields{"user_id": 22, "password": 222, "name": "zhanghua"}).Accessf(ctx, "this is a access log message")
		logger.Tracef(ctx, "Test message with %s", "formatting")
		logger.Debugf(ctx, "Test message with %s", "formatting")
		logger.Errorf(ctx, "Test message with %s", "formatting")
		logger.WithFields(Fields{"user_id": 22, "password": 222, "name": "zhanghua"}).Errorf(ctx, "Test message with %s", "format")
	}
	// Check if the log message contains the expected content
	//if !strings.Contains(buf.String(), "Test message with formatting") {
	//	t.Errorf("Expected log message not found in output: %s", buf.String())
	//}
}
