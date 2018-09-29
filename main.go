package main

import (
	"log"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaylevin/Mp32Cloud/config"
	"os"
	"encoding/json"
	"bytes"
)

var conf *config.Config

func main() {
	if _, err := os.Stat("./config.json"); os.IsNotExist(err) {
		log.Println("Could not find config file Generating a new one for you...")
		err := config.New()
		if err != nil {
			log.Fatal("Error generating config file:", err.Error())
		}
	}

	var err error
	conf, err = config.Load()
	if err != nil {
		log.Fatal("Error loading config file:\n", err.Error())
	}

	jBytes, _ := json.Marshal(conf)

	pretty := bytes.NewBuffer([]byte(``))
	json.Indent(pretty, jBytes, "", "    ")

	log.Println("Loaded config file:")
	log.Println(string(pretty.Bytes()))


	/* Set up router and listen for PUT request on localhost:8080/upload */
	r := mux.NewRouter()
	r.HandleFunc("/upload", handleMp3Upload).Methods("PUT")
	log.Println("Server is listening for mp3 upload requests on localhost:8080/upload...")
	addr := fmt.Sprintf(":%d", 8080)
	log.Fatal(http.ListenAndServe(addr, r))
}
