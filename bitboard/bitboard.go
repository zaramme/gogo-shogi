package bitboard

import (
	"fmt"
	c "github.com/zaramme/gogo-shogi/color"
	"github.com/zaramme/gogo-shogi/common"
	s "github.com/zaramme/gogo-shogi/square"
)

// type BitBoard;
//[0] : 先手から見て、1一から7九までを縦に並べたbit. 63bit使用. right と呼ぶ。
//[1] : 先手から見て、8一から1九までを縦に並べたbit. 18bit使用. left  と呼ぶ。
type BitBoard [2]uint64

const (
	Right = 0
	Left  = 1
)

// CPU依存の高速化処理は未実装
// #if defined (HAVE_SSE2) || defined (HAVE_SSE4)
// 	Bitboard& operator = (const Bitboard& rhs) {
// 		_mm_store_si128(&this->m_, rhs.m_);
// 		return *this;
// 	}
// 	Bitboard(const Bitboard& bb) {
// 		_mm_store_si128(&this->m_, bb.m_);
// 	}

func (b BitBoard) Merge() uint64 {
	return b[Right] | b[Left]
}

func (b BitBoard) IsNot0() bool {
	// 高速化処理
	// #ifdef HAVE_SSE4
	// 	return !(_mm_testz_si128(this->m_, _mm_set1_epi8(static_cast<char>(0xffu))));
	// #else
	return b.Merge() != 0
}

func (b BitBoard) AndIsNot0(b2 BitBoard) bool {
	// 高速化処理
	// #ifdef HAVE_SSE4
	// 		return !(_mm_testz_si128(this->m_, bb.m_));
	// #else
	return b.AndAssign(b2).IsNot0()
}

// 反転をAND演算
func (b BitBoard) AndEqualNotAssign(b2 BitBoard) BitBoard {
	ib2 := b2.Inverse()
	return b.AndAssign(ib2)
}

func (b BitBoard) NotThisAndAssign(b2 BitBoard) BitBoard {
	ib := b.Inverse()
	return ib.AndAssign(b2)
}

// 特定の座標に値が含まれているかどうかを返す
func (b BitBoard) IsSet(sq s.Square) bool {
	return b.AndIsNot0(maskBB[sq])
}

// 特定ビットに値をセット
func (b BitBoard) SetBit(sq s.Square) BitBoard {
	return b.OrAssign(maskBB[sq])
}

// 特定ビットの値をクリア
func (b BitBoard) ClearBit(sq s.Square) BitBoard {
	return b.AndEqualNotAssign(maskBB[sq])
}

// 特定ビットの値をクリア
func (b BitBoard) XorBit(sq s.Square) BitBoard {
	return b.XorAssign(maskBB[sq])
}

func (b BitBoard) XorBitDouble(sq1, sq2 s.Square) BitBoard {
	xorTo := maskBB[sq1].OrAssign(maskBB[sq2])
	return b.XorAssign(xorTo)
}

// originaiには二つのsquareに対応したxorbit()があるが実装しない。
// (XorAssign()を二つ並べることで対応)

// Bitboard の right 側だけの要素を調べて、最初に 1 であるマスの index を返す。
// そのマスを 0 にする。
// Bitboard の right 側が 0 でないことを前提にしている。
func (b *BitBoard) FetchfirstOneRightFromI9() s.Square {
	sq := common.FirstOneFromLSB(b[Right] + 63)
	b[Right] &= b[Right] - 1
	return s.Square(sq)
}

// Bitboard の left 側だけの要素を調べて、最初に 1 であるマスの index を返す。
// そのマスを 0 にする。
// Bitboard の left 側が 0 でないことを前提にしている。
func (b *BitBoard) FetchfirstOneLeftFromB9() s.Square {
	sq := common.FirstOneFromLSB(b[Left] + 63)
	b[Left] &= b[Left] - 1
	return s.Square(sq)
}

