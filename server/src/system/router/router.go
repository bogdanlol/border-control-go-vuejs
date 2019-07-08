package router

import (
	"learning/proj1/pkg/types/routes"

	V1SubRoutes "learning/proj1/src/controllers/v1/router"
	V2SubRoutes "learning/proj1/src/controllers/v2/router"
	V3SubRoutes "learning/proj1/src/controllers/v3/router"
	V4SubRoutes "learning/proj1/src/controllers/v4/router"
	V5SubRoutes "learning/proj1/src/controllers/v5/router"

	"github.com/go-xorm/xorm"

	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func (r *Router) Init(db *xorm.Engine) {
	r.Router.Use(MiddleWare)

	baseRoutes := GetRoutes(db)
	for _, route := range baseRoutes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	v1SubRoutes := V1SubRoutes.GetRoutes(db)
	for name, pack := range v1SubRoutes {
		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}
	v2SubRoutes := V2SubRoutes.GetRoutes(db)
	for name, pack := range v2SubRoutes {
		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}
	v3SubRoutes := V3SubRoutes.GetRoutes(db)
	for name, pack := range v3SubRoutes {
		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}
	v4SubRoutes := V4SubRoutes.GetRoutes(db)
	for name, pack := range v4SubRoutes {
		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}
	v5SubRoutes := V5SubRoutes.GetRoutes(db)
	for name, pack := range v5SubRoutes {
		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}
}

func (r *Router) AttachSubRouterWithMiddleware(path string, subroutes routes.Routes, middleware mux.MiddlewareFunc) (SubRouter *mux.Router) {

	SubRouter = r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return
}

func NewRouter() (r Router) {
	r.Router = mux.NewRouter().StrictSlash(true)

	return
}
