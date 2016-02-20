package position

import (
	//	"github.com/zaramme/gogo-shogi/common"
	bb "github.com/zaramme/gogo-shogi/bitboard"
	c "github.com/zaramme/gogo-shogi/color"
	com "github.com/zaramme/gogo-shogi/common"
	evl "github.com/zaramme/gogo-shogi/evallist"
	ev "github.com/zaramme/gogo-shogi/evaluate"
	h "github.com/zaramme/gogo-shogi/hand"
	m "github.com/zaramme/gogo-shogi/move"
	p "github.com/zaramme/gogo-shogi/piece"
	ps "github.com/zaramme/gogo-shogi/pieceScore"
	sc "github.com/zaramme/gogo-shogi/score"
	sq "github.com/zaramme/gogo-shogi/square"
	"math/rand"
)

// #include "common.hpp"
// #include "hand.hpp"
// #include "bitboard.hpp"
// #include "pieceScore.hpp"
// #include "evalList.hpp"
// #include <stack>
// #include <memory>

type RepetitionType int

// position sfen R8/2K1S1SSk/4B4/9/9/9/9/9/1L1L1L3 b PLNSGBR17p3n3g 1
// の局面が最大合法手局面で 593 手。番兵の分、+ 1 しておく。
const MaxLegalMoves = 593 + 1

const (
	NotRepetition      RepetitionType = 0
	RepetitionDraw     RepetitionType = 1
	RepetitionWin      RepetitionType = 2
	RepetitionLose     RepetitionType = 3
	RepetitionSuperior RepetitionType = 4
	RepetitionInferior RepetitionType = 5
)

type Position struct {
	ByTypeBB  [p.PieceTypeNum]bb.BitBoard
	ByColorBB [c.ColorNum]bb.BitBoard

	// 小駒の成り駒を含めた金の位置
	GoldsBB bb.BitBoard

	// 各マスの状態
	Piece      [sq.SquareNum]p.Piece
	KingSquare [c.ColorNum]sq.Square

	// 手駒
	Hand [c.ColorNum]h.Hand
	Turn c.Color

	EvalList   evl.EvalList
	startState StateInfo
	State      *StateInfo //aperyでは"st"

	// 時間管理に使用する。
	gamePly    sc.Ply
	thisThread threadInterface
	Nodes      uint64

	Sarcher *searchInterface
}

// static Key zobrist_[PieceTypeNum][SquareNum][ColorNum];
var zobrist [p.PieceTypeNum][sq.SquareNum][c.ColorNum]com.Key

// static const Key zobTurn_ = 1;
var zobTurn = 1

// static Key zobHand_[HandPieceNum][ColorNum];
var zobHand [p.HandPieceNum][c.ColorNum]com.Key

// static Key zobExclusion_; // todo: これが必要か、要検討
var zobExclusion com.Key

func NewPositionWithSfen(sfen string, th threadInterface, srch searchInterface) {
	pos := new(Position)
	pos.Set(sfen, th)

}

////////////////////////////
// interfaceとしてのgetter(Moveで使用)
func (pos Position) GetPiece(square sq.Square) p.Piece {
	return pos.Piece[square]
}

// = 演算子のオペレータ
func (pos *Position) assign(pos2 *Position) {
	//後で書く
}

// 盤面をセットする
func (pos Position) Set(sfen string, th threadInterface) {
	pos.thisThread = th
}

// 以下は直接変数へアクセスする事で対応
// Bitboard bbOf(const PieceType pt) const{
//	return byTypeBB_[pt];
//}
// Bitboard bbOf(const Color c) const {
//		return byColorBB_[c];
//}

// 直接ByTypeBB[pt]にアクセスしても良いて、evallistへのインターフェースのため必要
func (pos Position) PieceTypeBB(pt p.PieceType) bb.BitBoard {
	return pos.ByTypeBB[pt]
}

// bbOf(const PieceType pt, const Color c)
func (pos Position) PieceTypeColorBB(pt p.PieceType, co c.Color) bb.BitBoard {
	return pos.ByTypeBB[pt].AndAssign(pos.ByColorBB[co])
}

// bbOf(const PieceType pt1, const PieceType pt2)
func (pos Position) DoublePieceTypeBB(pt1, pt2 p.PieceType) bb.BitBoard {
	return pos.ByTypeBB[pt1].OrAssign(pos.ByTypeBB[pt2])
}

