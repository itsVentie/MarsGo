package engine

import (
	"errors"
	"strings"

	"github.com/notnil/chess"
)

type MarseilleMove struct {
	Move    *chess.Move
	IsCheck bool
}

type MarseilleGame struct {
	internal *chess.Game
	subTurn  SubTurn
	history  []MarseilleMove
}

func NewMarseilleGame() *MarseilleGame {
	return &MarseilleGame{
		internal: chess.NewGame(),
		subTurn:  FirstHalf,
	}
}

func (mg *MarseilleGame) PushMove(mStr string) error {
	move, err := mg.parseMove(mStr)
	if err != nil {
		return err
	}

	targetPiece := mg.internal.Position().Board().Piece(move.S2())
	if targetPiece.Type() == chess.King {
		return errors.New("cannot capture the king; it's a checkmate situation")
	}

	moverColor := mg.internal.Position().Turn()

	if err := mg.internal.Move(move); err != nil {
		return err
	}

	isCheck := false

	if move.HasTag(chess.Check) || mg.internal.Position().Status() == chess.Checkmate {
		isCheck = true
	}

	if !isCheck {
		board := mg.internal.Position().Board()
		oppColor := chess.White
		if moverColor == chess.White {
			oppColor = chess.Black
		}

		var kingSq chess.Square
		foundKing := false
		for sq := chess.A1; sq <= chess.H8; sq++ {
			p := board.Piece(sq)
			if p.Type() == chess.King && p.Color() == oppColor {
				kingSq = sq
				foundKing = true
				break
			}
		}

		if foundKing {
			fenParts := strings.Split(mg.internal.Position().String(), " ")
			if fenParts[1] == "w" {
				fenParts[1] = "b"
			} else {
				fenParts[1] = "w"
			}
			tempFen, _ := chess.FEN(strings.Join(fenParts, " "))
			tempGame := chess.NewGame(tempFen)

			for _, m := range tempGame.ValidMoves() {
				if m.S2() == kingSq {
					isCheck = true
					break
				}
			}
		}
	}

	mg.history = append(mg.history, MarseilleMove{Move: move, IsCheck: isCheck})

	if isCheck || mg.subTurn == SecondHalf {
		mg.subTurn = FirstHalf
	} else {
		mg.subTurn = SecondHalf
		mg.forceSamePlayerTurn()
	}

	return nil
}

func (mg *MarseilleGame) Undo() error {
	if len(mg.history) == 0 {
		return errors.New("no moves to undo")
	}

	mg.history = mg.history[:len(mg.history)-1]

	newGame := chess.NewGame()
	mg.subTurn = FirstHalf

	for _, m := range mg.history {
		newGame.Move(m.Move)
		if m.IsCheck || mg.subTurn == SecondHalf {
			mg.subTurn = FirstHalf
		} else {
			mg.subTurn = SecondHalf
			fen := newGame.Position().String()
			parts := strings.Split(fen, " ")
			if parts[1] == "w" {
				parts[1] = "b"
			} else {
				parts[1] = "w"
			}
			newFen, _ := chess.FEN(strings.Join(parts, " "))
			newGame = chess.NewGame(newFen)
		}
	}
	mg.internal = newGame
	return nil
}

func (mg *MarseilleGame) parseMove(mStr string) (*chess.Move, error) {
	for _, n := range []chess.Notation{chess.UCINotation{}, chess.AlgebraicNotation{}} {
		if m, err := n.Decode(mg.internal.Position(), mStr); err == nil {
			return m, nil
		}
	}
	return nil, errors.New("invalid move format")
}

func (mg *MarseilleGame) forceSamePlayerTurn() {
	fen := mg.internal.Position().String()
	parts := strings.Split(fen, " ")

	if len(parts) > 1 {
		if parts[1] == "w" {
			parts[1] = "b"
		} else {
			parts[1] = "w"
		}
	}

	newFen := strings.Join(parts, " ")
	fenFunc, _ := chess.FEN(newFen)
	mg.internal = chess.NewGame(fenFunc)
}

func (mg *MarseilleGame) Turn() chess.Color {
	return mg.internal.Position().Turn()
}

func (mg *MarseilleGame) SubTurn() SubTurn {
	return mg.subTurn
}

func (mg *MarseilleGame) FEN() string {
	return mg.internal.Position().String()
}

func (mg *MarseilleGame) DrawBoard() string {
	return mg.internal.Position().Board().Draw()
}

func (mg *MarseilleGame) GetValidMoves() []string {
	moves := mg.internal.ValidMoves()
	res := make([]string, len(moves))
	for i, m := range moves {
		res[i] = m.String()
	}
	return res
}

func (mg *MarseilleGame) History() []string {
	res := make([]string, len(mg.history))
	for i, m := range mg.history {
		res[i] = m.Move.String()
	}
	return res
}

func (mg *MarseilleGame) IsGameOver() bool {
	return mg.internal.Position().Status() != chess.NoMethod
}

func (mg *MarseilleGame) LoadFEN(fen string) error {
	fenFunc, err := chess.FEN(fen)
	if err != nil {
		return err
	}
	mg.internal = chess.NewGame(fenFunc)
	return nil
}
