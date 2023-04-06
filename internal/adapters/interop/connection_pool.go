package interop

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct {
	Id   int
	Conn *websocket.Conn
}

func NewConnection() *Connection {
	return &Connection{}
}

type ConnectionPool struct {
	Connections []*websocket.Conn
}

func (p *ConnectionPool) Add(c *websocket.Conn, m *sync.Mutex) (err error) {
	m.Lock()
	p.Connections = append(p.Connections, c)
	m.Unlock()
	return
}

func (p *ConnectionPool) Remove(c *websocket.Conn, m *sync.Mutex) (err error) {
	m.Lock()
	for i := range p.Connections {
		if p.Connections[i] == c {
			p.Connections[i] = nil
		}
	}
	m.Unlock()
	return
}

func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		Connections: make([]*websocket.Conn, 1),
	}
}
