package position

import (
	bb "github.com/zaramme/gogo-shogi/bitboard"
	c "github.com/zaramme/gogo-shogi/color"
	com "github.com/zaramme/gogo-shogi/common"
	h "github.com/zaramme/gogo-shogi/hand"
	p "github.com/zaramme/gogo-shogi/piece"
	sc "github.com/zaramme/gogo-shogi/score"
)

type CheckInfo struct {
	dcBB, pined bb.BitBoard
	checkBB     [p.PieceTypeNum]bb.BitBoard
}

// 現在の盤面で王手できる駒の情報をcheckBBに格納する
func (ch *CheckInfo) SetCheckInfo(pos *Position) {
	them := pos.Turn.Opposite()
	ksq := pos.KingSquare[them]

	ch.pined = pos.pinnedBB()
	ch.dcBB = pos.DisCoverdCheckBB()
	ch.checkBB[p.Pawn] = pos.AttacksFrom(p.Pawn, them, ksq)
	ch.checkBB[p.Lance] = pos.AttacksFrom(p.Lance, them, ksq)
	ch.checkBB[p.Knight] = pos.AttacksFrom(p.Knight, them, ksq)
	ch.checkBB[p.Silver] = pos.AttacksFrom(p.Silver, them, ksq)
	ch.checkBB[p.Bishop] = pos.AttacksFrom(p.Bishop, c.ColorNum, ksq)
	ch.checkBB[p.Rook] = pos.AttacksFrom(p.Rook, c.ColorNum, ksq)
	ch.checkBB[p.Gold] = pos.AttacksFrom(p.Gold, them, ksq)
	ch.checkBB[p.King] = bb.NewBitBoardAllZero()

	ch.checkBB[p.ProPawn] = ch.checkBB[p.Gold]
	ch.checkBB[p.ProLance] = ch.checkBB[p.Gold]
	ch.checkBB[p.ProKnight] = ch.checkBB[p.Gold]
	ch.checkBB[p.ProSilver] = ch.checkBB[p.Gold]
	ch.checkBB[p.Horse] = ch.checkBB[p.Bishop].AndAssign(pos.AttacksFrom(p.King, c.ColorNum, ksq))
	ch.checkBB[p.Dragon] = ch.checkBB[p.Dragon].AndAssign(pos.AttacksFrom(p.King, c.ColorNum, ksq))
}

// 一手の着手の中で変化する駒を表現
type ChangedLists struct {
	ClistPair [2]ChangedListPair // 一手で動く駒は最大2つ。(動く駒、取られる駒)
	ListIndex [2]int             // 一手で動く駒は最大2つ。(動く駒、取られる駒)
	Size      int
}

type ChangedListPair struct {
	newList [2]int
	oldList [2]int
}

type StateInfoMin struct {
	material        sc.Score
	pliesFromNull   int
	ContinuousCheck [c.ColorNum]int
}

type StateInfo struct {
	StateInfoMin
	boardKey     com.Key
	HandKey      com.Key
	CheckersBB   bb.BitBoard
	Previous     *StateInfo
	Hand         h.Hand
	ChangedLists ChangedLists //Aperyでは(cl)
}

type threadInterface interface {
	notifyone()
	cutoffOccurrred() bool
	isAvailableTo(*threadInterface) bool
	waitfor(bool)
	//	split(pos Position,ss *SarchStack, alpha,beta sc.Score, bestScore *sc.Score.
	//	bestMove mv.Move, depth Depth, move ThreadMove, moveCount int,
	//  mp *MovePicker, nodeType NodeType, cutNode bool)
}

type searchInterface interface {
}
