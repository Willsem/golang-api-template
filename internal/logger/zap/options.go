package zap

import (
	"io"

	"go.uber.org/zap/zapcore"
)

type options struct {
	output   io.Writer
	logLevel zapcore.Level
	name     string
	env      string
}

type OptionFunc func(*options)

//nolint:unused // this function is used by the variables of optionFunc type inside this package
func (f OptionFunc) apply(o *options) {
	f(o)
}

func Output(output io.Writer) OptionFunc {
	return func(o *options) {
		o.output = output
	}
}

func LogLevel(logLevel zapcore.Level) OptionFunc {
	return func(o *options) {
		o.logLevel = logLevel
	}
}

func Name(name string) OptionFunc {
	return func(o *options) {
		o.name = name
	}
}

func Env(env string) OptionFunc {
	return func(o *options) {
		o.env = env
	}
}
