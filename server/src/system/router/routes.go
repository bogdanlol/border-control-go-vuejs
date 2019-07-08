package router

import (
	"learning/proj1/pkg/types/routes"
	AuthHandler "learning/proj1/src/controllers/auth"
	HomeHandler "learning/proj1/src/controllers/home"
	"net/http"

	"github.com/go-xorm/xorm"
)

func MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) routes.Routes {
	AuthHandler.Init(db)
	HomeHandler.Init(db)
	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
		routes.Route{"AuthStore", "POST", "/auth/login", AuthHandler.Login},
		routes.Route{"AuthCheck", "GET", "/auth/check", AuthHandler.Check},
	}
}
