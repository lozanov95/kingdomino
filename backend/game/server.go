package game

import (
	"log"
	"net/http"
	"strconv"
	"sync"

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

func (s *Server) HandleJoinRoom(ws *websocket.Conn) {
	player := NewPlayer(ws)
	payload := player.GetInput()
	player.Name = payload.Name

	s.joinRoom(player)
}

func (s *Server) joinRoom(p *Player) {
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

func (s *Server) CloseRoom(id string) {
	s.mut.Lock()
	defer s.mut.Unlock()
	delete(s.GameRooms, id)
	log.Println("closed room", id)
	log.Println("active rooms", len(s.GameRooms))
}

func (s *Server) ListenAndServe(port int) error {
	go func() {
		for id := range s.CloseChan {
			s.CloseRoom(id)
		}
	}()

	http.Handle("/join", websocket.Handler(s.HandleJoinRoom))
	http.Handle("/", http.FileServer(http.Dir("./ui")))

	log.Println("Serving on", port)
	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
