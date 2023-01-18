package server

import (
	"io"
	"log"
	"sync"
	"time"

	"github.com/lozanov95/kingdomino/backend/cmd/game"
	"golang.org/x/net/websocket"
)

type ChatConn struct {
	Id   int64
	Conn *websocket.Conn
}

type Server struct {
	conns     map[int64]*ChatConn
	mut       sync.RWMutex
	GameRooms map[string]*GameRoom
	CloseChan chan string
}

func NewServer() *Server {
	s := &Server{
		conns:     make(map[int64]*ChatConn),
		mut:       sync.RWMutex{},
		GameRooms: map[string]*GameRoom{},
		CloseChan: make(chan string, 10),
	}

	return s
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	log.Println("new connection from client:", ws.RemoteAddr())

	conn := &ChatConn{Id: time.Now().UnixNano(), Conn: ws}

	s.mut.Lock()
	s.conns[conn.Id] = conn
	s.mut.Unlock()

	s.readLoop(conn)
}

func (s *Server) HandleJoinRoom(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	n, err := ws.Read(buf)
	if err != nil {
		if err == io.EOF {
			log.Println("eof err")
			return
		}
		log.Println(err)
		ws.Close()
		return
	}
	player := game.NewPlayer(buf[:n], ws)

	s.joinRoom(player)
}

func (s *Server) joinRoom(p *game.Player) {
	s.mut.Lock()
	joined := false
	for _, room := range s.GameRooms {
		if !room.IsFull() {
			room.Join(p)
			joined = true
			break
		}
	}
	if !joined {
		room := NewGameRoom(s.CloseChan)
		room.Join(p)
		s.GameRooms[room.ID] = room
	}
	s.mut.Unlock()

	p.GameStateLoop()
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

func (s *Server) CloseRoom(id string) {
	s.mut.Lock()
	defer s.mut.Unlock()
	delete(s.GameRooms, id)
	log.Println("closed room", id)
	log.Println("active rooms", len(s.GameRooms))
}
