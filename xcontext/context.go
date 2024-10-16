package xcontext

import (
	"context"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/xybor-x/snowflake"
	"github.com/xybor/x/logging"
	"github.com/xybor/x/scope"
)

type contextKey int

const (
	loggerKey contextKey = iota
	requestTimeKey
	requestUserIDKey
	scopeKey
)

func WithLogger(ctx context.Context, logger logging.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func Logger(ctx context.Context) logging.Logger {
	if val := ctx.Value(loggerKey); val != nil {
		return val.(logging.Logger)
	}

	return logging.NewSLogger(logging.LevelDebug).With("logger", "temporary")
}

func RequestID(ctx context.Context) string {
	return ctx.Value(middleware.RequestIDKey).(string)
}

func WithRequestTime(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, requestTimeKey, t)
}

func RequestTime(ctx context.Context) time.Time {
	if val := ctx.Value(requestTimeKey); val != nil {
		return val.(time.Time)
	}

	return time.Time{}
}

func WithRequestUserID(ctx context.Context, userID snowflake.ID) context.Context {
	return context.WithValue(ctx, requestUserIDKey, userID)
}

func RequestUserID(ctx context.Context) snowflake.ID {
	if val := ctx.Value(requestUserIDKey); val != nil {
		return val.(snowflake.ID)
	}

	return 0
}

func WithScope(ctx context.Context, scopes scope.Scopes) context.Context {
	return context.WithValue(ctx, scopeKey, scopes)
}

func Scope(ctx context.Context) scope.Scopes {
	if val := ctx.Value(scopeKey); val != nil {
		return val.(scope.Scopes)
	}

	return nil
}
