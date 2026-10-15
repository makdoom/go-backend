package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxKey int

const requestIdCtxKey ctxKey = iota

func RequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-Id")

		if id == "" {
			id = uuid.NewString()
		}

		w.Header().Add("X-Request-Id", id)
		ctx := context.WithValue(r.Context(), requestIdCtxKey, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequestIdFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIdCtxKey).(string)
	return id, ok
}
