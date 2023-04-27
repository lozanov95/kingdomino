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
	mut       sync.RWMutex
	GameRooms map[string]*GameRoom
	CloseChan chan string
}

func NewServer() *Server {
	s := &Server{
		mut:       sync.RWMutex{},
		GameRooms: make(map[string]*GameRoom),
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

func (s *Server) GetAvailableRoom() *GameRoom {
	for _, room := range s.GameRooms {
		if !room.IsFull() {
			return room
		}
	}
	new_room := NewGameRoom(s.CloseChan)
	s.GameRooms[new_room.ID] = new_room

	return new_room
}

func (s *Server) joinRoom(p *Player) {
	s.mut.Lock()
	room := s.GetAvailableRoom()
	room.Join(p)
	s.mut.Unlock()
	p.GameStateLoop(room.closeChan)
}

func (s *Server) CloseRoom(id string) {
	s.mut.Lock()
	defer s.mut.Unlock()
	close(s.GameRooms[id].closeChan)
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
