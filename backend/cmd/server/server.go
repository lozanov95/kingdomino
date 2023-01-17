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
	conns       map[int64]*ChatConn
	mut         sync.RWMutex
	GameRooms   []*GameRoom
	PlayersChan chan *game.Player
}

func NewServer() *Server {
	s := &Server{
		conns:       make(map[int64]*ChatConn),
		PlayersChan: make(chan *game.Player, 2),
	}

	go s.joinRoomLoop()

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

	s.PlayersChan <- player

	for {
		time.Sleep(500 * time.Millisecond)
	}

}

func (s *Server) joinRoomLoop() {
	for {
		p1 := <-s.PlayersChan
		room := NewGameRoom()
		room.Join(p1)

		p2 := <-s.PlayersChan
		room.Join(p2)

		p1.Conn.Write(p1.GetBoard())
		p2.Conn.Write(p2.GetBoard())
		go room.gameLoop()
	}
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
