package personal

import (
	"encoding/json"
	Personal "learning/proj1/src/controllers/v4/models/personal"
	DB "learning/proj1/src/system/db"
	"log"
	"net/http"
)

func Store(w http.ResponseWriter, r *http.Request) {
	var pers Personal.Personal

	pers.NUME = r.PostFormValue("nume")
	pers.PREN = r.PostFormValue("pren")
	pers.RANG = r.PostFormValue("rang")
	// encryptedPassword, err := Passwords.Encrypt(password)
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "Unable to encrypt password", http.StatusInternalServerError)
	// 	return
	// }

	// user.Password = encryptedPassword

	if err := DB.Store(db, &pers); err != nil {
		log.Println(err)
		http.Error(w, "Unable to store persoane", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(pers)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse users", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
