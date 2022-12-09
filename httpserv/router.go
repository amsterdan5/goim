package httpserv

import (
	"net/http"
)

type Router interface {
	Route()
}

// 加载路由执行方法
func Route(servMux *http.ServeMux, method, path string, handler http.Handler) {
	if len(method) == 0 {
		panic("请求方法不能为空")
	}

	if len(path) == 0 {
		panic("请求地址不能为空")
	}

	servMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.NotFound(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
