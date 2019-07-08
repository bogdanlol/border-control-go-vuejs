package status

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("V4 status is live"))
}
