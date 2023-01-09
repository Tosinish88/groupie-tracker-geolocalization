package main

import (
	"fmt"
	"gp/server"
	"log"
	"net/http"
	"time"
)

func main() {
	// monitoring the time needed to load

	start := time.Now()
	fmt.Println()
	fmt.Println("Welcome to Groupie-Tracker")
	fmt.Println()
	fmt.Println("Fetching data from the API...")
	fs := http.FileServer(http.Dir("static"))
	fmt.Println()

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", server.ServerHandler)
	http.HandleFunc("/artists/", server.ArtistHandler)

	fmt.Println("Data fetched successfully in", time.Since(start))

	fmt.Println()
	fmt.Println("Server is running on port 8080...")
	go log.Fatalln(http.ListenAndServe("0.0.0.0:8080", nil))

}
