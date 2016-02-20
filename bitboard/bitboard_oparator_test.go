package bitboard

import (
	"testing"
)

func Test_inverse(t *testing.T) {

	bb := BitBoard{0xFFFFFFFFFFFFFFFF, 0x0000000000000000}
	bb2 := bb.Inverse()

	if bb2[0] != 0x000000000000000 || bb2[1] != 0xFFFFFFFFFFFFFFFF {
		t.Errorf("[id=1]ビットボード反転に失敗しました。, [0]=%b, [1]=%b", bb2[0], bb2[1])
	}

	bb = BitBoard{0xF0F0F0F0F0F0F0F0, 0x0F0F0F0F0F0F0F0F}
	bb2 = bb.Inverse()

	if bb2[0] != 0x0F0F0F0F0F0F0F0F || bb2[1] != 0xF0F0F0F0F0F0F0F0 {
		t.Errorf("[id=2]ビットボード反転に失敗しました。, [0]=%b, [1]=%b", bb2[0], bb2[1])
	}
}

func Test_EqualTo(t *testing.T) {

	testCase := func(id *int, bb1, bb2 BitBoard, expect bool) {
		// actual := bb1.EqualTo(bb2)
		// if actual != expect {
		if (bb1 == bb2) != expect {
			t.Errorf("[id = %d]期待値と異なる結果です。bb1 = [%b,%b], bb2 = [%b,%b], expect = %t",
				*id, bb1[0], bb1[1], bb2[0], bb2[1], expect)
			t.Fatal()
		}
		*id++
	}

	var id int
	var bb1 BitBoard
	var bb2 BitBoard

	id = 0
	bb1 = BitBoard{0xFFFFFFFFFFFFFFFF, 0x0000000000000000}
	bb2 = BitBoard{0xFFFFFFFFFFFFFFFF, 0x0000000000000000}
	testCase(&id, bb1, bb2, true)

	bb1 = BitBoard{0x0000000000000000, 0xFFFFFFFFFFFFFFFF}
	bb2 = BitBoard{0x0000000000000000, 0xFFFFFFFFFFFFFFFF}
	testCase(&id, bb1, bb2, true)

	bb1 = BitBoard{0xFFFFFFFFFFFFFFFF, 0x0000000000000000}
	bb2 = BitBoard{0xFFFFFFFFFFFFFFFE, 0x0000000000000000}
	testCase(&id, bb1, bb2, false)

	bb1 = BitBoard{0xFFFFFFFFFFFFFFFF, 0x0000000000000000}
	bb2 = BitBoard{0xFFFFFFFFFFFFFFFF, 0x0000000000000001}
	testCase(&id, bb1, bb2, false)

	bb1 = BitBoard{0x1000000000000000, 0xFFFFFFFFFFFFFFFF}
	bb2 = BitBoard{0x0000000000000000, 0xFFFFFFFFFFFFFFFF}
	testCase(&id, bb1, bb2, false)

	bb1 = BitBoard{0x0000000000000000, 0xEFFFFFFFFFFFFFFF}
	bb2 = BitBoard{0x0000000000000000, 0xFFFFFFFFFFFFFFFF}
	testCase(&id, bb1, bb2, false)

}

func Test_NotEqualTo(t *testing.T) {

	testCase := func(id *int, bb1, bb2 BitBoard, expect bool) {
		//actual := bb1.NotEqualTo(bb2)
		//if actual != expect {
		if (bb1 != bb2) != expect {
			t.Errorf("[id = %d]期待値と異なる結果です。bb1 = [%b,%b], bb2 = [%b,%b], expect = %t ",
				*id, bb1[0], bb1[1], bb2[0], bb2[1], expect)
			t.Fatal()
		}
		*id++
	}

	var id int
	var bb1 BitBoard
	var bb2 BitBoard

	id = 0
	bb1 = BitBoard{0xFFFFFFFFFFFFFFFF, 0x0000000000000000}
	bb2 = BitBoard{0xFFFFFFFFFFFFFFFF, 0x0000000000000000}
	testCase(&id, bb1, bb2, false)

	bb1 = BitBoard{0x0000000000000000, 0xFFFFFFFFFFFFFFFF}
	bb2 = BitBoard{0x0000000000000000, 0xFFFFFFFFFFFFFFFF}
	testCase(&id, bb1, bb2, false)

	bb1 = BitBoard{0xFFFFFFFFFFFFFFFF, 0x0000000000000000}
	bb2 = BitBoard{0xFFFFFFFFFFFFFFFE, 0x0000000000000000}
	testCase(&id, bb1, bb2, true)

	bb1 = BitBoard{0xFFFFFFFFFFFFFFFF, 0x0000000000000000}
	bb2 = BitBoard{0xFFFFFFFFFFFFFFFF, 0x0000000000000001}
	testCase(&id, bb1, bb2, true)

	bb1 = BitBoard{0x1000000000000000, 0xFFFFFFFFFFFFFFFF}
	bb2 = BitBoard{0x0000000000000000, 0xFFFFFFFFFFFFFFFF}
	testCase(&id, bb1, bb2, true)

	bb1 = BitBoard{0x0000000000000000, 0xEFFFFFFFFFFFFFFF}
	bb2 = BitBoard{0x0000000000000000, 0xFFFFFFFFFFFFFFFF}
	testCase(&id, bb1, bb2, true)

}

func Test_AndAssign(t *testing.T) {

	testCase := func(id *int, bb1, bb2 BitBoard, expect BitBoard) {
		actual := bb1.AndAssign(bb2)
		if actual != expect {
			t.Errorf("[id = %d]期待値と異なる結果です。	\n bb1 = [%b,%b], \n bb2 = [%b,%b], \n expect = %v actual = %v",
				*id, bb1[0], bb1[1], bb2[0], bb2[1], expect, actual)
		}

		*id++
	}

	var id int
	var bb1 BitBoard
	var bb2 BitBoard
	var expect BitBoard
	id = 0

	bb1 = BitBoard{0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF}
	bb2 = BitBoard{0xFFFFFFFFFFFFFFFF, 0x0000000000000000}
	expect = BitBoard{0xFFFFFFFFFFFFFFFF, 0x0000000000000000}
	testCase(&id, bb1, bb2, expect)

	bb1 = BitBoard{0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF}
	bb2 = BitBoard{0x0000000000000000, 0xFFFFFFFFFFFFFFFF}
	expect = BitBoard{0x0000000000000000, 0xFFFFFFFFFFFFFFFF}
	testCase(&id, bb1, bb2, expect)

	bb1 = BitBoard{0xFFEEDDCCBBAA9988, 0x7766554433221100}
	bb2 = BitBoard{0xF0F0F0F0F0F0F0F0, 0x0F0F0F0F0F0F0F0F}
	expect = BitBoard{0xF0E0D0C0B0A09080, 0x0706050403020100}
	testCase(&id, bb1, bb2, expect)

	return
}
