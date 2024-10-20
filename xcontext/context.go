package xcontext

import (
	"context"
	"net/http"
	"time"

	"github.com/xybor-x/snowflake"
	"github.com/xybor/x/logging"
	"github.com/xybor/x/scope"
	"github.com/xybor/x/session"
)

type contextKey int

const (
	loggerKey contextKey = iota
	httpRequestKey
	requestIDKey
	requestTimeKey
	requestUserIDKey
	scopeKey
	sessionKey
	sessionManagerKey
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

func WithHTTPRequest(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, httpRequestKey, r)
}

func HTTPRequest(ctx context.Context) *http.Request {
	if val := ctx.Value(httpRequestKey); val != nil {
		return val.(*http.Request)
	}

	return nil
}

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func RequestID(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
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

func WithSession(ctx context.Context, session *session.Session) context.Context {
	return context.WithValue(ctx, sessionKey, session)
}

func Session(ctx context.Context) *session.Session {
	if val := ctx.Value(sessionKey); val != nil {
		return val.(*session.Session)
	}

	return &session.Session{}
}

func WithSessionManager(ctx context.Context, manager *session.Manager) context.Context {
	return context.WithValue(ctx, sessionManagerKey, manager)
}

func SessionManager(ctx context.Context) *session.Manager {
	return ctx.Value(sessionManagerKey).(*session.Manager)
}
