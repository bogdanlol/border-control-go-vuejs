package masini

import (
	"encoding/json"
	Masini "learning/proj1/src/controllers/v3/models/masini"
	DB "learning/proj1/src/system/db"
	"log"
	"net/http"
)

func Store(w http.ResponseWriter, r *http.Request) {
	var masina Masini.Masina

	masina.Numar = r.PostFormValue("numar")
	masina.Culoare = r.PostFormValue("culoare")

	// encryptedPassword, err := Passwords.Encrypt(password)
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "Unable to encrypt password", http.StatusInternalServerError)
	// 	return
	// }

	// user.Password = encryptedPassword

	if err := DB.Store(db, &masina); err != nil {
		log.Println(err)
		http.Error(w, "Unable to store persoane", http.StatusInternalServerError)
		return
	}

	packet, err := json.Marshal(masina)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse users", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
