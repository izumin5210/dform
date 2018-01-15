package log

import (
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/grpclog"
)

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

	// See https://github.com/grpc-ecosystem/go-grpc-middleware/blob/d0c54e68681ec7999ac17864470f3bee6521ba2b/logging/zap/grpclogger.go#L13-L18
	zgl := &zapGrpcLogger{logger.With(zap.String("system", "grpc"), zap.Bool("grpc_log", true))}
	grpclog.SetLogger(zgl)

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

// see https://github.com/grpc-ecosystem/go-grpc-middleware/blob/d0c54e68681ec7999ac17864470f3bee6521ba2b/logging/zap/grpclogger.go#L20-L46
type zapGrpcLogger struct {
	logger *zap.Logger
}

func (l *zapGrpcLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(fmt.Sprint(args...))
}

func (l *zapGrpcLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal(fmt.Sprintf(format, args...))
}

func (l *zapGrpcLogger) Fatalln(args ...interface{}) {
	l.logger.Fatal(fmt.Sprint(args...))
}

func (l *zapGrpcLogger) Print(args ...interface{}) {
	l.logger.Debug(fmt.Sprint(args...))
}

func (l *zapGrpcLogger) Printf(format string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, args...))
}

func (l *zapGrpcLogger) Println(args ...interface{}) {
	l.logger.Debug(fmt.Sprint(args...))
}
