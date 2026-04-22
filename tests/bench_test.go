package tests

import (
	"marsgo/internal/search"
	"marsgo/pkg/engine"
	"testing"
)

func BenchmarkSearchDepth4(b *testing.B) {
	game := engine.NewMarseilleGame()
	game.PushMove("e2e4")
	game.PushMove("e7e5")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search.GetBestMove(game, 4)
	}
}

func BenchmarkMoveGeneration(b *testing.B) {
	game := engine.NewMarseilleGame()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.ValidMovesStr()
	}
}