// bbOf(const PieceType pt1, const PieceType pt2)
func (pos Position) DoublePieceTypeColorBB(pt1, pt2 p.PieceType, co c.Color) bb.BitBoard {
	return pos.DoublePieceTypeBB(pt1, pt2).AndAssign(pos.ByColorBB[co])
}

// bbOf(const PieceType pt1, const PieceType pt2, const PieceType pt3)
func (pos Position) TriplePieceTypeBB(pt1, pt2, pt3 p.PieceType) bb.BitBoard {
	return pos.DoublePieceTypeBB(pt1, pt2).AndAssign(pos.ByTypeBB[pt3])
}

func (pos Position) QuadPieceTypeBB(pt1, pt2, pt3, pt4 p.PieceType) bb.BitBoard {
	return pos.TriplePieceTypeBB(pt1, pt2, pt3).AndAssign(pos.ByTypeBB[pt4])
}

func (pos Position) QuintPieceTypeBB(pt1, pt2, pt3, pt4, pt5 p.PieceType) bb.BitBoard {
	return pos.QuadPieceTypeBB(pt1, pt2, pt3, pt4).AndAssign(pos.ByTypeBB[pt5])
}

func (pos Position) emptyBB() bb.BitBoard {
	return pos.ByTypeBB[p.Occupied].XorAssign(bb.NewBitBoardAllOne())
}

func (pos Position) goldsBBWithColor(co c.Color) bb.BitBoard {
	return pos.GoldsBB.AndAssign(pos.ByColorBB[co])
}

// Occupied情報へのエイリアス(ByTypeBB[p.Occupied])
func (pos Position) OccupiedBB() bb.BitBoard {
	return pos.ByTypeBB[p.Occupied]
}

// turn() 側が pin されている Bitboard を返す。
// checkersBB が更新されている必要がある。
func (pos Position) pinnedBB() bb.BitBoard {
	return pos.HiddenCheckers(true, true)
}

// turn() 側が
// pin されて(して)いる駒の Bitboard を返す。
// BetweenIsUs == true  : 間の駒が自駒。
// BetweenIsUs == false : 間の駒が敵駒。
func (pos *Position) XorBBs(pt p.PieceType, square sq.Square, color c.Color) {
	pos.ByTypeBB[p.Occupied] = pos.ByTypeBB[p.Occupied].XorBit(square)
	pos.ByTypeBB[pt] = pos.ByTypeBB[pt].XorBit(square)
	pos.ByColorBB[color] = pos.ByColorBB[color].XorBit(square)
}

func (pos Position) HiddenCheckers(findPinned, betweenIsUs bool) bb.BitBoard {
	result := bb.NewBitBoardAllZero()
	us := pos.Turn
	them := us.Opposite()

	var betweenColor c.Color
	if betweenIsUs {
		betweenColor = us
	} else {
		betweenColor = them
	}

	// pin する遠隔駒
	// まずは自駒か敵駒かで大雑把に判別
	var pinners bb.BitBoard
	var ksq sq.Square
	if findPinned {
		pinners = pos.ByColorBB[them]
		ksq = pos.KingSquare[us]
	} else {
		pinners = pos.ByColorBB[us]
		ksq = pos.KingSquare[them]
	}

	lanceToAppend := pos.ByTypeBB[p.Lance]
	if findPinned {
		lanceToAppend = lanceToAppend.AndAssign(bb.LanceAttackToEdge[us][ksq])
	} else {
		lanceToAppend = lanceToAppend.AndAssign(bb.LanceAttackToEdge[them][ksq])
	}
	rookToAppend := pos.ByTypeBB[p.Rook].AndAssign(bb.RookAttackToEdge[ksq])
	bishopToAppend := pos.ByTypeBB[p.Bishop].AndAssign(bb.BishopAttackToEdge[ksq])

	toAppend := lanceToAppend.OrAssign(rookToAppend).OrAssign(bishopToAppend)
	pinners = pinners.AndAssign(toAppend)

	if pinners.IsNot0() {
		square := pinners.FirstOneFromI9()
		between := bb.BetweenBB[square][ksq].AndAssign(pos.OccupiedBB())

		// pin する遠隔駒と玉の間にある駒が1つで、かつ、引数の色のとき、その駒は(を) pin されて(して)いる。
		if between.IsNot0() &&
			between.IsOneBit() &&
			between.AndIsNot0(pos.ByColorBB[betweenColor]) {
			result = result.OrAssign(between)
		}

	}

	return result

}

