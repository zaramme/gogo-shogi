package score

import (
	"math"
)

type Ply int

const (
	MaxPly      Ply = 128
	MaxPlyPlus2 Ply = MaxPly + 2
)

// 評価値の状態を示す
type Bound int

const (
	BoundNone Bound = 0 // Bynary<0>
	Upper     Bound = 1 // Bynary<01> fail low  で正しい score が分からない。alpha 以下が確定という意味。
	Lower     Bound = 2 // Bynary<10> fail high  で正しい score が分からない。bata 以上が確定という意味。
	Exact     Bound = 3 // Bynary<11> alpha, bataの間で評価値が確定している。
)

// inline関数のため、高速化するには、直接処理を書く必要がある
func (b Bound) ExactOrLower() bool {
	v := b & Lower
	return v != 0
}

func (b Bound) ExactOrUpper() bool {
	v := b & Upper
	return v != 0
}

type Score int

const (
	Zero          Score = 0
	Draw          Score = 0
	MaxEvaluate   Score = 30000
	MateLong      Score = 30002
	Mate1Ply      Score = 32559
	Mate0Ply      Score = 32600
	MateInMaxPly  Score = Mate0Ply - Score(MaxPly)
	MatedInMaxPly Score = Zero - MateInMaxPly
	Inifinite     Score = 32601
	NotEvaluated  Score = math.MaxInt64
	ScoreNone     Score = 32602
)

func (s Score) MateIn(ply Ply) Score {
	return Mate0Ply - Score(ply)
}

func (s Score) MatedIn(ply Ply) Score {
	return -Mate0Ply + Score(ply)
}
