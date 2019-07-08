package personal

import (
	Personal "learning/proj1/src/controllers/v4/models/personal"

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
	pers := Personal.Personal{Id: int64(newId)}
	pers.NUME = r.PostFormValue("nume")
	pers.PREN = r.PostFormValue("pren")
	pers.RANG = r.PostFormValue("rang")

	if err := DB.Update(db, pers.Id, &pers); err != nil {
		log.Println(err)
		http.Error(w, "Unable to get user", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(pers)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse user", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
