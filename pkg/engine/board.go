package engine

import (
	"marsgo/internal/bitboard"

	"github.com/notnil/chess"
)

type Board struct {
	bb *bitboard.Board
}

func NewBoard(pos *chess.Position) *Board {
	b := &Board{bb: &bitboard.Board{}}
	b.importPosition(pos)
	return b
}

func (b *Board) importPosition(pos *chess.Position) {
	internalBoard := pos.Board()
	for sq := chess.A1; sq <= chess.H8; sq++ {
		p := internalBoard.Piece(sq)
		if p != chess.NoPiece {
			b.bb.SetPieceAt(int(sq), int(p.Type()), p.Color() == chess.White)
		}
	}
}

func (b *Board) Bitboard() *bitboard.Board {
	return b.bb
}

func (b *Board) IsSquareOccupied(sq int) bool {
	return (b.bb.Occupied>>sq)&1 != 0
}
