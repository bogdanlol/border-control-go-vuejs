package router

import (
	"net/http"

	"learning/proj1/pkg/types/routes"
	VerificariHandler "learning/proj1/src/controllers/v5/routes/verificari"
	StatusHandler "learning/proj1/src/controllers/v5/status"
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

		log.Println("Inside V5 Middleware")

		next.ServeHTTP(w, r)
	})
}

func GetRoutes(DB *xorm.Engine) (SubRoute map[string]routes.SubRoutePackage) {
	db = DB
	StatusHandler.Init(DB)
	VerificariHandler.Init(DB)

	/* ROUTES */
	SubRoute = map[string]routes.SubRoutePackage{
		"/v5": {
			Routes: routes.Routes{
				routes.Route{"Status", "GET", "/status", StatusHandler.Index},
				routes.Route{"VerificariIndex", "GET", "/verificari", VerificariHandler.Index},
				routes.Route{"VerificariStore", "POST", "/verificari", VerificariHandler.Store},
				routes.Route{"VerificariEdit", "GET", "/verificari/{id}/edit", VerificariHandler.Edit},
				routes.Route{"VerificariUpdate", "PUT", "/verificari/{id}", VerificariHandler.Update},
				routes.Route{"VerificariDestroy", "DELETE", "/verificari/{id}", VerificariHandler.Destroy},
			},
			Middleware: Middleware,
		},
	}

	return
}
