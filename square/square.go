package square

import (
	c "github.com/zaramme/gogo-shogi/color"
)

type Square int

const (
	I9            Square = iota
	I8            Square = iota
	I7            Square = iota
	I6            Square = iota
	I5            Square = iota
	I4            Square = iota
	I3            Square = iota
	I2            Square = iota
	I1            Square = iota
	H9            Square = iota
	H8            Square = iota
	H7            Square = iota
	H6            Square = iota
	H5            Square = iota
	H4            Square = iota
	H3            Square = iota
	H2            Square = iota
	H1            Square = iota
	G9            Square = iota
	G8            Square = iota
	G7            Square = iota
	G6            Square = iota
	G5            Square = iota
	G4            Square = iota
	G3            Square = iota
	G2            Square = iota
	G1            Square = iota
	F9            Square = iota
	F8            Square = iota
	F7            Square = iota
	F6            Square = iota
	F5            Square = iota
	F4            Square = iota
	F3            Square = iota
	F2            Square = iota
	F1            Square = iota
	E9            Square = iota
	E8            Square = iota
	E7            Square = iota
	E6            Square = iota
	E5            Square = iota
	E4            Square = iota
	E3            Square = iota
	E2            Square = iota
	E1            Square = iota
	D9            Square = iota
	D8            Square = iota
	D7            Square = iota
	D6            Square = iota
	D5            Square = iota
	D4            Square = iota
	D3            Square = iota
	D2            Square = iota
	D1            Square = iota
	C9            Square = iota
	C8            Square = iota
	C7            Square = iota
	C6            Square = iota
	C5            Square = iota
	C4            Square = iota
	C3            Square = iota
	C2            Square = iota
	C1            Square = iota
	B9            Square = iota
	B8            Square = iota
	B7            Square = iota
	B6            Square = iota
	B5            Square = iota
	B4            Square = iota
	B3            Square = iota
	B2            Square = iota
	B1            Square = iota
	A9            Square = iota
	A8            Square = iota
	A7            Square = iota
	A6            Square = iota
	A5            Square = iota
	A4            Square = iota
	A3            Square = iota
	A2            Square = iota
	A1            Square = iota
	SquareNum     Square = 81 // = 81 as sentinel
	B_hand_pawn   Square = 81
	B_hand_lance  Square = B_hand_pawn + 18
	B_hand_knight Square = B_hand_lance + 4
	B_hand_silver Square = B_hand_knight + 4
	B_hand_gold   Square = B_hand_silver + 4
	B_hand_bishop Square = B_hand_gold + 4
	B_hand_rook   Square = B_hand_bishop + 2
	W_hand_pawn   Square = B_hand_rook + 2
	W_hand_lance  Square = W_hand_pawn + 18
	W_hand_knight Square = W_hand_lance + 4
	W_hand_silver Square = W_hand_knight + 4
	W_hand_gold   Square = W_hand_silver + 4
	W_hand_bishop Square = W_hand_gold + 4
	W_hand_rook   Square = W_hand_bishop + 2
	SquareHandNum Square = W_hand_rook + 3 // as sentinel
)

type File int

const (
	FileI   File = 0
	FileH   File = 1
	FileG   File = 2
	FileF   File = 3
	FileE   File = 4
	FileD   File = 5
	FileC   File = 6
	FileB   File = 7
	FileA   File = 8
	FileNum File = 9 // as sentinel
)

func (f File) isLeftOf(target File, turn c.Color) bool {
	if turn == c.Black {
		return f < target
	}

	return target < f
}

func (f File) isRightOf(target File, turn c.Color) bool {
	b := f.isLeftOf(target, turn)
	return !b
}

func (f File) isInFile() bool {
	return f < FileNum
}

func (f File) isInRank() bool {
	return f < FileNum
}

type Rank int

const (
	Rank9   Rank = iota
	Rank8   Rank = iota
	Rank7   Rank = iota
	Rank6   Rank = iota
	Rank5   Rank = iota
	Rank4   Rank = iota
	Rank3   Rank = iota
	Rank2   Rank = iota
	Rank1   Rank = iota
	RankNum      // as sentinel
)

