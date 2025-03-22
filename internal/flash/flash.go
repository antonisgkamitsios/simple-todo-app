package flash

import (
	"context"
	"net/http"
)

type contextKey string

const flashKey = contextKey("flash")

const (
	FlashTypeSuccess string = "success"
	FlashTypeError   string = "error"
	FlashTypeWarning string = "warning"
)

type Flash struct {
	FlashType string
	Message   string
}

func ContextGetFlash(ctx context.Context) Flash {
	flash, ok := ctx.Value(flashKey).(Flash)
	if !ok {
		return Flash{}
	}

	return flash
}

func ContextSetFlash(r *http.Request, flashType, message string) *http.Request {
	flash := Flash{FlashType: flashType, Message: message}

	ctx := context.WithValue(r.Context(), flashKey, flash)

	return r.WithContext(ctx)
}
