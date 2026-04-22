package engine

import "github.com/notnil/chess"

func (mg *MarseilleGame) IsLegalMarseille(move *chess.Move) bool {

	return true
}

func (mg *MarseilleGame) GetStatus() string {
	status := mg.internal.Position().Status()

	switch status {
	case chess.Checkmate:
		return "CHECKMATE - Game Over"
	case chess.Stalemate:
		return "DRAW (Stalemate)"
	case chess.FivefoldRepetition:
		return "DRAW (Repetition)"
	}

	if mg.subTurn == SecondHalf {
		return "Waiting for second sub-move..."
	}

	return "Waiting for player move"
}
