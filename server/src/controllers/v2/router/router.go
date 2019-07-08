package router

import (
	"net/http"

	"learning/proj1/pkg/types/routes"
	PersoaneHandler "learning/proj1/src/controllers/v2/routes/persoane"
	StatusHandler "learning/proj1/src/controllers/v2/status"
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

		log.Println("Inside V1 Middleware")

		next.ServeHTTP(w, r)
	})
}

func GetRoutes(DB *xorm.Engine) (SubRoute map[string]routes.SubRoutePackage) {
	db = DB
	StatusHandler.Init(DB)
	PersoaneHandler.Init(DB)

	/* ROUTES */
	SubRoute = map[string]routes.SubRoutePackage{
		"/v2": {
			Routes: routes.Routes{
				routes.Route{"Status", "GET", "/status", StatusHandler.Index},
				routes.Route{"PersoaneIndex", "GET", "/persoane", PersoaneHandler.Index},
				routes.Route{"PersoaneStore", "POST", "/persoane", PersoaneHandler.Store},
				routes.Route{"PersoaneEdit", "GET", "/persoane/{id}/edit", PersoaneHandler.Edit},
				routes.Route{"PersoaneUpdate", "PUT", "/persoane/{id}", PersoaneHandler.Update},
				routes.Route{"PersoaneDestroy", "DELETE", "/persoane/{id}", PersoaneHandler.Destroy},
			},
			Middleware: Middleware,
		},
	}

	return
}
