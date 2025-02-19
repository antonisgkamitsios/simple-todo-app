package htmx

import (
	"context"
	"net/http"
)

type contextKey string

const htmxRequestKey = contextKey("htmxRequestKey")

func ContextSetHtmxRequest(r *http.Request, isHtmxRequest bool) *http.Request {
	ctx := context.WithValue(r.Context(), htmxRequestKey, isHtmxRequest)

	return r.WithContext(ctx)
}

func ContextGetHtmxRequest(ctx context.Context) bool {
	if comesFromHtmx, ok := ctx.Value(htmxRequestKey).(bool); ok {
		return comesFromHtmx
	}

	return false
}
