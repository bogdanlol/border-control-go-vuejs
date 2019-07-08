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

func Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	newId, _ := strconv.Atoi(id)

	masina := Masini.Masina{Id: int64(newId)}

	if err := DB.FindBy(db, &masina); err != nil {
		log.Println(err)
		http.Error(w, "Unable to get user", http.StatusInternalServerError)
		return
	}

	//user.Password = ""

	packet, err := json.Marshal(masina)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse user", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
