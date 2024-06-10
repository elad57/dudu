package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"

	// "github.com/elad57/dudu/bl"
	"github.com/elad57/dudu/database"
	"github.com/elad57/dudu/modules"
)


func main() {
	router := mux.NewRouter()
	
	db := database.InitDb("dudu.db")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var url modules.Url
    	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
        	http.Error(w, err.Error(), http.StatusBadRequest)
        	return
    	}

		err := database.CreateRoute(url.Original_url, url.Short_url, db)
		if  err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("everything is cool!"))
	}).Methods("POST")

	router.HandleFunc("/{shortUrl}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		short_url := vars["shortUrl"]
		original_url, err := database.GetOriginalUrlOfShortUrl(short_url, db)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, original_url, http.StatusFound)
	}).Methods("GET")

	port := 8080
	
	fmt.Println(fmt.Sprintf("server is up on port %d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))

}