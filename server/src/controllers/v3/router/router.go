package router

import (
	"net/http"

	"learning/proj1/pkg/types/routes"
	MasiniHandler "learning/proj1/src/controllers/v3/routes/masini"
	StatusHandler "learning/proj1/src/controllers/v3/status"
	"log"

	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")
		if len(token) < 1 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		log.Println("Inside V3 Middleware")

		next.ServeHTTP(w, r)
	})
}

func GetRoutes(DB *xorm.Engine) (SubRoute map[string]routes.SubRoutePackage) {
	db = DB
	StatusHandler.Init(DB)
	MasiniHandler.Init(DB)

	/* ROUTES */
	SubRoute = map[string]routes.SubRoutePackage{
		"/v3": {
			Routes: routes.Routes{
				routes.Route{"Status", "GET", "/status", StatusHandler.Index},
				routes.Route{"MasiniIndex", "GET", "/masini", MasiniHandler.Index},
				routes.Route{"MasiniStore", "POST", "/masini", MasiniHandler.Store},
				routes.Route{"MasiniEdit", "GET", "/masini/{id}/edit", MasiniHandler.Edit},
				routes.Route{"MasiniUpdate", "PUT", "/masini/{id}", MasiniHandler.Update},
				routes.Route{"MasiniDestroy", "DELETE", "/masini/{id}", MasiniHandler.Destroy},
			},
			Middleware: Middleware,
		},
	}

	return
}
