package utils

import (
	"context"
	"net/http"
	"suregem/src/middleware"
)

func ForwardBackendResponse(ctx context.Context, backendResp *http.Response) {
	if backendResp == nil {
		return
	}

	appCtx := middleware.GetAppContext(ctx)
	if appCtx == nil {
		return
	}

	// forward cookies
	for _, cookie := range backendResp.Cookies() {
		http.SetCookie(appCtx.Writer, cookie)
	}

	// IMP: headers to skip
	var skipHeaders = map[string]bool{
		"Content-Length":    true,
		"Content-Type":      true,
		"Transfer-Encoding": true,
	}

	// forward headers
	for key, values := range backendResp.Header {
		if skipHeaders[key] {
			continue
		}
		for _, value := range values {
			appCtx.Writer.Header().Add(key, value)
		}
	}
}