// 間の駒を動かすことで、turn() 側が空き王手が出来る駒のBitboardを返す。
// checkersBB が更新されている必要はない。
// BetweenIsUs == true  : 間の駒が自駒。
// BetweenIsUs == false : 間の駒が敵駒。
//template <bool BetweenIsUs = true> Bitboard discoveredCheckBB() const { return hiddenCheckers<false, BetweenIsUs>(); }
func (pos Position) DisCoverdCheckBB() bb.BitBoard {
	// あとでかく
	return bb.NewBitBoardAllZero()
}

// toFile と同じ筋に us の歩がないなら true
//return !bbOf(Pawn, us).andIsNot0(fileMask(toFile));
func (pos Position) noPawns(us c.Color, toFile sq.File) bool {
	pawnBB := pos.PieceTypeColorBB(p.Pawn, us)
	fileBB := bb.FileMask[toFile]

	return pawnBB.AndIsNot0(fileBB)
}

// 打ち歩詰め判定
//bool isPawnDropCheckMate(const Color us, const Square sq) const;

// Pinされているfromの駒がtoに移動出来なければtrueを返す。

//template <bool IsKnight = false>
func (pos Position) isPinnedIllegal(from, to, ksq sq.Square, pinned *bb.BitBoard) bool {
	// 後で書く
	// 桂馬ならどこに動いても駄目。
	// return pinned.isSet(from) && (IsKnight || !isAligned<true>(from, to, ksq));
	return false
}

// 空き王手かどうか。
// template <bool IsKnight = false>
func (pos Position) isDiscoverdCheck(from, to, ksq sq.Square, dcBB bb.BitBoard) bool {
	// 桂馬ならどこに動いても空き王手になる。
	// 	return dcBB.isSet(from) && (IsKnight || !isAligned<true>(from, to, ksq));
	return false
}

func (pos Position) CheckersBB() bb.BitBoard {
	return pos.State.CheckersBB
}

func (pos Position) prevCheckersBB() bb.BitBoard {
	return pos.State.Previous.CheckersBB
}

// 王手が掛かっているか。
//	bool inCheck() const            { return checkersBB().isNot0(); }
func (pos Position) IsCheck() bool {
	return pos.CheckersBB().IsNot0()
}

// func (pos Position) material()sc.Score{
//後で書く
// 	pos.State.mate
// }

//	Score materialDiff() const { return st_->material - st_->previous->material; }

// 単に直接pos.kingSquareへの参照で対応する
//func (pos Position) kingSquare(const Color c)sq.Square

//	bool moveGivesCheck(const Move m) const;
//	bool moveGivesCheck(const Move move, const CheckInfo& ci) const;

// bool moveGivesCheck(const Move m) const;
// bool moveGivesCheck(const Move move, const CheckInfo& ci) const;

//////////////////////////////////////////////////////////////
// Move

// template <bool MUSTNOTDROP, bool FROMMUSTNOTBEKING>
// bool pseudoLegalMoveIsLegal(const Move move, const Bitboard& pinned) const;
// bool pseudoLegalMoveIsEvasion(const Move move, const Bitboard& pinned) const;

// bool moveIsPseudoLegal(const Move move, const bool checkPawnDrop = false) const;
// //#if !defined NDEBUG
// bool moveIsLegal(const Move move) const;

// void doMove(const Move move, StateInfo& newSt);

