package uci

import (
	"bufio"
	"fmt"
	"marsgo/internal/search"
	"marsgo/pkg/engine"
	"os"
	"strings"
)

type Handler struct {
	game *engine.MarseilleGame
}

func NewHandler() *Handler {
	return &Handler{game: engine.NewMarseilleGame()}
}

func (h *Handler) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		parts := strings.Fields(cmd)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "uci":
			fmt.Println("id name MarsGo")
			fmt.Println("id author Ventie Ravelle")
			fmt.Println("uciok")
		case "isready":
			fmt.Println("readyok")
		case "position":
			h.handlePosition(parts[1:])
		case "go":
			h.handleGo()
		case "quit":
			return
		}
	}
}

func (h *Handler) handlePosition(params []string) {
	if len(params) > 0 && params[0] == "startpos" {
		h.game = engine.NewMarseilleGame()
	}
}

func (h *Handler) handleGo() {
	move := search.GetBestMove(h.game, 4)
	fmt.Printf("bestmove %s\n", move)
}
