package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lozanov95/kingdomino/cmd/server"
	"golang.org/x/net/websocket"
)

const (
	PORT = ":8080"
)

func main() {
	srv := server.NewServer()

	http.Handle("/ws", websocket.Handler(srv.HandleWS))
	http.Handle("/", http.FileServer(http.Dir("./ui")))

	log.Println("Serving on", PORT)

	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal("failed to serve", err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome")
}
