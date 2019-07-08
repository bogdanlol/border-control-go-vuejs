package personal

import (
	"encoding/json"
	Personal "learning/proj1/src/controllers/v4/models/personal"

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
	pers := Personal.Personal{Id: int64(newId)}

	if err := DB.Destroy(db, pers.Id, &pers); err != nil {
		log.Println(err)
		http.Error(w, "Unable to get persoana", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(pers)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse persoana", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
