package main

import (
	"flag"
	"log"

	"github.com/lozanov95/kingdomino/backend/cmd/server"
)

var (
	PORT int
)

func main() {
	flag.IntVar(&PORT, "port", 80, "Specify the server's port")
	flag.Parse()

	srv := server.NewServer()
	if err := srv.ListenAndServe(PORT); err != nil {
		log.Fatal("failed to start the server", err)
	}
}
