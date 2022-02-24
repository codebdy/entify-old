package authentication

import (
	"context"
	"net/http"
	"time"
)

// ContextValue is a context key
type ContextValue map[string]interface{}

// AuthMiddleware 传递公共参数中间件
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//为了测试loading状态，生产版需要删掉
		time.Sleep(time.Duration(300) * time.Millisecond)
		data := ContextValue{
			"1": "one",
			"2": "two",
		}
		// 赋值
		ctx := context.WithValue(r.Context(), "data", data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//设置跨域
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		//w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
