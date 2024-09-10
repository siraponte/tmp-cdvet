// Provides a custom router that embeds an http.ServeMux and adds nested prefix and middleware support.
package router

import (
	"net/http"
	"path"
	"slices"
	"strings"
)

type Router interface {
	ListenAndServe(addr string) error
	Group(fn GroupMember)
	GroupPrefix(prefix string, fn GroupMember)
	Use(middlewares ...Middleware)
	Get(path string, fn http.HandlerFunc, middlewares ...Middleware)
	Post(path string, fn http.HandlerFunc, middlewares ...Middleware)
	Put(path string, fn http.HandlerFunc, middlewares ...Middleware)
	Delete(path string, fn http.HandlerFunc, middlewares ...Middleware)
	Head(path string, fn http.HandlerFunc, middlewares ...Middleware)
	Options(path string, fn http.HandlerFunc, middlewares ...Middleware)
}

type (
	Middleware func(http.Handler) http.Handler

	GroupMember func(r Router)

	router struct {
		*http.ServeMux
		prefix string
		chain  []Middleware
	}
)

func NewRouter(prefix string, middlewares ...Middleware) Router {
	cleanPrefix := cleanPath(prefix)

	return &router{ServeMux: &http.ServeMux{}, prefix: cleanPrefix, chain: middlewares}
}

func (r *router) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, r)
}

func (r *router) Group(fn GroupMember) {
	// Create a new subrouter with the new prefix and the same chain.
	subrouter := &router{ServeMux: r.ServeMux, prefix: r.prefix, chain: slices.Clone(r.chain)}

	// Call the provided function with the new subrouter.
	fn(subrouter)
}

func (r *router) GroupPrefix(prefix string, fn GroupMember) {
	cleanPrefix := cleanPath(r.prefix + prefix)

	// Create a new subrouter with the new prefix and the same chain.
	subrouter := &router{ServeMux: r.ServeMux, prefix: cleanPrefix, chain: slices.Clone(r.chain)}

	// Call the provided function with the new subrouter.
	fn(subrouter)
}

func (r *router) Use(middlewares ...Middleware) {
	r.chain = append(r.chain, middlewares...)
}

func (r *router) Get(path string, fn http.HandlerFunc, middlewares ...Middleware) {
	r.handle(http.MethodGet, path, fn, middlewares)
}

func (r *router) Post(path string, fn http.HandlerFunc, middlewares ...Middleware) {
	r.handle(http.MethodPost, path, fn, middlewares)
}

func (r *router) Put(path string, fn http.HandlerFunc, middlewares ...Middleware) {
	r.handle(http.MethodPut, path, fn, middlewares)
}

func (r *router) Delete(path string, fn http.HandlerFunc, middlewares ...Middleware) {
	r.handle(http.MethodDelete, path, fn, middlewares)
}

func (r *router) Head(path string, fn http.HandlerFunc, middlewares ...Middleware) {
	r.handle(http.MethodHead, path, fn, middlewares)
}

func (r *router) Options(path string, fn http.HandlerFunc, middlewares ...Middleware) {
	r.handle(http.MethodOptions, path, fn, middlewares)
}

/* -------------------------------------------------------------------------- */
/*                                  INTERNAL                                  */
/* -------------------------------------------------------------------------- */

func (r *router) handle(method, path string, fn http.HandlerFunc, middlewares []Middleware) {
	// Add the prefix to the path.
	fullPath := cleanPath(r.prefix + path)

	// Add the route to the router.
	r.Handle(method+" "+fullPath, r.wrap(fn, middlewares))
}

func (r *router) wrap(fn http.HandlerFunc, middlewares []Middleware) http.Handler {
	// Create a new handler with the provided function.
	newHandler := http.Handler(fn)

	// Append the new middlewares to the list.
	middlewares = append(middlewares, r.chain...)

	// Reverse the middleware slice.
	slices.Reverse(middlewares)

	// Apply the middleware to the handler.
	for _, mid := range middlewares {
		newHandler = mid(newHandler)
	}

	return newHandler
}

func cleanPath(p string) string {
	// Use path.Clean to clean up paths, but ensure it retains a leading "/"
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}

	return path.Clean(p)
}