// void doMove(const Move move, StateInfo& newSt, const CheckInfo& ci, const bool moveIsCheck);
func (pos *Position) Domove(mv m.Move, newSt *StateInfo, checkInfo *CheckInfo, moveIsCheck bool) {
	// assert(isOK());
	// assert(!move.isNone());
	// assert(&newSt != st_);

	boardKey := pos.State.boardKey
	boardKey ^= com.Key(zobTurn)
	handKey := pos.State.HandKey

	newSt.Previous = pos.State
	pos.State = newSt

	newSt.ChangedLists.Size = 1

	us := pos.Turn
	to := mv.To()
	ptCaptured := mv.Cap()

	if mv.IsDrop() {
		ptTo := mv.PieceTypeDropped()
		hpTo := ptTo.ToHandPiece()

		handKey -= zobHand[hpTo][us]
		boardKey += zobrist[ptTo][to][us]

		// メモリ関連の命令なのでスキップ
		//prefetch(csearcher()->tt.firstEntry(boardKey + handKey));

		handNum := int(h.Hand(us).NumOf(hpTo))
		listIndex := pos.EvalList.SquareHandToList[evl.HandPieceToSquareHand[us][hpTo]] + int(handNum)
		pcTo := p.NewPieceWithColorAndPieceType(us, ptTo)

		pos.State.ChangedLists.ListIndex[0] = listIndex
		pos.State.ChangedLists.ClistPair[0].oldList[0] = pos.EvalList.List0[listIndex]
		pos.State.ChangedLists.ClistPair[0].oldList[1] = pos.EvalList.List1[listIndex]

		pos.EvalList.List0[listIndex] = ev.KPPArray[pcTo] + int(to)
		pos.EvalList.List1[listIndex] = ev.KPPArray[pcTo.Inverse()] + int(to.Inverse())
		pos.EvalList.ListToSquareHand[listIndex] = to
		pos.EvalList.SquareHandToList[to] = listIndex

		pos.State.ChangedLists.ClistPair[0].newList[0] = pos.EvalList.List0[listIndex]
		pos.State.ChangedLists.ClistPair[0].newList[1] = pos.EvalList.List1[listIndex]

		pos.Hand[us] = pos.Hand[us].MinusOne(hpTo)
		pos.XorBBs(ptTo, to, us)
		pos.Piece[to] = p.NewPieceWithColorAndPieceType(us, ptTo)

		if moveIsCheck {
			// Direct checks
			pos.State.CheckersBB = bb.GetMaskBB(to)
			pos.State.ContinuousCheck[us] += 2
		} else {
			pos.State.CheckersBB = bb.NewBitBoardAllZero()
			pos.State.ContinuousCheck[us] = 0
		}
	} else {
		from := mv.From()
		ptFrom := mv.PieceTypeFrom()
		ptTo := mv.PieceTypeTo()

		pos.ByTypeBB[ptFrom] = pos.ByTypeBB[ptFrom].XorBit(from)
		pos.ByTypeBB[ptTo] = pos.ByTypeBB[ptTo].XorBit(to)
		pos.ByColorBB[us] = pos.ByColorBB[us].XorBitDouble(from, to)

		pos.Piece[from] = p.Empty
		pos.Piece[to] = p.NewPieceWithColorAndPieceType(us, ptTo)

		boardKey -= zobrist[ptFrom][from][us]
		boardKey += zobrist[ptTo][to][us]

		if int(ptCaptured) != 0 {
			//駒を取ったとき
			hpCaptured := ptCaptured.ToHandPiece()
			them := us.Opposite()

			boardKey -= zobrist[ptCaptured][to][them]
			handKey += zobHand[hpCaptured][us]

			pos.ByTypeBB[ptCaptured] = pos.ByTypeBB[ptCaptured].XorBit(to)
			pos.ByColorBB[them].XorBit(to)

			pos.Hand[us] = pos.Hand[us].PlusOne(hpCaptured)
			toListIndex := pos.EvalList.SquareHandToList[to]
			pos.State.ChangedLists.ListIndex[1] = toListIndex
			pos.State.ChangedLists.ClistPair[1].oldList[0] = pos.EvalList.List0[toListIndex]
			pos.State.ChangedLists.ClistPair[1].oldList[0] = pos.EvalList.List1[toListIndex]

			pos.State.ChangedLists.Size = 2

			handnum := pos.Hand[us].NumOf(hpCaptured)

			pos.EvalList.List0[toListIndex] = ev.KPPHandArray[us][hpCaptured] + int(handnum)
			pos.EvalList.List1[toListIndex] = ev.KPPHandArray[them][hpCaptured] + int(handnum)
			squareHand := evl.HandPieceSToSquareHand[us][hpCaptured] + sq.Square(handnum)
			pos.EvalList.ListToSquareHand[toListIndex] = squareHand
			pos.EvalList.SquareHandToList[squareHand] = toListIndex

			pos.State.ChangedLists.ClistPair[1].newList[0] = pos.EvalList.List0[toListIndex]
			pos.State.ChangedLists.ClistPair[1].newList[1] = pos.EvalList.List1[toListIndex]

			if us == c.Black {
				pos.State.material += ps.CapturedPieceScore[ptCaptured]
			} else {
				pos.State.material -= ps.CapturedPieceScore[ptCaptured]
			}
		}

		//prefetch(csearcher()->tt.firstEntry(boardKey + handKey));
		pos.ByTypeBB[p.Occupied] = pos.ByColorBB[c.Black].OrAssign(pos.ByColorBB[c.White])

		if ptTo == p.King {
			pos.KingSquare[c.Black] = to
		} else {
			pcTo := p.NewPieceWithColorAndPieceType(us, ptTo)
			fromListIndex := pos.EvalList.SquareHandToList[from]

			pos.State.ChangedLists.ListIndex[0] = fromListIndex
			pos.State.ChangedLists.ClistPair[0].oldList[0] = pos.EvalList.List0[fromListIndex]
			pos.State.ChangedLists.ClistPair[0].oldList[1] = pos.EvalList.List1[fromListIndex]

			pos.EvalList.List0[fromListIndex] = ev.KPPArray[pcTo] + int(to)
			pos.EvalList.List1[fromListIndex] = ev.KPPArray[pcTo.Inverse()] + int(to.Inverse())
			pos.EvalList.ListToSquareHand[fromListIndex] = to
			pos.EvalList.SquareHandToList[to] = fromListIndex

			pos.State.ChangedLists.ClistPair[0].newList[0] = pos.EvalList.List0[fromListIndex]
			pos.State.ChangedLists.ClistPair[0].newList[1] = pos.EvalList.List1[fromListIndex]
		}

		if mv.IsPromotion() != 0 {
			if us == c.Black {
				pos.State.material += ps.PieceScore[ptTo] - ps.PieceScore[ptFrom]
			}
			if us == c.White {
				pos.State.material -= ps.PieceScore[ptTo] - ps.PieceScore[ptFrom]
			}
		}

		if moveIsCheck {
			pos.State.CheckersBB = checkInfo.checkBB[ptTo].AndAssign(bb.GetMaskBB(to))

			ksq := pos.KingSquare[us.Opposite()]
			if pos.isDiscoverdCheck(from, to, ksq, checkInfo.dcBB) {
				switch sq.SquareRelation[from][ksq] { //(squareRelation(from, ksq))
				case sq.DirecFile:
					// from の位置から縦に利きを調べると相手玉と、空き王手している駒に当たっているはず。味方の駒が空き王手している駒。
					appendBB := bb.GetRookAttackFile(from, pos.OccupiedBB()).AndAssign(pos.ByColorBB[us])
					pos.State.CheckersBB = pos.State.CheckersBB.OrAssign(appendBB)
				case sq.DirecRank:
					appendBB := pos.AttacksFrom(p.Rook, c.ColorNum, ksq).AndAssign(pos.DoublePieceTypeColorBB(p.Rook, p.Dragon, us))
					pos.State.CheckersBB = pos.State.CheckersBB.OrAssign(appendBB)
				case sq.DirecDiagNESW:
					appendBB := pos.AttacksFrom(p.Bishop, c.ColorNum, ksq).AndAssign(pos.DoublePieceTypeColorBB(p.Bishop, p.Horse, us))
					pos.State.CheckersBB = pos.State.CheckersBB.OrAssign(appendBB)
				case sq.DirecDiagNWSE:
					appendBB := pos.AttacksFrom(p.Bishop, c.ColorNum, ksq).AndAssign(pos.DoublePieceTypeColorBB(p.Bishop, p.Horse, us))
					pos.State.CheckersBB = pos.State.CheckersBB.OrAssign(appendBB)
				}
			}
			pos.State.ContinuousCheck[us] += 2

		} else {
			pos.State.CheckersBB = bb.NewBitBoardAllZero()
			pos.State.ContinuousCheck[us] = 0
		}

	}

	pos.GoldsBB = pos.QuintPieceTypeBB(p.Gold, p.ProPawn, p.ProLance, p.ProKnight, p.ProSilver)

	pos.State.boardKey = boardKey
	pos.State.HandKey = handKey
	pos.State.pliesFromNull++

	pos.Turn = us.Opposite()
	pos.State.Hand = pos.Hand[pos.Turn]

	//assert(isOK());
}

