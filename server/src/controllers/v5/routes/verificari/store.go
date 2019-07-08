package verificari

import (
	"encoding/json"
	Verificari "learning/proj1/src/controllers/v5/models/verificari"
	DB "learning/proj1/src/system/db"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Store(w http.ResponseWriter, r *http.Request) {
	var verificare Verificari.Verificare
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

	// user.Password = encryptedPassword

	if err := DB.Store(db, &verificare); err != nil {
		log.Println(err)
		http.Error(w, "Unable to store persoane", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(verificare)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse users", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
