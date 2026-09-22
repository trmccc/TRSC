package ws

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"

	"trsc/internal/process"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleConsole(w http.ResponseWriter, r *http.Request, proc *process.Proc, onInput func(string) error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for _, line := range proc.History() {
		_ = conn.WriteJSON(map[string]any{"type": "log", "t": line.T, "text": line.Text})
	}
	_ = conn.WriteJSON(map[string]any{"type": "ready"})

	ch := proc.Subscribe()
	defer proc.Unsubscribe(ch)

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				return
			}
			var msg struct {
				Type string `json:"type"`
				Data string `json:"data"`
			}
			if json.Unmarshal(data, &msg) != nil {
				continue
			}
			if msg.Type == "input" && onInput != nil {
				_ = onInput(msg.Data)
			}
		}
	}()

	for {
		select {
		case <-done:
			return
		case line, ok := <-ch:
			if !ok {
				return
			}
			if err := conn.WriteJSON(map[string]any{"type": "log", "t": line.T, "text": line.Text}); err != nil {
				return
			}
		}
	}
}
