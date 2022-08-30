package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("serverconf.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	httpPort := ":" + os.Getenv("PORT")
	log.Println("Listening on ", httpPort)

	http.ListenAndServe(httpPort, nil)

	router := mux.NewRouter()
	router.HandleFunc("/", getBlockchainHandler).Methods("GET")
	router.HandleFunc("/", writeBlockHandler).Methods("POST")
}
