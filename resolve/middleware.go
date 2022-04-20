package resolve

import (
	"context"
	"net/http"

	"rxdrag.com/entity-engine/consts"
)

func LoadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), consts.LOADERS, CreateDataLoaders())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
