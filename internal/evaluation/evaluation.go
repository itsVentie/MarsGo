package evaluation

import (
	"github.com/notnil/chess"
)

var pieceValues = map[chess.PieceType]int{
	chess.Pawn:   100,
	chess.Knight: 320,
	chess.Bishop: 330,
	chess.Rook:   500,
	chess.Queen:  900,
	chess.King:   20000,
}

func EvaluatePosition(pos *chess.Position) int {
	score := 0
	board := pos.Board()

	for sq := chess.A1; sq <= chess.H8; sq++ {
		piece := board.Piece(sq)
		if piece != chess.NoPiece {
			val := pieceValues[piece.Type()]
			if piece.Color() == chess.White {
				score += val
			} else {
				score -= val
			}
		}
	}

	if pos.Turn() == chess.Black {
		return -score
	}
	return score
}
