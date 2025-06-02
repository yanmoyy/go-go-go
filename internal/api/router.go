package api

import "net/http"

type router struct {
	mux *http.ServeMux
}

func newRouter() *router {
	return &router{mux: http.NewServeMux()}
}

func (r *router) Get(endpoint string, handlerFn http.HandlerFunc) {
	r.mux.HandleFunc("GET"+" "+endpoint, handlerFn)
}

func (r *router) Post(endpoint string, handlerFn http.HandlerFunc) {
	r.mux.HandleFunc("POST"+" "+endpoint, handlerFn)
}

func (r *router) Put(endpoint string, handlerFn http.HandlerFunc) {
	r.mux.HandleFunc("PUT"+" "+endpoint, handlerFn)
}

func (r *router) Delete(endpoint string, handlerFn http.HandlerFunc) {
	r.mux.HandleFunc("DELETE"+" "+endpoint, handlerFn)
}

func (r *router) ServeStatic(path string, dir string) {
	fs := http.FileServer(http.Dir(dir))
	r.mux.Handle(path, http.StripPrefix(path, fs))
}
