package tt

import (
	mv "github.com/zaramme/gogo-shogi/move"
	sc "github.com/zaramme/gogo-shogi/score"
)

type Depth int

const (
	OnePly           Depth = 2
	Depth0           Depth = 0
	Depth1           Depth = 1
	DepthQChecks     Depth = -1 * OnePly
	DepthQNoChecks   Depth = -2 * OnePly
	DepthQRecaptures Depth = -5 * OnePly
	DepthNone        Depth = -127 * OnePly
)

type TTEntry struct {
	key32       uint32
	move16      uint16
	bound       uint8
	generation8 uint8
	score16     int16
	depth16     int16
	evalScore   int16
}

////////////////////////////
// getter Methods
func (t TTEntry) Key() uint32 {
	return t.key32
}

func (t TTEntry) Depth() Depth {
	return Depth(t.depth16)
}

func (t TTEntry) Score() sc.Score {
	return sc.Score(t.score16)
}

func (t TTEntry) Move() mv.Move {
	return mv.Move(t.move16)
}

func (t TTEntry) Type() sc.Bound {
	return sc.Bound(t.bound)
}

func (t TTEntry) SetGeneration(g uint8) {
	t.generation8 = g
}

func (t TTEntry) Save(dp Depth,
	sc sc.Score, mv mv.Move,
	posKeyHigh32 uint32, bo sc.Bound,
	gen uint8, esc sc.Score) {
	t.key32 = posKeyHigh32
	t.move16 = uint16(mv)
	t.bound = uint8(bo)
	t.generation8 = gen
	t.score16 = int16(sc)
	t.depth16 = int16(dp)
	t.evalScore = int16(esc)
}
