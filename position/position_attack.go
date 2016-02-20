package position

import (
	bb "github.com/zaramme/gogo-shogi/bitboard"
	c "github.com/zaramme/gogo-shogi/color"
	h "github.com/zaramme/gogo-shogi/hand"
	p "github.com/zaramme/gogo-shogi/piece"
	psc "github.com/zaramme/gogo-shogi/pieceScore"
	sc "github.com/zaramme/gogo-shogi/score"
	sq "github.com/zaramme/gogo-shogi/square"
)

//func attackersTo(const Square sq, const Bitboard& occupied) const;

//玉以外で sq へ移動可能な c 側の駒の Bitboard を返す。
func (pos Position) attackersToExceptKing(color c.Color, square sq.Square) bb.BitBoard {
	opposite := color.Opposite()

	attackersBB := bb.NewBitBoardAllZero()
	setMinorPiece := func(pt p.PieceType) {
		// 利きが生じるすべての座標を取得し、
		ptBB := pos.AttacksFrom(pt, opposite, square)
		// その中からposに実際にその駒が存在している座標のみを抽出
		ptBB = ptBB.AndAssign(pos.ByTypeBB[pt])
		attackersBB = attackersBB.OrAssign(ptBB)
	}

	setMinorPiece(p.Pawn)
	setMinorPiece(p.Lance)
	setMinorPiece(p.Knight)

	//| (attacksFrom<Silver>(opposite, sq) & bbOf(Silver, Dragon))
	silvers := pos.AttacksFrom(p.Silver, color, square)
	silvers = silvers.AndAssign(pos.DoublePieceTypeBB(p.Silver, p.Dragon))
	attackersBB = attackersBB.OrAssign(silvers)

	//| (attacksFrom<Gold  >(opposite, sq) & (goldsBB() | bbOf(Horse)))
	golds := pos.AttacksFrom(p.Gold, color, square)
	goldsBBandHorse := pos.GoldsBB.OrAssign(pos.ByTypeBB[p.Horse])
	golds = golds.AndAssign(goldsBBandHorse)
	attackersBB = attackersBB.OrAssign(golds)

	// 		| (attacksFrom<Bishop>(          sq, occupied) & bbOf(Bishop, Horse        ))
	bishopBB := pos.AttacksFrom(p.Bishop, color, square)
	bishopBB = bishopBB.AndAssign(pos.DoublePieceTypeBB(p.Bishop, p.Horse))
	attackersBB = attackersBB.OrAssign(bishopBB)

	// 		| (attacksFrom<Rook  >(          sq, occupied) & bbOf(Rook  , Dragon       ))
	rookBB := pos.AttacksFrom(p.Rook, color, square)
	rookBB = rookBB.AndAssign(pos.DoublePieceTypeBB(p.Bishop, p.Horse))
	attackersBB = attackersBB.OrAssign(bishopBB)

	attackersBB = attackersBB.AndAssign(pos.ByColorBB[color])
	return attackersBB
}

// Occupiedを任意指定する場合のAttarckersTo。
// 玉の利きも考慮される。
func (pos Position) attackersToExOccupied(color c.Color, square sq.Square, occupied *bb.BitBoard) bb.BitBoard {
	opposite := color.Opposite()

	attackersBB := bb.NewBitBoardAllZero()

	setMinorPiece := func(pt p.PieceType) {
		// 利きが生じるすべての座標を取得し、
		ptBB := pos.AttacksFromExOccupied(pt, opposite, square, occupied)
		// その中からposに実際にその駒が存在している座標のみを抽出
		ptBB = ptBB.AndAssign(pos.ByTypeBB[pt])
		attackersBB = attackersBB.OrAssign(ptBB)
	}
	// return ((attacksFrom<Pawn  >(opposite, sq          ) & bbOf(Pawn  ))
	// 		| (attacksFrom<Lance >(opposite, sq, occupied) & bbOf(Lance ))
	// 		| (attacksFrom<Knight>(opposite, sq          ) & bbOf(Knight))
	// 		| (attacksFrom<Silver>(opposite, sq          ) & bbOf(Silver))
	setMinorPiece(p.Pawn)
	setMinorPiece(p.Lance)
	setMinorPiece(p.Knight)
	setMinorPiece(p.Silver)
	// 		| (attacksFrom<Gold  >(opposite, sq          ) & goldsBB())
	goldsBB := pos.AttacksFromExOccupied(p.Gold, color, square, occupied)
	goldsBB = goldsBB.AndAssign(pos.GoldsBB)
	attackersBB = attackersBB.OrAssign(goldsBB)
	// 		| (attacksFrom<Bishop>(          sq, occupied) & bbOf(Bishop, Horse        ))
	bishopBB := pos.AttacksFromExOccupied(p.Bishop, color, square, occupied)
	bishopBB = bishopBB.AndAssign(pos.DoublePieceTypeBB(p.Bishop, p.Horse))
	attackersBB = attackersBB.OrAssign(bishopBB)

	// 		| (attacksFrom<Rook  >(          sq, occupied) & bbOf(Rook  , Dragon       ))
	rookBB := pos.AttacksFromExOccupied(p.Rook, color, square, occupied)
	rookBB = rookBB.AndAssign(pos.DoublePieceTypeBB(p.Bishop, p.Horse))
	attackersBB = attackersBB.OrAssign(bishopBB)

	// 		| (attacksFrom<King  >(          sq          ) & bbOf(King  , Horse, Dragon)))
	kingBB := pos.AttacksFromExOccupied(p.King, color, square, occupied)
	kingBB = kingBB.AndAssign(pos.TriplePieceTypeBB(p.King, p.Bishop, p.Horse))
	attackersBB = attackersBB.OrAssign(kingBB)

	// 	& bbOf(c);
	// 最後に先後でフィルタ
	attackersBB = attackersBB.AndAssign(pos.ByColorBB[color])

	return attackersBB
}

