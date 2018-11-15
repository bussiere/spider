package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Data struct {
	Fruit string
}

func test(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		fmt.Println("Post method")
	case http.MethodGet:
		fmt.Println("Get method")
	}
	var test = Data{}
	var reception = Data{}
	test.Fruit = "Banana"

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &reception)
	fmt.Println(reception.Fruit)
	js, err := json.Marshal(test)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	router := mux.NewRouter()
	router.HandleFunc("/test", test).Methods("POST")
	corsObj := handlers.AllowedOrigins([]string{"*"})
	http.ListenAndServe(":8066", handlers.CORS(corsObj)(router))
	// NEW
}
