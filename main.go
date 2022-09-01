package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetBlockchainHandler).Methods("GET")
	router.HandleFunc("/", WriteBlockHandler).Methods("POST")

	err := godotenv.Load("serverconf.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	httpPort := ":" + os.Getenv("PORT")
	log.Println("Listening on ", httpPort)

	server := &http.Server{
		Addr:           httpPort,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
	}()

}
