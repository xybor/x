package logging

import (
	"errors"

	"github.com/xybor/x/errorx"
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

func Serverity2Level(s errorx.Serverity) Level {
	switch s {
	case errorx.ServerityDebug:
		return LevelDebug
	case errorx.ServerityInfo:
		return LevelInfo
	case errorx.ServerityWarn:
		return LevelWarn
	case errorx.ServerityCritical:
		return LevelCritical
	default:
		panic("invalid serverity")
	}
}

func LogError(logger Logger, err error, a ...any) {
	var serviceErr errorx.ServiceError
	switch {
	case errors.As(err, &serviceErr):
		logger.Log(Serverity2Level(serviceErr.Serverity), err.Error(), a...)
	default:
		logger.Warn(err.Error(), a...)
	}
}

func Ensure(logger Logger, f func(logger Logger)) {
	if logger != nil {
		f(logger)
	}
}
