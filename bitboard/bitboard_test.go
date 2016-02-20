package bitboard

import (
	"testing"
)

func Test_Merge(t *testing.T) {

	testCase := func(id *int, input BitBoard, expect uint64) {
		actual := input.Merge()
		if actual != expect {
			t.Errorf("[id = %s]期待値と実際値が異なります。input=%v,expect=%v,actual=%v", id, input, expect, actual)
		}
		*id++
	}

	var input BitBoard
	var id int

	id = 0
	input = BitBoard{0xF0F0F0F0F0F0F0F0, 0x0F0F0F0F0F0F0F0F}
	testCase(&id, input, 0xFFFFFFFFFFFFFFFF)

	input = BitBoard{0xCCCCCCCCCCCCCCCC, 0x3333333333333333}
	testCase(&id, input, 0xFFFFFFFFFFFFFFFF)

}

func Test_isNot0(t *testing.T) {
	testCase := func(id *int, input BitBoard, expect bool) {
		actual := input.IsNot0()

		if actual != expect {
			t.Errorf("[id = %s]期待値と実際値が異なります。input=%v,expect=%v,actual=%v", id, input, expect, actual)
		}
		*id++
	}

	var input BitBoard
	var id int

	id = 0
	input = BitBoard{0, 0}
	testCase(&id, input, false)

	input = BitBoard{1, 0}
	testCase(&id, input, true)

	input = BitBoard{0, 1}
	testCase(&id, input, true)

}

func Test_AndIsNot0(t *testing.T) {
	testCase := func(id *int, b1, b2 BitBoard, expect bool) {
		actual := b1.AndIsNot0(b2)

		if actual != expect {
			t.Errorf("[id = %s]期待値と実際値が異なります。\nb1={%X}\nb2={%X}\n,expect=%v,actual=%v", id, b1, b2, expect, actual)
		}
		*id++
	}

	var b1, b2 BitBoard
	var id int

	id = 0
	b1 = BitBoard{0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF}
	b2 = BitBoard{0x00000000FFFFFFFF, 0xFFFFFFFF00000000}
	testCase(&id, b1, b2, true)

	b1 = BitBoard{0xFFFFFFFF00000000, 0x00000000FFFFFFFF}
	b2 = BitBoard{0x00000000FFFFFFFF, 0xFFFFFFFF00000000}
	testCase(&id, b1, b2, false)

}

func Test_Print(t *testing.T) {
	t.SkipNow()
	b1 := BitBoard{0xFFFFFFFFFFFFFFFF, 0x0000000000000000}
	b1.Print()
}
