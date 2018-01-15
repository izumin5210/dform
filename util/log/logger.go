package log

import "go.uber.org/zap"

var (
	logger  *zap.Logger
	sLogger *zap.SugaredLogger
)

func init() {
	SetLogger(zap.NewNop())
}

// SetLogger sets logger implementation
func SetLogger(l *zap.Logger) {
	logger = l
	sLogger = l.Sugar()
	Debug("logger replaced")
}

// Logger returns current logger object
func Logger() *zap.Logger {
	return logger
}

// Close closes the logger
func Close() error {
	var errs []error
	if err := logger.Sync(); err != nil {
		errs = append(errs, err)
	}
	if err := sLogger.Sync(); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}

// Debug logs message and key-values pairs as DEBUG
func Debug(msg string, args ...interface{}) {
	sLogger.Debugw(msg, args...)
}

// Info logs message and key-values pairs as INFO
func Info(msg string, args ...interface{}) {
	sLogger.Infow(msg, args...)
}

// Warn logs message and key-values pairs in as WARN
func Warn(msg string, args ...interface{}) {
	sLogger.Warnw(msg, args...)
}

// Error logs message and key-values pairs in as ERROR
func Error(msg string, args ...interface{}) {
	sLogger.Errorw(msg, args...)
}