func (pos *Position) undoMove(mv m.Move) {
	// assert(isOK());
	// assert(!move.isNone());

	them := pos.Turn
	us := them.Opposite()
	to := mv.To()

	// ここで先に turn_ を戻したので、以下、move は us の指し手とする。
	if mv.IsDrop() {
		ptTo := mv.PieceTypeDropped()
		pos.ByTypeBB[ptTo] = pos.ByTypeBB[ptTo].XorBit(to)
		pos.ByColorBB[us] = pos.ByColorBB[us].XorBit(to)
		pos.Piece[to] = p.Empty

		hp := ptTo.ToHandPiece()
		pos.Hand[us] = pos.Hand[us].PlusOne(hp)

		toListIndex := pos.EvalList.SquareHandToList[to]
		handnum := int(pos.Hand[us].NumOf(hp))
		pos.EvalList.List0[toListIndex] = ev.KPPHandArray[us][hp] + handnum
		pos.EvalList.List1[toListIndex] = ev.KPPHandArray[them][hp] + handnum

		squareHand := evl.HandPieceSToSquareHand[us][hp] + sq.Square(handnum)
		pos.EvalList.ListToSquareHand[toListIndex] = squareHand
		pos.EvalList.SquareHandToList[squareHand] = toListIndex
	} else {
		from := mv.From()
		ptFrom := mv.PieceTypeFrom()
		ptTo := mv.PieceTypeToFast(ptFrom)
		ptCaptured := mv.Cap()

		if ptTo == p.King {
			pos.KingSquare[us] = from
		} else {
			pcFrom := p.NewPieceWithColorAndPieceType(us, ptFrom)
			toListIndex := pos.EvalList.SquareHandToList[to]
			pos.EvalList.List0[toListIndex] = ev.KPPArray[pcFrom] + int(to)
			pos.EvalList.List1[toListIndex] = ev.KPPArray[pcFrom.Inverse()] + int(to.Inverse())
			pos.EvalList.ListToSquareHand[toListIndex] = from
			pos.EvalList.SquareHandToList[from] = toListIndex
		}

		if ptCaptured != 0 {
			//駒を取った時
			pos.ByTypeBB[ptCaptured] = pos.ByTypeBB[ptCaptured].XorBit(to)
			pos.ByColorBB[them] = pos.ByColorBB[them].XorBit(to)

			hpCaptured := ptCaptured.ToHandPiece()
			pcCaptured := p.NewPieceWithColorAndPieceType(them, ptCaptured)
			pos.Piece[to] = pcCaptured

			handnum := pos.Hand[us].NumOf(hpCaptured)
			toListIndex := pos.EvalList.SquareHandToList[uint32(evl.HandPieceSToSquareHand[us][hpCaptured])+handnum]
			pos.EvalList.List0[toListIndex] = ev.KPPArray[pcCaptured] + int(to)
			pos.EvalList.List1[toListIndex] = ev.KPPArray[pcCaptured.Inverse()] + int(to.Inverse())
			pos.EvalList.ListToSquareHand[toListIndex] = to
			pos.EvalList.SquareHandToList[to] = toListIndex

			pos.Hand[us] = pos.Hand[us].MinusOne(hpCaptured)
		} else {
			// 駒を取らないときは、colorAndPieceTypeToPiece(us, ptCaptured) は 0 または 16 になる。
			// 16 になると困るので、駒を取らないときは明示的に Empty にする。
			pos.Piece[to] = p.Empty
		}

		pos.ByTypeBB[ptFrom] = pos.ByTypeBB[ptFrom].XorBit(from)
		pos.ByTypeBB[ptTo] = pos.ByTypeBB[ptTo].XorBit(to)
		pos.ByColorBB[us] = pos.ByColorBB[us].XorBitDouble(from, to)
		pos.Piece[from] = p.NewPieceWithColorAndPieceType(us, ptFrom)
	}

	pos.ByTypeBB[p.Occupied] = pos.ByColorBB[c.Black].OrAssign(pos.ByColorBB[c.White])
	pos.GoldsBB = pos.QuintPieceTypeBB(p.Gold, p.ProPawn, p.ProLance, p.ProKnight, p.ProSilver)

	// key などは StateInfo にまとめられているので、
	// previous のポインタを st_ に代入するだけで良い。
	pos.State = pos.State.Previous

	//assert(isOK());
}

