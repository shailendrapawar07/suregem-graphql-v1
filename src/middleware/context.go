package middleware

import (
	"context"
	"net/http"
)

// AppContext holds HTTP objects needed in resolvers
type AppContext struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

// contextKey type for safe context storage
type contextKey string

const AppContextKey contextKey = "app_context"

// WithAppContext injects AppContext into context
func WithAppContext(ctx context.Context, appCtx *AppContext) context.Context {
	return context.WithValue(ctx, AppContextKey, appCtx)
}

// GetAppContext retrieves AppContext from context
func GetAppContext(ctx context.Context) *AppContext {
	if appCtx, ok := ctx.Value(AppContextKey).(*AppContext); ok {
		return appCtx
	}
	return nil
}
