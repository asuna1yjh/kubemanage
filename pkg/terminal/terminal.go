package terminal

import "github.com/gorilla/websocket"

type TerminalSession struct {
	Conn *websocket.Conn
}
type TerminalSessioner interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
	Close()
}

func (t *TerminalSession) Read(p []byte) (n int, err error) {
	_, message, err := t.Conn.ReadMessage()
	if err != nil {
		return 0, err
	}
	if string(message) == "exit" {
		t.Close()
	}
	return copy(p, message), nil
}
func (t *TerminalSession) Write(p []byte) (n int, err error) {
	err = t.Conn.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (t *TerminalSession) Close() {
	t.Conn.Close()
}

type TerminalMessage struct {
	Operation string `json:"operation"`
	Data      string `json:"data"`
}

func NewTerminalSession(conn *websocket.Conn) TerminalSessioner {
	return &TerminalSession{Conn: conn}
}
