package bitboard

type Bitboard uint64

type Board struct {
	Pieces   [12]Bitboard
	White    Bitboard
	Black    Bitboard
	Occupied Bitboard
}

func (b *Board) GetPieceAt(sq int) int {
	for i, bb := range b.Pieces {
		if (bb>>sq)&1 != 0 {
			return i
		}
	}
	return -1
}

func (b *Board) SetPieceAt(sq int, piece int, isWhite bool) {
	mask := Bitboard(1) << sq
	b.Pieces[piece] |= mask
	b.Occupied |= mask

	if isWhite {
		b.White |= mask
	} else {
		b.Black |= mask
	}
}

func (b *Board) ClearPieceAt(sq int, piece int, isWhite bool) {
	mask := ^(Bitboard(1) << sq)
	b.Pieces[piece] &= mask
	b.Occupied &= mask

	if isWhite {
		b.White &= mask
	} else {
		b.Black &= mask
	}
}
