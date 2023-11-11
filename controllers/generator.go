package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortner/models"
)

func URLgenerator(w http.ResponseWriter, r *http.Request) {
	res := json.NewDecoder(r.Body)
	var data models.REQUEST_PAYLOAD
	err := res.Decode(&data)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received message: %s\n", data)
}
