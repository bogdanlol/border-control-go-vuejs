package status

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("V3 status is live"))
}
