package move

import (
	sq "github.com/zaramme/gogo-shogi/square"
	//	"math"
	"testing"
)

func Test_moveTo(t *testing.T) {
	testcase := func(id *int, m Move, expect sq.Square) {
		actual := m.To()
		if actual != expect {
			t.Errorf("[%d]期待値と実測値が異なります。expect = %x, actual = %x", *id, expect, actual)
		}
		*id++
	}

	id := 1

	testcase(&id, Move(0xFFFFFF00), sq.I9)
	testcase(&id, Move(0xFFFFFF50), sq.A1)

}

func Test_moveFrom(t *testing.T) {
	testcase := func(id *int, m Move, expect sq.Square) {
		actual := m.From()
		if actual != expect {
			t.Errorf("[%d]期待値と実測値が異なります。expect = %x, actual = %x", *id, expect, actual)
		}
		*id++
	}

	id := 1

	testcase(&id, Move(0xffffc07f), sq.I9)
	testcase(&id, Move(0xffffe87f), sq.A1)

}

func Test_moveFromAndTo(t *testing.T) {
	testCase := func(id *int, m Move, expect uint32) {
		actual := m.FromAndTo()
		if actual != expect {
			t.Errorf("[%d]期待値と実測値が異なります。expect = %x, actual = %x", *id, expect, actual)
		}
		*id++
	}

	id := 0
	m := Move(0x9fce2986)
	expect := m.To() | (m.From() << 7)
	testCase(&id, m, uint32(expect))

	m = Move(0xe324acb8)
	expect = m.To() | (m.From() << 7)
	testCase(&id, m, uint32(expect))
}