// 先手、後手に関わらず、sq へ移動可能な駒の利きを Bitboard を返す。
func (pos Position) AttackersToAnyColor(square sq.Square, occupied *bb.BitBoard) bb.BitBoard {

	setMinorPiece := func(target *bb.BitBoard, pt p.PieceType, color c.Color) {
		// 利きが生じるすべての座標を取得し、
		ptBB := pos.AttacksFromExOccupied(pt, color, square, occupied)
		// その中からposに実際にその駒が存在している座標のみを抽出
		ptBB = ptBB.AndAssign(pos.ByTypeBB[pt])
		*target = target.OrAssign(ptBB)
	}

	golds := pos.GoldsBB
	setGolds := func(target *bb.BitBoard, color c.Color) {
		// 利きが生じるすべての座標を取得し、
		ptBB := pos.AttacksFromExOccupied(p.Gold, color, square, occupied)
		// その中からposに実際にその駒が存在している座標のみを抽出
		ptBB = ptBB.AndAssign(golds)
		*target = target.OrAssign(ptBB)
	}

	blacks := bb.NewBitBoardAllZero()
	setMinorPiece(&blacks, p.Pawn, c.Black)
	setMinorPiece(&blacks, p.Lance, c.Black)
	setMinorPiece(&blacks, p.Knight, c.Black)
	setMinorPiece(&blacks, p.Silver, c.Black)
	setGolds(&blacks, c.Black)

	blacks = blacks.AndAssign(pos.ByColorBB[c.Black])

	whites := bb.NewBitBoardAllZero()
	// pawn
	setMinorPiece(&whites, p.Pawn, c.White)
	setMinorPiece(&whites, p.Lance, c.White)
	setMinorPiece(&whites, p.Knight, c.White)
	setMinorPiece(&whites, p.Silver, c.White)
	setGolds(&whites, c.White)

	whites = whites.AndAssign(pos.ByColorBB[c.White])

	bishops := pos.AttacksFromExOccupied(p.Bishop, c.ColorNum, square, occupied)
	bishops = bishops.AndAssign(pos.DoublePieceTypeBB(p.Bishop, p.Horse))

	rooks := pos.AttacksFromExOccupied(p.Rook, c.ColorNum, square, occupied)
	rooks = rooks.AndAssign(pos.DoublePieceTypeBB(p.Rook, p.Dragon))

	kings := pos.AttacksFromExOccupied(p.King, c.ColorNum, square, occupied)
	kings = kings.AndAssign(pos.TriplePieceTypeBB(p.King, p.Horse, p.Dragon))

	return blacks.
		AndAssign(whites).
		AndAssign(bishops).
		AndAssign(rooks).
		AndAssign(kings)
}

// 以下は実装せずattackersTo().isNot0で対応
// bool attackersToIsNot0(const Color c, const Square sq) const { return attackersTo(c, sq).isNot0(); }
// bool attackersToIsNot0(const Color c, const Square sq, const Bitboard& occupied) const {

// 移動王手が味方の利きに支えられているか。false なら相手玉で取れば詰まない。
func (pos Position) unDropCheckIsSupported(color c.Color, square sq.Square) bool {
	return pos.attackersToExceptKing(color, square).IsNot0()
}

func (pos Position) computeMaterial() sc.Score {
	score := sc.Zero
	for pt := p.Pawn; pt < p.PieceTypeNum; pt++ {
		num := pos.PieceTypeColorBB(pt, c.Black).PopCount() - pos.PieceTypeColorBB(pt, c.White).PopCount()
		score += sc.Score(num) * psc.PieceScore[pt]
	}
	for pt := p.Pawn; pt < p.King; pt++ {
		num := h.Hand(c.Black).NumOf(pt.ToHandPiece()) - h.Hand(c.White).NumOf(pt.ToHandPiece())
		score += sc.Score(num) * psc.PieceScore[pt]
	}
	return score
}

// 利きを生成する。
// Occupied情報はposition内部から取得する。
func (pos Position) AttacksFrom(
	pt p.PieceType, color c.Color, square sq.Square) bb.BitBoard {
	occupied := pos.OccupiedBB()
	return pos.AttacksFromExOccupied(pt, color, square, &occupied)
}

// 任意のsq座標とoccupied情報に対して、特定の駒、先後で利いている駒の位置のbitBoardを返す。
func (pos Position) AttacksFromExOccupied(
	pt p.PieceType, color c.Color, square sq.Square, occupied *bb.BitBoard) bb.BitBoard {
	// C++ではマクロを用いて条件分岐を回避しているため、
	// ここは処理が超絶遅くなっていいる
	switch pt {
	case p.Occupied:
		return bb.NewBitBoardAllZero()
	case p.Pawn:
		return bb.PawnAttack[color][square]
	case p.Lance:
		return bb.GetLanceAttack(color, square, occupied)
	case p.Knight:
		return bb.KnightAttack[color][square]
	case p.Silver:
		return bb.SilverAttack[color][square]
	case p.Bishop:
		return bb.GetBishopAttack(square, occupied)
	case p.Rook:
		//return bb.GetRookAttack(square, occupied)
	case p.King:
		//return bb.KingAttack[sq]
	case p.Horse:
		//return bb.HorceAttack(square, occupied)
	case p.Dragon:
		//return bb.DragonAttack(square, occupied)
	}

	// 通常では到達しない
	return bb.NewBitBoardAllZero()
}
