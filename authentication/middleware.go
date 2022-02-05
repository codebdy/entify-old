package authentication

import (
	"context"
	"net/http"
)

// ContextValue is a context key
type ContextValue map[string]interface{}

// AuthMiddleware 传递公共参数中间件
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		data := ContextValue{
			"1": "one",
			"2": "two",
		}
		// 赋值
		ctx := context.WithValue(r.Context(), "data", data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
