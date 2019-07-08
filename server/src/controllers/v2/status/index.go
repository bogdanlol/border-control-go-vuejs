package status

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("V2 status is live"))
}
