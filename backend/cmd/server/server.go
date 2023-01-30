package server

import (
	"log"
	"sync"

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

func (s *Server) HandleJoinRoom(ws *websocket.Conn) {
	player := game.NewPlayer(ws)
	msg, err := player.GetInput()
	if err != nil {
		log.Println("failed to get player input", err)
		return
	}
	player.Name = msg.Name

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

func (s *Server) CloseRoom(id string) {
	s.mut.Lock()
	defer s.mut.Unlock()
	delete(s.GameRooms, id)
	log.Println("closed room", id)
	log.Println("active rooms", len(s.GameRooms))
}
