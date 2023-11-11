package routes

import (
	"fmt"
	"net/http"
	"url-shortner/controllers"

	"github.com/gorilla/mux"
)

var PORT = 8000

func IndexRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.URLgenerator).Methods("POST")
	http.Handle("/", r)
	fmt.Println("Server up and running😁 on PORT", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)

	if err != nil {
		fmt.Println("An error occured", err)
		return
	}

}
