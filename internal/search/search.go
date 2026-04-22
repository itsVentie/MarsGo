package search

import (
	"marsgo/internal/evaluation"
	"marsgo/pkg/engine"
)

const (
	Inf = 50000
)

func GetBestMove(game *engine.MarseilleGame, depth int) string {
	bestMove := ""
	bestScore := -Inf

	validMoves := game.ValidMovesStr()

	for _, moveStr := range validMoves {
		gameCopy := game.Clone()
		gameCopy.PushMove(moveStr)

		score := -negamax(gameCopy, depth-1, -Inf, Inf)

		if score > bestScore {
			bestScore = score
			bestMove = moveStr
		}
	}

	return bestMove
}

func negamax(game *engine.MarseilleGame, depth, alpha, beta int) int {
	if depth == 0 {
		return evaluation.EvaluatePosition(game.Position())
	}

	moves := game.ValidMovesStr()
	if len(moves) == 0 {
		return -Inf
	}

	bestScore := -Inf

	for _, moveStr := range moves {
		gameCopy := game.Clone()
		gameCopy.PushMove(moveStr)

		score := -negamax(gameCopy, depth-1, -beta, -alpha)

		if score > bestScore {
			bestScore = score
		}
		if alpha < score {
			alpha = score
		}
		if alpha >= beta {
			break
		}
	}
	return bestScore
}
