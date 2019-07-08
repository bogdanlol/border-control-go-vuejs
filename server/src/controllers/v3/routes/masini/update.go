package masini

import (
	Masini "learning/proj1/src/controllers/v3/models/masini"

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
	masina := Masini.Masina{Id: int64(newId)}
	masina.Numar = r.PostFormValue("numar")
	masina.Culoare = r.PostFormValue("culoare")

	if err := DB.Update(db, masina.Id, &masina); err != nil {
		log.Println(err)
		http.Error(w, "Unable to get user", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(masina)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse user", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
