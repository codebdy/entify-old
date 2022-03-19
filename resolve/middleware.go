package resolve

import (
	"context"
	"net/http"
)

func LoadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "loaders", CreateDataLoaders())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
