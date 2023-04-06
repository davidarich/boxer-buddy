package interop

import (
	"net/http"
	"sync"

	"github.com/davidarich/boxer-buddy/internal/adapters/log"
	"github.com/davidarich/boxer-buddy/internal/ports"
	"github.com/gorilla/websocket"
)

// server for handling websocket connections from a UI
type Server struct {
	addr           string
	connectionPool *ConnectionPool
	poolMutex      *sync.Mutex
	messageType    int

	logger log.Logger
	router ports.MessageRouter
}

func (s *Server) handleSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.Error("failed to upgrade connection from http to websocket", err)
		return
	}
	s.connectionPool.Add(conn, s.poolMutex)
	defer conn.Close()
	s.logger.Info("websocket connection listening for messages")
	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			s.logger.Error(err)
			break
		}
		s.messageType = mt
		s.logger.Info("received websocket message")
		response := make(chan []byte)
		s.router.Route(msg, response)
		close(response)
	}
	s.connectionPool.Remove(conn, s.poolMutex)
}

func (s *Server) Write(message []byte) {
	for i := range s.connectionPool.Connections {
		if s.connectionPool.Connections[i] == nil {
			continue
		}
		err := s.connectionPool.Connections[i].WriteMessage(s.messageType, message)
		if err != nil {
			s.logger.Error(err)
		}
	}
}

func (s *Server) StartIO() {
	s.logger.Info("Starting websocket")
	http.HandleFunc("/", s.handleSocket)
	go http.ListenAndServe(s.addr, nil)
}

func NewServer(logger log.Logger, router ports.MessageRouter, cfg ports.Settings) *Server {
	settings, _ := cfg.Get()

	return &Server{
		addr:           settings.InteropOptions.GetAddress(),
		connectionPool: NewConnectionPool(),
		poolMutex:      &sync.Mutex{},
		logger:         logger,
		router:         router,
	}
}
