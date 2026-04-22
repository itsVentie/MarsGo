package engine

import "github.com/notnil/chess"

func (mg *MarseilleGame) ValidMovesStr() []string {
	moves := mg.internal.ValidMoves()
	var res []string
	for _, m := range moves {
		res = append(res, m.String())
	}
	return res
}

func (mg *MarseilleGame) Clone() *MarseilleGame {
	clone := &MarseilleGame{
		internal: mg.internal.Clone(),
		subTurn:  mg.subTurn,
		history:  make([]MarseilleMove, len(mg.history)),
	}
	copy(clone.history, mg.history)
	return clone
}

func (mg *MarseilleGame) Position() *chess.Position {
	return mg.internal.Position()
}
