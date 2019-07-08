package masini

import (
	"encoding/json"
	Masini "learning/proj1/src/controllers/v3/models/masini"
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
	masini := Masini.Masina{Id: int64(newId)}

	if err := DB.Destroy(db, masini.Id, &masini); err != nil {
		log.Println(err)
		http.Error(w, "Unable to get persoana", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(masini)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse persoana", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
