package logger

const (
	MsgKey   = "msg"
	TimeKey  = "ts"
	LevelKey = "level"
	NameKey  = "name"
	EnvKey   = "env"
	InstKey  = "inst"
	PidKey   = "pid"

	ComponentKey = "comp"
	ErrorKey     = "err"
)

type Logger interface {
	With(key string, value interface{}) Logger
	WithError(err error) Logger

	Debugf(msg string, args ...interface{})
	Debug(msg string)

	Infof(msg string, args ...interface{})
	Info(msg string)

	Warnf(msg string, args ...interface{})
	Warn(msg string)

	Errorf(msg string, args ...interface{})
	Error(msg string)

	Fatalf(msg string, args ...interface{})
	Fatal(msg string)
}
