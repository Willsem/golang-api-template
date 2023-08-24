package zap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Willsem/golang-api-template/internal/logger"
)

type LoggerImpl struct {
	log *zap.SugaredLogger
}

func NewLogger(opts ...OptionFunc) *LoggerImpl {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	pid := os.Getpid()

	options := &options{
		output:   os.Stdout,
		logLevel: zapcore.DebugLevel,
		name:     "unknown",
		env:      "unknown",
	}

	for _, o := range opts {
		o(options)
	}

	encoder := zapcore.NewJSONEncoder(newEncoderConfig())
	sink := zapcore.Lock(zapcore.AddSync(options.output))
	level := zap.NewAtomicLevelAt(options.logLevel)
	core := zapcore.NewCore(encoder, sink, level)

	log := zap.New(
		core,
		zap.ErrorOutput(sink),
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	).Sugar()

	return &LoggerImpl{
		log: log.With(
			logger.NameKey, options.name,
			logger.EnvKey, options.env,
			logger.InstKey, hostname,
			logger.PidKey, pid,
		),
	}
}

func encodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(l.String())
}

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey: logger.MsgKey,

		LevelKey:    logger.LevelKey,
		EncodeLevel: encodeLevel,

		TimeKey:    logger.TimeKey,
		EncodeTime: zapcore.ISO8601TimeEncoder,
	}
}

func (l *LoggerImpl) With(key string, value interface{}) logger.Logger {
	return &LoggerImpl{
		log: l.log.With(key, value),
	}
}

func (l *LoggerImpl) WithError(err error) logger.Logger {
	return l.With(logger.ErrorKey, err)
}

func (l *LoggerImpl) Debugf(msg string, args ...interface{}) {
	l.log.Debugf(msg, args...)
}

func (l *LoggerImpl) Debug(msg string) {
	l.log.Debug(msg)
}

func (l *LoggerImpl) Infof(msg string, args ...interface{}) {
	l.log.Infof(msg, args...)
}

func (l *LoggerImpl) Info(msg string) {
	l.log.Info(msg)
}

func (l *LoggerImpl) Warnf(msg string, args ...interface{}) {
	l.log.Warnf(msg, args...)
}

func (l *LoggerImpl) Warn(msg string) {
	l.log.Warn(msg)
}

func (l *LoggerImpl) Errorf(msg string, args ...interface{}) {
	l.log.Errorf(msg, args...)
}

func (l *LoggerImpl) Error(msg string) {
	l.log.Error(msg)
}

func (l *LoggerImpl) Fatalf(msg string, args ...interface{}) {
	l.log.Fatalf(msg, args...)
}

func (l *LoggerImpl) Fatal(msg string) {
	l.log.Fatal(msg)
}
