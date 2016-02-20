package evallist

import (
	b "github.com/zaramme/gogo-shogi/bitboard"
	c "github.com/zaramme/gogo-shogi/color"
	"github.com/zaramme/gogo-shogi/evaluate"
	h "github.com/zaramme/gogo-shogi/hand"
	p "github.com/zaramme/gogo-shogi/piece"
	sq "github.com/zaramme/gogo-shogi/square"
)

const ListSize = 38

var handB = h.Hand(c.Black)
var handW = h.Hand(c.White)

var HandPieceSToSquareHand [c.ColorNum][p.HandPieceNum]sq.Square

var HandPieceToSquareHand = [c.ColorNum][p.HandPieceNum]sq.Square{
	[p.HandPieceNum]sq.Square{
		sq.W_hand_pawn, sq.W_hand_lance, sq.W_hand_knight, sq.W_hand_silver,
		sq.W_hand_gold, sq.W_hand_bishop, sq.W_hand_rook},
	[p.HandPieceNum]sq.Square{
		sq.B_hand_pawn, sq.B_hand_lance, sq.B_hand_knight, sq.B_hand_silver,
		sq.B_hand_gold, sq.B_hand_bishop, sq.B_hand_rook}}

type EvalList struct {
	List0            [ListSize]int
	List1            [ListSize]int
	ListToSquareHand [ListSize]sq.Square
	SquareHandToList [sq.SquareHandNum]int
}

type positionInterFace interface {
	getHand(color c.Color) h.Hand
	getPiece(sq.Square) p.Piece
	PieceTypeBB(p.PieceType) b.BitBoard
	OccupiedBB() b.BitBoard
}

func (ev *EvalList) Set(pos positionInterFace) {
	handB := pos.getHand(c.Black)
	handW := pos.getHand(c.White)

	nList := 0

	// fooは本来DEFINEでマクロ化されているため、この処理はAperyより低速
	foo := func(hand h.Hand, hp p.HandPiece, list0_index, list1_index int, color c.Color) {
		max := int(hand.NumOf(hp))
		for i := 1; i <= max; i++ {
			ev.List0[nList] = list0_index + i
			ev.List1[nList] = list1_index + i
			squareHand := HandPieceSToSquareHand[color][hp] + sq.Square(i)
			ev.ListToSquareHand[nList] = squareHand
			ev.SquareHandToList[squareHand] = nList
			nList++
		}
	}

	foo(handB, p.HPawn, evaluate.F_HandPawn, evaluate.E_HandPawn, c.Black)
	foo(handW, p.HPawn, evaluate.E_HandBishop, evaluate.F_HandPawn, c.White)
	foo(handB, p.HLance, evaluate.F_HandLance, evaluate.E_HandLance, c.Black)
	foo(handW, p.HLance, evaluate.E_HandLance, evaluate.F_HandLance, c.White)
	foo(handB, p.HKnight, evaluate.F_HandKnight, evaluate.E_HandKnight, c.Black)
	foo(handW, p.HKnight, evaluate.E_HandKnight, evaluate.F_HandKnight, c.White)
	foo(handB, p.HSilver, evaluate.F_HandSilver, evaluate.E_HandSilver, c.Black)
	foo(handW, p.HSilver, evaluate.E_HandSilver, evaluate.F_HandSilver, c.White)
	foo(handB, p.HGold, evaluate.F_HandGold, evaluate.E_HandGold, c.Black)
	foo(handW, p.HGold, evaluate.E_HandGold, evaluate.F_HandGold, c.White)
	foo(handB, p.HBishop, evaluate.F_HandBishop, evaluate.E_HandBishop, c.Black)
	foo(handW, p.HBishop, evaluate.E_HandBishop, evaluate.F_HandBishop, c.White)
	foo(handB, p.HRook, evaluate.F_HandRook, evaluate.E_HandRook, c.Black)
	foo(handW, p.HRook, evaluate.E_HandRook, evaluate.F_HandRook, c.White)

	bb := pos.PieceTypeBB(p.King).NotThisAndAssign(pos.OccupiedBB())
	for bb.IsNot0() {
		square := bb.FirstOneFromI9()
		pc := pos.getPiece(square)
		ev.ListToSquareHand[nList] = square
		ev.SquareHandToList[square] = nList
		ev.List0[nList] = evaluate.KPPArray[pc] + int(square)
		nList++
		ev.List1[nList] = evaluate.KPPArray[pc.Inverse()] + int(square.Inverse())

	}
}
