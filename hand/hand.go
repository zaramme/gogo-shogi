package hand

import (
	p "github.com/zaramme/gogo-shogi/piece"
)

// 手駒
// 手駒の状態 (32bit に pack する)
// 手駒の優劣判定を高速に行う為に各駒の間を1bit空ける。
// xxxxxxxx xxxxxxxx xxxxxxxx xxx11111  Pawn
// xxxxxxxx xxxxxxxx xxxxxxx1 11xxxxxx  Lance
// xxxxxxxx xxxxxxxx xxx111xx xxxxxxxx  Knight
// xxxxxxxx xxxxxxx1 11xxxxxx xxxxxxxx  Silver
// xxxxxxxx xxx111xx xxxxxxxx xxxxxxxx  Gold
// xxxxxxxx 11xxxxxx xxxxxxxx xxxxxxxx  Bishop
// xxxxx11x xxxxxxxx xxxxxxxx xxxxxxxx  Rook

//手駒の状態
type Hand uint32

// 特定の種類の持ち駒を返す
func (h Hand) NumOf(hp p.HandPiece) uint32 {
	return (uint32(h) & uint32(handPieceMask[hp])) >> handPieceShiftBits[hp]
}

// 特定の種類の持ち駒が存在しているかどうか
// 存在している場合はその種類のみのフィルタを返す
func (h Hand) Exists(hp p.HandPiece) (exists bool, value Hand) {
	value = h & Hand(handPieceMask[hp])
	exists = value != 0
	return exists, value
}

// 歩以外の持ち駒が存在しているかどうか
// 存在している場合はその歩以外のフィルタを返す
func (h Hand) ExceptPawnExists() (exists bool, value Hand) {
	value = h & Hand(handPieceExceptPawnMask)
	exists = value != 0
	return exists, value
}

// 特定の種類の持ち駒を特定値以上に
func (h Hand) OrEqual(num uint32, hp p.HandPiece) Hand {
	return h | Hand(num)<<handPieceShiftBits[hp]
}

//　持ち駒を１加算
func (h Hand) PlusOne(hp p.HandPiece) Hand {
	return h + Hand(handPieceOne[hp])
}

// 持ち駒を１減算
func (h Hand) MinusOne(hp p.HandPiece) Hand {
	return h - Hand(handPieceOne[hp])
}

// 手駒の優劣判定
// 手駒が ref と同じか、勝っていれば true
// 勝っている状態とは、全ての種類の手駒が、ref 以上の枚数があることを言う。
func (h Hand) IsEqualOrSuperior(ref Hand) bool {
	// こちらは、同じ意味でより高速
	// ref の方がどれか一つでも多くの枚数の駒を持っていれば、Borrow の位置のビットが立つ。
	return (h-ref)&Hand(borrowMask) == 0
}

const (
	hPawnShiftBits   uint32 = 0
	hLanceShiftBits  uint32 = 6
	hKnightShiftBits uint32 = 10
	hSilverShiftBits uint32 = 14
	hGoldShiftBits   uint32 = 18
	hBishopShiftBits uint32 = 22
	hRookShiftBits   uint32 = 25
)

const (
	hPawnMask               uint32 = 0x1f << hPawnShiftBits
	hLanceMask              uint32 = 0x7 << hLanceShiftBits
	hKnightMask             uint32 = 0x7 << hKnightShiftBits
	hSilverMask             uint32 = 0x7 << hSilverShiftBits
	hGoldMask               uint32 = 0x7 << hGoldShiftBits
	hBishopMask             uint32 = 0x3 << hBishopShiftBits
	hRookMask               uint32 = 0x3 << hRookShiftBits
	handPieceExceptPawnMask uint32 = (hLanceMask | hKnightMask |
		hSilverMask | hGoldMask |
		hBishopMask | hRookMask)
	borrowMask uint32 = ((hPawnMask + (1 << hPawnShiftBits)) |
		(hLanceMask + (1 << hLanceShiftBits)) |
		(hKnightMask + (1 << hKnightShiftBits)) |
		(hSilverMask + (1 << hSilverShiftBits)) |
		(hGoldMask + (1 << hGoldShiftBits)) |
		(hBishopMask + (1 << hBishopShiftBits)) |
		(hRookMask + (1 << hRookShiftBits)))
)

// 特定の駒を先頭に持ってくるシフトビット数
var handPieceShiftBits = [p.HandPieceNum]uint32{
	hPawnShiftBits, hLanceShiftBits, hKnightShiftBits, hSilverShiftBits,
	hGoldShiftBits, hBishopShiftBits, hRookShiftBits}

// 特定の種類の持ち駒を取り出すフィルタ
var handPieceMask = [p.HandPieceNum]uint32{
	hPawnMask, hLanceMask, hKnightMask, hSilverMask,
	hGoldMask, hBishopMask, hRookMask}

// 特定の種類の持ち駒を１コマ分増減するためのフィルタ
var handPieceOne = [p.HandPieceNum]uint32{
	(1 << hPawnShiftBits), (1 << hLanceShiftBits),
	(1 << hKnightShiftBits), (1 << hSilverShiftBits),
	(1 << hGoldShiftBits), (1 << hBishopShiftBits),
	(1 << hRookShiftBits)}
