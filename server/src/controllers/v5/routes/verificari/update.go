package verificari

import (
	Verificari "learning/proj1/src/controllers/v5/models/verificari"
	"time"

	DB "learning/proj1/src/system/db"

	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	newId, _ := strconv.Atoi(id)
	verificare := Verificari.Verificare{Id: int64(newId)}
	codm := r.PostFormValue("codm")
	verificare.Codm, _ = strconv.ParseInt(codm, 10, 64)
	codp := r.PostFormValue("codp")
	verificare.Codp, _ = strconv.ParseInt(codp, 10, 64)
	codo := r.PostFormValue("codo")
	verificare.Codo, _ = strconv.ParseInt(codo, 10, 64)
	layout := "2006-01-02T15:04:05.000Z"
	orap := r.PostFormValue("orap")
	verificare.Orap, _ = time.Parse(layout, orap)
	orav := r.PostFormValue("orav")
	verificare.Orav, _ = time.Parse(layout, orav)

	if err := DB.Update(db, verificare.Id, &verificare); err != nil {
		log.Println(err)
		http.Error(w, "Unable to get user", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(verificare)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse user", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
