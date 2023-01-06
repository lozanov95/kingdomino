package server

import (
	"io"
	"log"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	log.Println("new connection from client:", ws.RemoteAddr())

	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	defer log.Println("dropped connection", ws.RemoteAddr())
	defer delete(s.conns, ws)
	defer ws.Close()
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("read error", err)
			continue
		}
		msg := string(buf[:n])
		ws.Write([]byte("go the msg from you " + msg))
	}
}