// Score see(const Move move, const int asymmThreshold = 0) const;
// Score seeSign(const Move move) const;

// template <Color US> Move mateMoveIn1Ply();
// Move mateMoveIn1Ply();

//////////////////////////////////////////////////////////////
// Keyやら　Plyやらのgetter

// u64 nodesSearched() const          { return nodes_; }
// void setNodesSearched(const u64 n) { nodes_ = n; }
// RepetitionType isDraw(const int checkMaxPly = std::numeric_limits<int>::max()) const;

// Thread* thisThread() const { return thisThread_; }

// void setStartPosPly(const Ply ply) { gamePly_ = ply; }

// static constexpr int nlist() { return EvalList::ListSize; }
// int list0(const int index) const { return evalList_.list0[index]; }
// int list1(const int index) const { return evalList_.list1[index]; }
// int squareHandToList(const Square sq) const { return evalList_.squareHandToList[sq]; }
// Square listToSquareHand(const int i) const { return evalList_.listToSquareHand[i]; }
// int* plist0() { return &evalList_.list0[0]; }
// int* plist1() { return &evalList_.list1[0]; }
// const int* cplist0() const { return &evalList_.list0[0]; }
// const int* cplist1() const { return &evalList_.list1[0]; }
// const ChangedLists& cl() const { return st_->cl; }

