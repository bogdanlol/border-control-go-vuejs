package verificari

import (
	"encoding/json"
	Verificari "learning/proj1/src/controllers/v5/models/verificari"
	DB "learning/proj1/src/system/db"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Destroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	newId, _ := strconv.Atoi(id)
	verificari := Verificari.Verificare{Id: int64(newId)}

	if err := DB.Destroy(db, verificari.Id, &verificari); err != nil {
		log.Println(err)
		http.Error(w, "Unable to get persoana", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(verificari)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse persoana", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
