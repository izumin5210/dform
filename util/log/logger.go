package log

import "go.uber.org/zap"

var (
	logger = zap.NewNop().Sugar()
)

// SetLogger sets logger implementation
func SetLogger(l *zap.SugaredLogger) {
	logger = l
}

// Close closes the logger
func Close() error {
	return logger.Sync()
}

// Debug logs message and key-values pairs as DEBUG
func Debug(msg string, args ...interface{}) {
	logger.Debugw(msg, args)
}

// Info logs message and key-values pairs as INFO
func Info(msg string, args ...interface{}) {
	logger.Infow(msg, args)
}

// Warn logs message and key-values pairs in as WARN
func Warn(msg string, args ...interface{}) {
	logger.Warnw(msg, args)
}

// Error logs message and key-values pairs in as ERROR
func Error(msg string, args ...interface{}) {
	logger.Errorw(msg, args)
}
