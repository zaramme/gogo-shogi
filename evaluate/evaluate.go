package evaluate

import (
	c "github.com/zaramme/gogo-shogi/color"
	p "github.com/zaramme/gogo-shogi/piece"
	p "github.com/zaramme/gogo-shogi/pieceScore"
	s "github.com/zaramme/gogo-shogi/square"
)

// 評価関数テーブルのオフセット。
// f_xxx が味方の駒、e_xxx が敵の駒
// Bonanza の影響で持ち駒 0 の場合のインデックスが存在するが、参照する事は無い。
// todo: 持ち駒 0 の位置を詰めてテーブルを少しでも小さくする。(キャッシュに少しは乗りやすい?)
const (
	F_HandPawn   = 0
	E_HandPawn   = F_HandPawn + 19
	F_HandLance  = E_HandPawn + 19
	E_HandLance  = F_HandLance + 5
	F_HandKnight = E_HandLance + 5
	E_HandKnight = F_HandKnight + 5
	F_HandSilver = E_HandKnight + 5
	E_HandSilver = F_HandSilver + 5
	F_HandGold   = E_HandSilver + 5
	E_HandGold   = F_HandGold + 5
	F_HandBishop = E_HandGold + 5
	E_HandBishop = F_HandBishop + 3
	F_HandRook   = E_HandBishop + 3
	E_HandRook   = F_HandRook + 3
	FE_HandEnd   = E_HandRook + 3
)
const (
	F_Pawn   = FE_HandEnd
	E_Pawn   = F_Pawn + 81
	F_Lance  = E_Pawn + 81
	E_Lance  = F_Lance + 81
	F_Knight = E_Lance + 81
	E_Knight = F_Knight + 81
	F_Silver = E_Knight + 81
	E_Silver = F_Silver + 81
	F_Gold   = E_Silver + 81
	E_Gold   = F_Gold + 81
	F_Bishop = E_Gold + 81
	E_Bishop = F_Bishop + 81
	F_Horse  = E_Bishop + 81
	E_Horse  = F_Horse + 81
	F_Rook   = E_Horse + 81
	E_Rook   = F_Rook + 81
	F_Dragon = E_Rook + 81
	E_Dragon = F_Dragon + 81
	FE_End   = E_Dragon + 8
)

const FVScale = 32

var KPPIndexArray = []int{
	F_HandPawn, E_HandPawn, F_HandLance, E_HandLance, F_HandKnight,
	E_HandKnight, F_HandSilver, E_HandSilver, F_HandGold, E_HandGold,
	F_HandBishop, E_HandBishop, F_HandRook, E_HandRook, /*fe_Handend,*/
	F_Pawn, E_Pawn, F_Lance, E_Lance, F_Knight, E_Knight, F_Silver, E_Silver,
	F_Gold, E_Gold, F_Bishop, E_Bishop, F_Horse, E_Horse, F_Rook, E_Rook,
	F_Dragon, E_Dragon, FE_End}

var KPPArray [31]int
var KKPArray [15]int

var KPPHandArray = [c.ColorNum][p.HandPieceNum]int{
	[p.HandPieceNum]int{F_HandPawn, F_HandLance, F_HandKnight, F_HandSilver,
		F_HandGold, F_HandBishop, F_HandRook},
	[p.HandPieceNum]int{E_HandPawn, E_HandLance, E_HandKnight, E_HandSilver,
		E_HandGold, E_HandBishop, E_HandRook}}

type positionInterface interface {
}

func KppIndexToSquare(i int) s.Square {
	// const auto it = std::upper_bound(std::begin(KPPIndexArray), std::end(KPPIndexArray), i);
	// return static_cast<Square>(i - *(it - 1));
	return s.B9
}