// const Searcher* csearcher() const { return searcher_; }
// Searcher* searcher() const { return searcher_; }
// void setSearcher(Searcher* s) { searcher_ = s; }

//////////////////////////////////////////////////////////////
// Zobrist

func initZobrist() {
	// zobTurn_ は 1 であり、その他は 1桁目を使わない。
	// zobTurn のみ xor で更新する為、他の桁に影響しないようにする為。
	// hash値の更新は普通は全て xor を使うが、持ち駒の更新の為に +, - を使用した方が都合が良い。
	for pt := p.Occupied; pt < p.PieceTypeNum; pt++ {
		for square := sq.I9; square < sq.SquareNum; square++ {
			for color := c.Black; color < c.ColorNum; color++ {
				zobrist[pt][square][color] = com.Key(rand.Int63()) & ^com.Key(1)
			}
		}
	}
	for hp := p.HPawn; hp < p.HandPieceNum; hp++ {
		zobHand[hp][c.Black] = com.Key(rand.Int63()) & ^com.Key(1)
		zobHand[hp][c.White] = com.Key(rand.Int63()) & ^com.Key(1)
	}

	zobExclusion = com.Key(rand.Int63()) & ^com.Key(1)
}

//////////////////////////////////////////////////////////////
// Private Methods
func (pos *Position) clear() {

}

// 駒を置く
func (pos *Position) setPiece(piece p.Piece, square sq.Square) {
	color := piece.Color()
	pieceType := piece.PieceType()

	pos.Piece[square] = piece

	pos.ByTypeBB[pieceType].SetBit(square)
	pos.ByColorBB[color].SetBit(square)
	pos.ByTypeBB[p.Occupied].SetBit(square)
}

// 持ち駒をセット(handPiece)
func (pos *Position) setHandWithHandPiece(hp p.HandPiece, color c.Color, num uint32) {
	pos.Hand[color].OrEqual(num, hp)
}

// 持ち駒をセット(Piece)
func (pos *Position) setHandWithPiece(piece p.Piece, num uint32) {
	color := piece.Color()
	pt := piece.PieceType()
	hp := pt.ToHandPiece()

	pos.setHandWithHandPiece(hp, color, num)
}

// 手番側の玉へ check している駒を全て探して checkersBB_ にセットする。
// 最後の手が何か覚えておけば、attackersTo() を使用しなくても良いはずで、処理が軽くなる。
func (pos *Position) findCheckers() {
	pos.State.CheckersBB = pos.attackersToExceptKing(pos.Turn.Opposite(), pos.KingSquare[pos.Turn])
}

//void xorBBs(const PieceType pt, const Square sq, const Color c);

// turn() 側が
// pin されて(して)いる駒の Bitboard を返す。
// BetweenIsUs == true  : 間の駒が自駒。
// BetweenIsUs == false : 間の駒が敵駒。
//template <bool FindPinned, bool BetweenIsUs> Bitboard hiddenCheckers() const {

//	Key key() const { return boardKey + handKey; }
// Key computeBoardKey() const;
// Key computeHandKey() const;
// Key computeKey() const { return computeBoardKey() + computeHandKey(); }

// StateStackPtrはユニークポインタ。Goにはないのでロジックレベルでの書き換え必須
//using StateStackPtr = std::unique_ptr<std::stack<StateInfo> >;