func (r Rank) isInFrontOf(target Rank, turn c.Color) bool {
	if turn == c.Black {
		return target > r
	}
	return target < r
}

func (r Rank) isBehind(target Rank, turn c.Color) bool {
	f := r.isInFrontOf(target, turn)
	return !f
}

func (s Square) isInSquare() bool {
	return s < SquareNum
}

// 速度が必要な場面で使用するなら、テーブル引きの方が有効だと思う。
func MakeSquare(f File, r Rank) Square {
	return Square(int(f)*9 + int(r))
}

type SquareDelta int

const (
	DeltaNothing SquareDelta = 0
	DeltaN       SquareDelta = -1
	DeltaE       SquareDelta = -9
	DeltaS       SquareDelta = 1
	DeltaW       SquareDelta = 9
	DeltaNE      SquareDelta = DeltaN + DeltaE
	DeltaSE      SquareDelta = DeltaS + DeltaE
	DeltaSW      SquareDelta = DeltaS + DeltaW
	DeltaNW      SquareDelta = DeltaN + DeltaW
)

var squareToRank = [81]Rank{
	Rank9, Rank8, Rank7, Rank6, Rank5, Rank4, Rank3, Rank2, Rank1,
	Rank9, Rank8, Rank7, Rank6, Rank5, Rank4, Rank3, Rank2, Rank1,
	Rank9, Rank8, Rank7, Rank6, Rank5, Rank4, Rank3, Rank2, Rank1,
	Rank9, Rank8, Rank7, Rank6, Rank5, Rank4, Rank3, Rank2, Rank1,
	Rank9, Rank8, Rank7, Rank6, Rank5, Rank4, Rank3, Rank2, Rank1,
	Rank9, Rank8, Rank7, Rank6, Rank5, Rank4, Rank3, Rank2, Rank1,
	Rank9, Rank8, Rank7, Rank6, Rank5, Rank4, Rank3, Rank2, Rank1,
	Rank9, Rank8, Rank7, Rank6, Rank5, Rank4, Rank3, Rank2, Rank1,
	Rank9, Rank8, Rank7, Rank6, Rank5, Rank4, Rank3, Rank2, Rank1}

var squareToFile = [81]File{
	FileI, FileI, FileI, FileI, FileI, FileI, FileI, FileI, FileI,
	FileH, FileH, FileH, FileH, FileH, FileH, FileH, FileH, FileH,
	FileG, FileG, FileG, FileG, FileG, FileG, FileG, FileG, FileG,
	FileF, FileF, FileF, FileF, FileF, FileF, FileF, FileF, FileF,
	FileE, FileE, FileE, FileE, FileE, FileE, FileE, FileE, FileE,
	FileD, FileD, FileD, FileD, FileD, FileD, FileD, FileD, FileD,
	FileC, FileC, FileC, FileC, FileC, FileC, FileC, FileC, FileC,
	FileB, FileB, FileB, FileB, FileB, FileB, FileB, FileB, FileB,
	FileA, FileA, FileA, FileA, FileA, FileA, FileA, FileA, FileA}

func SquareToRank(s Square) Rank {
	return squareToRank[s]
}

func SquareToFile(s Square) File {
	return squareToFile[s]
}

// C++では速度面で別メソットだが、とりあえず内部処理を共通に
func MakeRank(s Square) Rank {
	return SquareToRank(s)
}

func MakeFile(s Square) File {
	return SquareToFile(s)
}

type Direction int

// 位置関係、方向
// ボナンザそのまま
// でもあまり使わないので普通の enum と同様に 0 から順に値を付けて行けば良いと思う。
const (
	DirecMisc     Direction = 0   // 縦、横、斜めの位置に無い場合
	DirecFile     Direction = 10  // 縦
	DirecRank     Direction = 11  // 横
	DirecDiagNESW Direction = 100 // 右上から左下
	DirecDiagNWSE Direction = 101 // 左上から右下
	DirecCross    Direction = 10  // 縦、横
	DirecDiag     Direction = 100 // 斜め
)

