package router

import "net/http"

type FBRouter struct {
	addr    string
	handler http.Handler
}

func NewFBRouter(addr string, handler http.Handler) *FBRouter {
	return &FBRouter{addr, handler}
}

func (r *FBRouter) Run() error {
	return http.ListenAndServe(r.addr, r.handler)
}
