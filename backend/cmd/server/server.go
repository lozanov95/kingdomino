package server

import (
	"io"
	"log"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

type ChatConn struct {
	Id   int64
	Conn *websocket.Conn
}

type Server struct {
	conns map[int64]*ChatConn
	mut   sync.RWMutex
}

func NewServer() *Server {
	return &Server{
		conns: make(map[int64]*ChatConn),
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	log.Println("new connection from client:", ws.RemoteAddr())

	conn := &ChatConn{Id: time.Now().UnixNano(), Conn: ws}

	s.mut.Lock()
	s.conns[conn.Id] = conn
	s.mut.Unlock()

	s.readLoop(conn)
}

func (s *Server) readLoop(ChatConn *ChatConn) {
	ws := ChatConn.Conn
	defer log.Println("dropped connection", ws.RemoteAddr())
	defer delete(s.conns, ChatConn.Id)
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

		for id := range s.conns {
			if id != ChatConn.Id {
				s.conns[id].Conn.Write([]byte(msg))
			}
		}
	}
}