// 2つの位置関係のテーブル
var SquareRelation [SquareNum][SquareNum]Direction

// // 何かの駒で一手で行ける位置関係についての距離のテーブル。桂馬の位置は距離1とする。
var SquareDistance [SquareNum][SquareNum]int

// from, to, ksq が 縦横斜めの同一ライン上にあれば true を返す。
func (own Square) IsAligned(to Square, ksq Square) bool {
	direc := SquareRelation[own][ksq]
	if direc == DirecMisc {
		return false
	}
	return direc == SquareRelation[own][to]
}

// IsAlignedの高速版。
// to, ksqにdirecMusicが含まれない前提
func (own Square) IsAlignedFast(to Square, ksq Square) bool {
	direc := SquareRelation[own][ksq]
	return direc == SquareRelation[own][to]
}

// template <bool FROM_KSQ_NEVER_BE_DIRECMISC>
// inline bool isAligned(const Square from, const Square to, const Square ksq) {
// 	const Direction direc = squareRelation(from, ksq);
// 	if (FROM_KSQ_NEVER_BE_DIRECMISC) {
// 		assert(direc != DirecMisc);
// 		return (direc == squareRelation(from, to));
// 	}
// 	else {
// 		return (direc != DirecMisc && direc == squareRelation(from, to));
// 	}
// }

// inline char fileToCharUSI(const File f) { return '1' + f; }
// // todo: アルファベットが辞書順に並んでいない処理系があるなら対応すること。
// inline char rankToCharUSI(const Rank r) {
// 	static_assert('a' + 1 == 'b', "");
// 	static_assert('a' + 2 == 'c', "");
// 	static_assert('a' + 3 == 'd', "");
// 	static_assert('a' + 4 == 'e', "");
// 	static_assert('a' + 5 == 'f', "");
// 	static_assert('a' + 6 == 'g', "");
// 	static_assert('a' + 7 == 'h', "");
// 	static_assert('a' + 8 == 'i', "");
// 	return 'a' + r;
// }
// inline std::string squareToStringUSI(const Square sq) {
// 	const Rank r = makeRank(sq);
// 	const File f = makeFile(sq);
// 	const char ch[] = {fileToCharUSI(f), rankToCharUSI(r), '\0'};
// 	return std::string(ch);
// }

// inline char fileToCharCSA(const File f) { return '1' + f; }
// inline char rankToCharCSA(const Rank r) { return '1' + r; }
// inline std::string squareToStringCSA(const Square sq) {
// 	const Rank r = makeRank(sq);
// 	const File f = makeFile(sq);
// 	const char ch[] = {fileToCharCSA(f), rankToCharCSA(r), '\0'};
// 	return std::string(ch);
// }

// inline File charCSAToFile(const char c) { return static_cast<File>(c - '1'); }
// inline Rank charCSAToRank(const char c) { return static_cast<Rank>(c - '1'); }
// inline File charUSIToFile(const char c) { return static_cast<File>(c - '1'); }
// inline Rank charUSIToRank(const char c) { return static_cast<Rank>(c - 'a'); }

// 後手の位置を先手の位置へ変換
func (s Square) Inverse() Square {
	return SquareNum - 1 - s
}

// 左右変換
func (f File) Inverse() File {
	return FileNum - 1 - f
}

// 上下変換
func (r Rank) Inverse() Rank {
	return RankNum - 1 - r
}

// Square の左右だけ変換
func (s Square) InverseFile() Square {
	return MakeSquare(MakeFile(s).Inverse(), MakeRank(s).Inverse())
}

func (s Square) InverseIfWhite(color c.Color) Square {
	if color == c.Black {
		return s
	}
	return s.Inverse()
}

func (r Rank) CanPromote(color c.Color) bool {
	return 1 == (uint32(0x1c00007) & (uint32(1) << (uint32(color<<4) + uint32(r))))
}
