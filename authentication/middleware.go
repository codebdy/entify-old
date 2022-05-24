package authentication

import (
	"context"
	"net/http"
	"strings"
	"time"

	"rxdrag.com/entify/consts"
)

// ContextValue is a context key
type ContextValue map[string]interface{}

// AuthMiddleware 传递公共参数中间件
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//为了测试loading状态，生产版需要删掉
		time.Sleep(time.Duration(300) * time.Millisecond)

		reqToken := r.Header.Get(consts.AUTHORIZATION)
		splitToken := strings.Split(reqToken, consts.BEARER)
		if len(splitToken) == 2 {
			reqToken = splitToken[1]
			// 附加token
			ctx := context.WithValue(r.Context(), consts.TOKEN, reqToken)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//设置跨域
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		//w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