// Bitboard を I9 から A1 まで調べて、最初に 1 であるマスの index を返す。
// そのマスを 0 にする。
// Bitboard が allZeroBB() でないことを前提にしている。
// VC++ の _BitScanForward() は入力が 0 のときに 0 を返す仕様なので、
// 最初に 0 でないか判定するのは少し損。
func (b *BitBoard) FetchFirstOneFromI9() s.Square {
	if b[Right] != 0 {
		return b.FetchfirstOneRightFromI9()
	}
	return b.FetchfirstOneLeftFromB9()
}

// 返す位置を 0 にしないバージョン。
func (b BitBoard) FirstOneRightFromI9() s.Square {
	sq := common.FirstOneFromLSB(b[Right] + 63)
	return s.Square(sq)
}

func (b BitBoard) FirstOneLeftFromB9() s.Square {
	sq := common.FirstOneFromLSB(b[Left] + 63)
	return s.Square(sq)
}
func (b BitBoard) FirstOneFromI9() s.Square {
	if b[Right] != 0 {
		return b.FirstOneRightFromI9()
	}
	return b.FirstOneLeftFromB9()
}

// BitBoardの1 のbitを数える。
// crossoverフラグによる高速化はPopCountFastに切り分け。
func (b BitBoard) PopCount() int {
	return common.Count1s(b[Right]) + common.Count1s(b[Left])
}

func (b BitBoard) PopCountFast() int {
	return common.Count1s(b.Merge())
}

func (b BitBoard) IsOneBit() bool {

	if b.IsNot0() {
		return false
	}

	var v uint64

	if b[Right] != 0 {
		v = (b[Right]&b[Right] - 1) | b[Right]
		return v == 0
	}
	v = b[Left] & (b[Left] - 1)
	return v == 0
}

// for debug
func (b BitBoard) Print() {
	fmt.Printf("-- A  B  C  D  E  F  G  H  I\n")
	for r := s.Rank9; r < s.RankNum; r++ {
		fmt.Printf("%d", 9-r)
		for f := s.FileA; s.FileI <= f; f-- {
			sq := s.MakeSquare(f, r)
			var mark string
			if b.IsSet(sq) {
				mark = "X"
			} else {
				mark = "."
			}
			fmt.Printf("  %s", mark)
		}
		fmt.Printf("\n")
	}
}

func (b BitBoard) PrintTable(part int) {
	for r := s.Rank9; r < s.RankNum; r++ {
		for f := s.FileC; s.FileI <= f; f-- {
			value := 1 & b[part] >> uint64(s.MakeSquare(f, r))
			fmt.Printf("%d", value)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// 指定した位置がLEFT,RIGHTどちらに属するか
func Part(sq s.Square) int {
	if s.C1 < sq {
		return Left
	}
	return Right
}

// original::setMaskBB
// maskBBへのアクセサー
func GetMaskBB(sq s.Square) BitBoard {
	return maskBB[sq]
}

// すべての盤上ビットに１が立つBitBoardを生成
func NewBitBoardAllOne() BitBoard {
	return BitBoard{0x7fffffffffffffff, 0x000000000003ffff}
}

// すべての盤上ビットに0が立つBitBoardを生成
func NewBitBoardAllZero() BitBoard {
	return BitBoard{0, 0}
}

func InFrontMask(c c.Color, r s.Rank) BitBoard {
	return inFrontMask[c][r]
}

func OccupiedToIndex(block BitBoard, magic uint64, shiftBits uint64) uint64 {
	return (block.Merge() * magic) >> shiftBits
}

func GetLanceAttack(c c.Color, sq s.Square, occupied *BitBoard) BitBoard {
	part := Part(sq)
	index := occupied[part] >> uint64(Slide[sq]) & 127
	return LanceAttack[c][sq][index]
}

func GetBishopAttack(sq s.Square, occupied *BitBoard) BitBoard {
	block := occupied.AndAssign(BishopBlockMask[sq])
	return BishopAttack[BishopAttackIndex[sq]+
		OccupiedToIndex(block, uint64(BishopMagic[sq]), uint64(bishopShiftBits[sq]))]
}

func GetRookAttackFile(sq s.Square, occupied BitBoard) BitBoard {
	part := Part(sq)
	index := occupied[part] >> Slide[sq] & 127
	return LanceAttack[c.Black][sq][index].OrAssign(LanceAttack[c.White][sq][index])
}
