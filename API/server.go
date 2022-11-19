package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	r := mux.NewRouter()
	server := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:3000",
	}

	log.Info().Msg("Starting EscapAR API server")
	InitDB()

	// r.HandleFunc("/", ).Methods("GET")
	// r.HandleFunc("/index.html", ).Methods("GET")
	r.HandleFunc("/upload/entry", postEntry).Methods("POST")
	r.HandleFunc("/upload/path", postPath).Methods("POST")
	r.HandleFunc("/assets", getAssets).Methods("GET")

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	wait := 15 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	server.Shutdown(ctx)
	log.Info().Msg("shutting down...")
	os.Exit(0)
}

func postEntry(w http.ResponseWriter, r *http.Request) {
	var entry UserEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		panic(err)
	}
	res := addEntry(entry)
	json.NewEncoder(w).Encode(res)
}

func postPath(w http.ResponseWriter, r *http.Request) {
	res := make([]UserEntry, 0)
	var ids []string
	if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}

func addEntry(entry UserEntry) UserEntry {
	return DBAddEntry(&entry)
}

func getAssets(w http.ResponseWriter, r *http.Request) {
	var ids []string
	if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
		panic(err)
	}
	res := make([]UserEntry, 0)
	for _, id := range ids {
		entry, exists := DBGetEntry(id)
		if !exists {
			continue
		}
		res = append(res, entry)
	}
	json.NewEncoder(w).Encode(res)
}
