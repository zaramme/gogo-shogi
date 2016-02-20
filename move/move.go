package move

import (
	p "github.com/zaramme/gogo-shogi/piece"
	s "github.com/zaramme/gogo-shogi/square"
)

// xxxxxxxx xxxxxxxx xxxxxxxx x1111111  移動先
// xxxxxxxx xxxxxxxx xx111111 1xxxxxxx  移動元。駒打ちの際には、PieceType + SquareNum - 1
// xxxxxxxx xxxxxxxx x1xxxxxx xxxxxxxx  1 なら成り
// xxxxxxxx xxxx1111 xxxxxxxx xxxxxxxx  移動する駒の PieceType 駒打ちの際には使用しない。
// xxxxxxxx 1111xxxx xxxxxxxx xxxxxxxx  取られた駒の PieceType

type Move uint32

const (
	PromoteFlag Move = 1 << 14
	MoveNone    Move = 0
	MoveNull    Move = 129
	MovePVsEnd       = 1 << 15 // for learn
)

type positionInterFace interface {
	GetPiece(square s.Square) p.Piece
}

// C++ constructors
// TODO:必要かどうかを検討する
// explicit Move(const u32 u) : value_(u) {}
// Move& operator = (const Move& m) { value_ = m.value_; return *this; }
// Move& operator = (const volatile Move& m) { value_ = m.value_; return *this; }
// // volatile Move& 型の *this を返すとなぜか警告が出るので、const Move& 型の m を返すことにする。
// const Move& operator = (const Move& m) volatile { value_ = m.value_; return m; }
// Move(const Move& m) { value_ = m.value_; }
// Move(const volatile Move& m) { value_ = m.value_; }

func (m Move) To() s.Square {
	v := (m >> 0) & 0x7f
	return s.Square(v)
}

func (m Move) From() s.Square {
	v := (m >> 7) & 0x7f
	return s.Square(v)
}

func (m Move) FromAndTo() uint32 {
	v := (m >> 0) & 0x3fff
	return uint32(v)
}

func (m Move) ProFromAndTo() uint32 {
	v := m >> 0 & 0x7fff
	return uint32(v)
}

func (m Move) Cap() p.PieceType {
	v := m >> 20 & 0xf
	return p.PieceType(v)
}

func (m Move) IsPromotion() uint32 {
	v := m & PromoteFlag
	return uint32(v)
}

func (m Move) PieceTypeFrom() p.PieceType {
	v := m >> 16
	return p.PieceType(v)
}

func (m Move) PieceTypeTo() p.PieceType {
	v := m >> 20
	return p.PieceType(v)
}

// 高速化
// 条件：PtFromが確定している時。
func (m Move) PieceTypeToFast(ptFrom p.PieceType) p.PieceType {
	v := ptFrom + p.PieceType((m&PromoteFlag)>>11)
	return v
}

func (m Move) IsDrop() bool {
	v := m.From() >> 81
	return v != 0
}

func (m Move) isCapture() bool {
	v := m & 0xf00000
	return v != 0
}

func (m Move) isCaputureOrPromotion() bool {
	v := m & 0xf04000
	return v != 0
}

func (m Move) isCaptureOrPawnPromotion() bool {
	if m.isCapture() {
		return true
	}
	return m.IsPromotion() != 0 && (m.PieceTypeFrom() == p.Pawn)
}

func (m Move) PieceTypeDropped() p.PieceType {
	v := m.From() - s.SquareNum + 1
	return p.PieceType(v)
}

func (m Move) PieceTypeFromOrDropped() p.PieceType {
	if m.IsDrop() {
		return m.PieceTypeDropped()
	}
	return m.PieceTypeFrom()
}

func (m Move) handPieceDropped() p.HandPiece {
	// assert(m.IsDrop())
	return p.HandPiece(m.PieceTypeDropped())
}

func (m Move) isNone() bool {
	return m == MoveNone
}

func (m Move) PromoteFlagToStringUSI() string {
	if m.IsPromotion() != 0 {
		return "+"
	}
	return ""
}

//	std::string toUSI() const;
//	std::string toCSA() const;

func NewMoveFromToSquare(to s.Square) Move {
	return Move(to << 0)
}

func NewMoveFromFromSquare(from s.Square) Move {
	return Move(from << 7)
}

// 駒打ちの駒の種類から移動元に変換
// todo: PieceType を HandPiece に変更
func Drop2From(pt p.PieceType) s.Square {
	v := s.SquareNum - 1 + s.Square(pt)
	return v
}

func NewMoveFromDropPieceType(pt p.PieceType) Move {
	v := Drop2From(pt)
	return Move(v)
}

func NewMoveFromPieceType(pt p.PieceType) Move {
	v := pt << 16
	return Move(v)
}

func From2Drop(from s.Square) p.PieceType {
	v := from - s.SquareNum + 1
	return p.PieceType(v)
}

func NewMoveFromCapturedPieceType(captured p.PieceType) Move {
	v := captured << 20
	return Move(v)
}

// 移動先と、Position から 取った駒の種類を判別し、指し手に変換
// 駒を取らないときは、0 (MoveNone) を返す。
// C++ : capturedPieceType2Move
func NewMoveFromCapturedPosition(to s.Square, pos positionInterFace) Move {
	captured := p.PieceType(pos.GetPiece(to))
	return NewMoveFromCapturedPieceType(captured)
}

// 移動元、移動先、移動する駒の種類から指し手に変換
//inline前提なので、関数呼び出しを無くした方が早いかも？
func NewMove(pt p.PieceType, from, to s.Square) Move {
	return NewMoveFromPieceType(pt) | NewMoveFromFromSquare(from) | NewMoveFromToSquare(to)
}

// 取った駒を判別する必要がある。
// この関数は駒を取らないときにも使える。
func NewCaptureMove(pt p.PieceType, from, to s.Square, pos positionInterFace) Move {
	return NewMoveFromCapturedPosition(to, pos) | NewMove(pt, from, to)
}

func NewCapturePromoteMove(pt p.PieceType, from, to s.Square, pos positionInterFace) Move {
	return NewCaptureMove(pt, from, to, pos) | PromoteFlag
}

func NewDropMove(pt p.PieceType, to s.Square) Move {
	return NewMoveFromToSquare(Drop2From(pt)) | NewMoveFromToSquare(to)
}

type MoveStack struct {
	Move  Move
	Score int
}

type MoveStacks []MoveStack

func (mt MoveStack) GreaterThan(mt2 MoveStack) bool {
	return mt.Score > mt2.Score
}

func (mt MoveStack) LessThan(mt2 MoveStack) bool {
	return mt.Score < mt2.Score
}

//template <typename T, bool UseSentinel = false> inline void insertionSort(T first, T last) {

func (mts MoveStacks) PickBest() MoveStack {
	max := MoveStack{0, 0}
	for _, mt := range mts {
		if mt.GreaterThan(max) {
			max = mt
		}
	}
	return max
}

func (m Move) move16ToMove(pos positionInterFace) Move {
	if m.isNone() {
		return MoveNone
	}
	if m.IsDrop() {
		return m
	}
	from := m.From()
	ptFrom := p.PieceType(pos.GetPiece(from))
	return m | NewMoveFromPieceType(ptFrom) | NewMoveFromCapturedPosition(m.To(), pos)
}
