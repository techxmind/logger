package logger

import (
	"flag"
	"net/http"
	"os"
	"regexp"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	_atom    = zap.NewAtomicLevel()
	_slogger *zap.SugaredLogger
	_logger  *zap.Logger
)

func init() {
	var (
		lvl            = zap.InfoLevel
		levelFlagName  = "log-level"
		levelFlagUsage = "minimum enabled logging level. debug|info|warn|error|dpanic|panic|fatal"
	)

	flag.Var(&lvl, levelFlagName, levelFlagUsage)

	logArgs := regexp.MustCompile(`-{1,2}` + levelFlagName + `(?:\s+|\s*=\s*)(\w+)`).
		FindString(strings.Join(os.Args[1:], " "))

	if logArgs != "" {
		// use local FlagSet to parse immediately
		flagSet := flag.NewFlagSet("logger", flag.ContinueOnError)
		flagSet.Var(&lvl, levelFlagName, levelFlagUsage)
		flagSet.Parse(regexp.MustCompile(`\s+`).Split(logArgs, 2))
		_atom.SetLevel(lvl)
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	//encoderCfg.TimeKey = ""

	_logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		_atom,
	))
	_slogger = _logger.Sugar()
}

func Debug(args ...interface{}) {
	_slogger.Debug(args...)
}

func Debugw(msg string, kvs ...interface{}) {
	_slogger.Debugw(msg, kvs...)
}

func Debugf(msg string, args ...interface{}) {
	_slogger.Debugf(msg, args...)
}

func Info(args ...interface{}) {
	_slogger.Info(args...)
}

func Infow(msg string, kvs ...interface{}) {
	_slogger.Infow(msg, kvs...)
}

func Infof(msg string, args ...interface{}) {
	_slogger.Infof(msg, args...)
}

func Warn(args ...interface{}) {
	_slogger.Warn(args...)
}

func Warnw(msg string, kvs ...interface{}) {
	_slogger.Warnw(msg, kvs...)
}

func Warnf(msg string, args ...interface{}) {
	_slogger.Warnf(msg, args...)
}

func Error(args ...interface{}) {
	_slogger.Error(args...)
}

func Errorw(msg string, kvs ...interface{}) {
	_slogger.Errorw(msg, kvs...)
}

func Errorf(msg string, args ...interface{}) {
	_slogger.Errorf(msg, args...)
}

func DPanic(args ...interface{}) {
	_slogger.DPanic(args...)
}

func DPanicw(msg string, kvs ...interface{}) {
	_slogger.DPanicw(msg, kvs...)
}

func DPanicf(msg string, args ...interface{}) {
	_slogger.DPanicf(msg, args...)
}

func Panic(args ...interface{}) {
	_slogger.Panic(args...)
}

func Panicw(msg string, kvs ...interface{}) {
	_slogger.Panicw(msg, kvs...)
}

func Panicf(msg string, args ...interface{}) {
	_slogger.Panicf(msg, args...)
}

func Fatal(args ...interface{}) {
	_slogger.Fatal(args...)
}

func Fatalw(msg string, kvs ...interface{}) {
	_slogger.Fatalw(msg, kvs...)
}

func Fatalf(msg string, args ...interface{}) {
	_slogger.Fatalf(msg, args...)
}

func Named(s string) *zap.SugaredLogger {
	return _slogger.Named(s)
}

func With(args ...interface{}) *zap.SugaredLogger {
	return _slogger.With(args...)
}

func HttpHandler() http.Handler {
	return _atom
}

func Sync() {
	_slogger.Sync()
}

func Logger() *zap.Logger {
	return _logger
}
