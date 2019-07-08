package persoane

import (
	Persoane "learning/proj1/src/controllers/v2/models/persoane"

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
	persoana := Persoane.Persoana{Id: int64(newId)}
	persoana.NUME = r.PostFormValue("nume")
	persoana.PREN = r.PostFormValue("pren")
	persoana.CNP = r.PostFormValue("cnp")
	varsta := r.PostFormValue("varsta")

	persoana.Varsta, _ = strconv.ParseInt(varsta, 10, 64)

	codm := r.PostFormValue("codm")
	persoana.Codm, _ = strconv.ParseInt(codm, 10, 64)

	if err := DB.Update(db, persoana.Id, &persoana); err != nil {
		log.Println(err)
		http.Error(w, "Unable to get user", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(persoana)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse user", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
