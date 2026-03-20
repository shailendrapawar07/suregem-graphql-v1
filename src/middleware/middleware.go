package middleware

import (
	"github.com/gin-gonic/gin"
)

// ContextMiddleware injects AppContext into the request context
func ContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// create AppContext from Gin request
		appCtx := &AppContext{
			Writer:  c.Writer,
			Request: c.Request,
		}

		// inject AppContext into the context
		ctx := WithAppContext(c.Request.Context(), appCtx)

		// replace the original request context with the enriched one
		c.Request = c.Request.WithContext(ctx)

		c.Next() // pass control to next handler
	}
}
