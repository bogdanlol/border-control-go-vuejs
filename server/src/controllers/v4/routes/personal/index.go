package personal

import (
	"encoding/json"
	Personal "learning/proj1/src/controllers/v4/models/personal"

	"log"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var offset int
	limit := int(25)

	if keys, ok := r.URL.Query()["limit"]; ok {
		limit, _ = strconv.Atoi(keys[0])
	}
	if keys, ok := r.URL.Query()["page"]; ok {
		page, _ := strconv.Atoi(keys[0])
		offset = (page - 1) * limit
	}

	pers, err := Personal.FindBy(db, limit, offset)

	packet, err := json.Marshal(pers)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse users", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "apllication/json")
	w.Write(packet)
}
