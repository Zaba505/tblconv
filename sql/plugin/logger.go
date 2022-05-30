package plugin

import (
	"io"
	"log"

	"github.com/hashicorp/go-hclog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	Logger *zap.SugaredLogger

	name string
}

func newLogger(s *zap.SugaredLogger) hclog.Logger {
	return &logger{
		Logger: s,
	}
}

func (l logger) Log(level hclog.Level, msg string, args ...interface{}) {
	switch level {
	case hclog.Trace, hclog.Debug:
		l.Debug(msg, args...)
	case hclog.NoLevel, hclog.Info: // default is info level
		l.Info(msg, args...)
	case hclog.Warn:
		l.Warn(msg, args...)
	case hclog.Error:
		l.Error(msg, args...)
	}
}

func (l logger) Trace(msg string, args ...interface{}) {
	l.Logger.Debugw(msg, args...)
}

func (l logger) Debug(msg string, args ...interface{}) {
	l.Logger.Debugw(msg, args...)
}

func (l logger) Info(msg string, args ...interface{}) {
	l.Logger.Infow(msg, args...)
}

func (l logger) Warn(msg string, args ...interface{}) {
	l.Logger.Warnw(msg, args...)
}

func (l logger) Error(msg string, args ...interface{}) {
	l.Logger.Errorw(msg, args...)
}

func (l logger) IsTrace() bool {
	return l.Logger.Desugar().Check(zapcore.DebugLevel, "trace") != nil
}

func (l logger) IsDebug() bool {
	return l.Logger.Desugar().Check(zapcore.DebugLevel, "debug") != nil
}

func (l logger) IsInfo() bool {
	return l.Logger.Desugar().Check(zapcore.InfoLevel, "info") != nil
}

func (l logger) IsWarn() bool {
	return l.Logger.Desugar().Check(zapcore.WarnLevel, "warn") != nil
}

func (l logger) IsError() bool {
	return l.Logger.Desugar().Check(zapcore.ErrorLevel, "error") != nil
}

func (l logger) ImpliedArgs() []interface{} {
	return nil
}

func (l logger) With(args ...interface{}) hclog.Logger {
	return logger{
		Logger: l.Logger.With(args...),
	}
}

func (l logger) Name() string {
	return l.name
}

func (l logger) Named(name string) hclog.Logger {
	return logger{
		Logger: l.Logger.Named(name),
		name:   name,
	}
}

func (l logger) ResetNamed(name string) hclog.Logger {
	return nil
}

var hclogToZapLevel = map[hclog.Level]zapcore.Level{
	hclog.NoLevel: zapcore.InfoLevel,
	hclog.Trace:   zapcore.DebugLevel,
	hclog.Debug:   zapcore.DebugLevel,
	hclog.Info:    zapcore.InfoLevel,
	hclog.Warn:    zapcore.WarnLevel,
	hclog.Error:   zapcore.ErrorLevel,
	hclog.Off:     zapcore.Level(127),
}

func (l logger) SetLevel(level hclog.Level) {
	l.Logger = l.Logger.Desugar().WithOptions(zap.IncreaseLevel(hclogToZapLevel[level])).Sugar()
}

func (l logger) StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger {
	return nil // TODO
}

func (l logger) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	return nil
}
