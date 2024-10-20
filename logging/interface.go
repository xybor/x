package logging

import (
	"github.com/xybor/x/xerror"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelCritical
)

type Logger interface {
	With(a ...any) Logger
	Log(level Level, msg string, a ...any)
	Debug(msg string, a ...any)
	Info(msg string, a ...any)
	Warn(msg string, a ...any)
	Critical(msg string, a ...any)
}

func Serverity2Level(s xerror.Serverity) Level {
	switch s {
	case xerror.ServerityDebug:
		return LevelDebug
	case xerror.ServerityInfo:
		return LevelInfo
	case xerror.ServerityWarn:
		return LevelWarn
	case xerror.ServerityCritical:
		return LevelCritical
	default:
		panic("invalid serverity")
	}
}

func LogError(logger Logger, err error, a ...any) {
	parameter := append(a, "details", xerror.MessageOf(err))
	logger.Log(Serverity2Level(xerror.ServerityOf(err, xerror.ServerityWarn)), err.Error(), parameter...)
}
