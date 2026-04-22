package notation

import (
	"fmt"
	"marsgo/pkg/engine"
	"strings"
)

func EncodeMarseilleFEN(game *engine.MarseilleGame) string {
	baseFEN := game.FEN()
	return fmt.Sprintf("%s %d", baseFEN, game.SubTurn())
}

func DecodeMarseilleFEN(fen string) (string, engine.SubTurn) {
	parts := strings.Split(fen, " ")
	if len(parts) < 7 {
		return fen, engine.FirstHalf
	}

	subTurnStr := parts[6]
	actualFEN := strings.Join(parts[:6], " ")

	if subTurnStr == "2" {
		return actualFEN, engine.SecondHalf
	}
	return actualFEN, engine.FirstHalf
}
