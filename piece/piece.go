package piece

import (
	c "github.com/zaramme/gogo-shogi/color"
)

type PieceType uint16

const (
	Promote      PieceType = 8 // 成り変換の加算値
	Occupied     PieceType = 0
	Pawn         PieceType = 1
	Lance        PieceType = 2
	Knight       PieceType = 3
	Silver       PieceType = 4
	Bishop       PieceType = 5
	Rook         PieceType = 6
	Gold         PieceType = 7
	King         PieceType = 8
	ProPawn      PieceType = 9
	ProLance     PieceType = 10
	ProKnight    PieceType = 11
	ProSilver    PieceType = 12
	Horse        PieceType = 13
	Dragon       PieceType = 14
	PieceTypeNum PieceType = 15 // as sentinel
)

type Piece uint16

const (
	Empty      Piece = 0
	Unpromoted Piece = 0
	Promoted   Piece = 8
	BPawn      Piece = 1
	BLance     Piece = 2
	BKnight    Piece = 3
	BSilver    Piece = 4
	BBishop    Piece = 5
	BRook      Piece = 6
	BGold      Piece = 7
	BKing      Piece = 8
	BProPawn   Piece = 9
	BProLance  Piece = 10
	BProKnight Piece = 11
	BProSilver Piece = 12
	BHorse     Piece = 13
	BDragon    Piece = 14
	WPawn      Piece = 17
	WLance     Piece = 18
	WKnight    Piece = 19
	WSilver    Piece = 20
	WBishop    Piece = 21
	WRook      Piece = 22
	WGold      Piece = 23
	WKing      Piece = 24
	WProPawn   Piece = 25
	WProLance  Piece = 26
	WProKnight Piece = 27
	WProSilver Piece = 28
	WHorse     Piece = 29
	WDragon    Piece = 30
	PieceNone  Piece = 31 // PieceNone = 31  これを 32 にした方が多重配列のときに有利か。
)

// 先手・後手を入れ替える
func (p Piece) Inverse() Piece {
	return p ^ 0x10
}

type HandPiece uint16

const (
	HPawn        HandPiece = 0
	HLance       HandPiece = 1
	HKnight      HandPiece = 2
	HSilver      HandPiece = 3
	HGold        HandPiece = 4
	HBishop      HandPiece = 5
	HRook        HandPiece = 6
	HandPieceNum HandPiece = 7 //as sentinel
)

func (p Piece) PieceType() PieceType {
	return PieceType(p & 15)
}

func (p Piece) Color() c.Color {
	return c.Color(p >> 4)
}

func NewPieceWithColorAndPieceType(c c.Color, pt PieceType) Piece {
	return Piece(uint32(c<<4) | uint32(pt))
}

var isSliderVal = 0x60646064

// 飛び駒(飛車角香車)かどうか
func (p Piece) IsSlider() bool {
	return (isSliderVal & (1 << p)) != 0
}

// 飛び駒(飛車角香車)かどうか
func (pt PieceType) IsSlider() bool {
	return (isSliderVal & (1 << pt)) != 0
}

var pieceTypeToHandPiece = [PieceTypeNum]HandPiece{
	HandPieceNum, HPawn, HLance, HKnight, HSilver, HBishop, HRook, HGold,
	HandPieceNum, HPawn, HLance, HKnight, HSilver, HBishop, HRook}

func (pt PieceType) ToHandPiece() HandPiece {
	return pieceTypeToHandPiece[pt]
}
