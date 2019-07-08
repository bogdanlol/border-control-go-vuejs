package router

import (
	"net/http"

	"learning/proj1/pkg/types/routes"
	PersonalHandler "learning/proj1/src/controllers/v4/routes/personal"
	StatusHandler "learning/proj1/src/controllers/v4/status"
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

		log.Println("Inside V4 Middleware")

		next.ServeHTTP(w, r)
	})
}

func GetRoutes(DB *xorm.Engine) (SubRoute map[string]routes.SubRoutePackage) {
	db = DB
	StatusHandler.Init(DB)
	PersonalHandler.Init(DB)

	/* ROUTES */
	SubRoute = map[string]routes.SubRoutePackage{
		"/v4": {
			Routes: routes.Routes{
				routes.Route{"Status", "GET", "/status", StatusHandler.Index},
				routes.Route{"PersonalIndex", "GET", "/personal", PersonalHandler.Index},
				routes.Route{"PersonalStore", "POST", "/personal", PersonalHandler.Store},
				routes.Route{"PersonalEdit", "GET", "/personal/{id}/edit", PersonalHandler.Edit},
				routes.Route{"PersonalUpdate", "PUT", "/personal/{id}", PersonalHandler.Update},
				routes.Route{"PersonalDestroy", "DELETE", "/personal/{id}", PersonalHandler.Destroy},
			},
			Middleware: Middleware,
		},
	}

	return
}
