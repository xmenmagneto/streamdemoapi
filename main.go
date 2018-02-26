package main

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/xmenmagneto/streamdemoapi/Cassandra"
  	"github.com/xmenmagneto/streamdemoapi/Users"
)

type heartbeatResponse struct {
	Status string `json:"status"`
	Code int `json:"code"`
}

func main() {
 	CassandraSession := Cassandra.Session
  	defer CassandraSession.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartbeat)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}