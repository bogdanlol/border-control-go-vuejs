package persoane

import (
	"encoding/json"
	Persoane "learning/proj1/src/controllers/v2/models/persoane"
	DB "learning/proj1/src/system/db"
	"log"
	"net/http"
	"strconv"
)

func Store(w http.ResponseWriter, r *http.Request) {
	var persoana Persoane.Persoana

	persoana.NUME = r.PostFormValue("nume")
	persoana.PREN = r.PostFormValue("pren")

	persoana.CNP = r.PostFormValue("cnp")
	varsta := r.PostFormValue("varsta")

	persoana.Varsta, _ = strconv.ParseInt(varsta, 10, 64)

	codm := r.PostFormValue("codm")
	persoana.Codm, _ = strconv.ParseInt(codm, 10, 64)

	// encryptedPassword, err := Passwords.Encrypt(password)
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "Unable to encrypt password", http.StatusInternalServerError)
	// 	return
	// }

	// user.Password = encryptedPassword

	if err := DB.Store(db, &persoana); err != nil {
		log.Println(err)
		http.Error(w, "Unable to store persoane", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(persoana)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse users", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
