package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/lozanov95/kingdomino/backend/cmd/server"
	"golang.org/x/net/websocket"
)

var (
	PORT int
)

func main() {
	flag.IntVar(&PORT, "port", 80, "Specify the server's port")
	flag.Parse()

	srv := server.NewServer()

	http.Handle("/ws", websocket.Handler(srv.HandleWS))
	http.Handle("/", http.FileServer(http.Dir("./ui")))

	log.Println("Serving on", PORT)

	if err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil); err != nil {
		log.Fatal("failed to serve", err)
	}
}