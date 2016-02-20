package hand

import (
	p "github.com/zaramme/gogo-shogi/piece"
	"strconv"
	"strings"
	"testing"
)

// 二進数文字列をintに
func stob(s string) uint32 {
	// split -> join で半角スペースを削除
	sa := strings.Split(s, " ")
	sb := strings.Join(sa, "")
	v, _ := strconv.ParseInt(sb, 2, 64)
	return uint32(v)
}

// すべての駒がゼロ
func handZero() Hand {
	return Hand(stob("00000000 00000000 00000000 00000000"))
}

// すべての駒が1
func handOne() Hand {
	return Hand(stob("00000010 01000100 01000100 01000001"))
}

// すべての駒がルール上のMAX値
func handMax() Hand {
	return Hand(stob("00000100 10010001 00010001 00010010"))
}

func Test_NumOf(t *testing.T) {

	testCase := func(id *int, input uint32, hp p.HandPiece, expect uint32) {
		h := Hand(input)
		numOfPawn := h.NumOf(hp)
		if numOfPawn != expect {
			t.Errorf("[id = %d]数値が間違っています。 input = %b expect = %d, output = %d", *id, input, expect, numOfPawn)
		}
		*id++
	}

	id := 1
	testCase(&id, stob("0"), p.HPawn, 0)
	testCase(&id, stob("0"), p.HLance, 0)
	testCase(&id, stob("0"), p.HKnight, 0)
	testCase(&id, stob("0"), p.HSilver, 0)
	testCase(&id, stob("0"), p.HGold, 0)
	testCase(&id, stob("0"), p.HBishop, 0)
	testCase(&id, stob("0"), p.HRook, 0)

	id = 10
	testCase(&id, stob("00000000 00000001"), p.HPawn, 1)
	testCase(&id, stob("00000000 00010010"), p.HPawn, 18)

	id = 20
	testCase(&id, stob("00000000 01000000"), p.HLance, 1)
	testCase(&id, stob("00000001 00000000"), p.HLance, 4)

	id = 30
	testCase(&id, stob("00000100 00000000"), p.HKnight, 1)
	testCase(&id, stob("00010000 00000000"), p.HKnight, 4)

	id = 40
	testCase(&id, stob("0 01000000 00000000"), p.HSilver, 1)
	testCase(&id, stob("1 00000000 00000000"), p.HSilver, 4)

	id = 50
	testCase(&id, stob("00100 01000000 00000000"), p.HGold, 1)
	testCase(&id, stob("10001 00000000 00000000"), p.HGold, 4)

	id = 60
	testCase(&id, stob("01000000 00000000 00000000"), p.HBishop, 1)
	testCase(&id, stob("10000000 00000000 00000000"), p.HBishop, 2)

	id = 70
	testCase(&id, stob("010 01000000 00000000 00000000"), p.HRook, 1)
	testCase(&id, stob("100 10000000 00000000 00000000"), p.HRook, 2)

}

func Test_Exists(t *testing.T) {

	testCase := func(id *int, input uint32, hp p.HandPiece, expectB bool, expectI uint32) {
		h := Hand(input)
		expectH := Hand(expectI)
		actualB, actualH := h.Exists(hp)
		if actualB != expectB {
			t.Errorf("[id = %d]論理値が間違っています。 input = %b expect = %b, output = %b", *id, input, expectB, actualB)
		}
		if actualH != expectH {
			t.Errorf("[id = %d]戻り値が間違っています。 input = %b expect = %b, output = %b", *id, input, expectH, actualH)
		}
		*id++
	}

	id := 0
	testCase(&id, stob("0"), p.HPawn, false, 0)
	testCase(&id, stob("0"), p.HLance, false, 0)
	testCase(&id, stob("0"), p.HKnight, false, 0)
	testCase(&id, stob("0"), p.HSilver, false, 0)
	testCase(&id, stob("0"), p.HGold, false, 0)
	testCase(&id, stob("0"), p.HBishop, false, 0)
	testCase(&id, stob("0"), p.HRook, false, 0)

	id = 10
	testCase(&id, stob("11111111 11111111 11111111 11100001"), p.HPawn,
		true, stob("00000000 00000000 00000000 00000001"))
	testCase(&id, stob("11111111 11111111 11111111 11110010"), p.HPawn,
		true, stob("00000000 00000000 00000000 00010010"))

	id = 20
	testCase(&id, stob("11111111 11111111 11111110 01111111"), p.HLance,
		true, stob("00000000 00000000 00000000 01000000"))
	testCase(&id, stob("11111111 11111111 11111111 00111111"), p.HLance,
		true, stob("00000000 00000000 00000001 00000000"))

	id = 30
	testCase(&id, stob("11111111 11111111 11100111 11111111"), p.HKnight,
		true, stob("00000000 00000000 00000100 00000000"))
	testCase(&id, stob("11111111 11111111 11110011 11111111"), p.HKnight,
		true, stob("00000000 00000000 00010000 00000000"))

	id = 40
	testCase(&id, stob("11111111 11111110 01111111 11111111"), p.HSilver,
		true, stob("00000000 00000000 01000000 00000000"))
	testCase(&id, stob("11111111 11111111 00111111 11111111"), p.HSilver,
		true, stob("00000000 00000001 00000000 00000000"))

	id = 50
	testCase(&id, stob("11111111 11100111 11111111 11111111"), p.HGold,
		true, stob("00000000 00000100 00000000 00000000"))
	testCase(&id, stob("11111111 11110011 11111111 11111111"), p.HGold,
		true, stob("00000000 00010000 00000000 00000000"))

	id = 60
	testCase(&id, stob("11111110 01111111 11111111 11111111"), p.HBishop,
		true, stob("00000000 01000000 00000000 00000000"))
	testCase(&id, stob("11111110 10111111 11111111 11111111"), p.HBishop,
		true, stob("00000000 10000000 00000000 00000000"))

	id = 70
	testCase(&id, stob("11111011 01111111 11111111 11111111"), p.HRook,
		true, stob("00000010 00000000 00000000 00000000"))
	testCase(&id, stob("11111101 10111111 11111111 11111111"), p.HRook,
		true, stob("00000100 00000000 00000000 00000000"))

}

func Test_ExceptPawnExists(t *testing.T) {

	testCase := func(id *int, input uint32, expectB bool, expectI uint32) {
		h := Hand(input)
		expectH := Hand(expectI)
		actualB, actualH := h.ExceptPawnExists()
		if actualB != expectB {
			t.Errorf("[id = %d]論理値が間違っています。 input = %b expect = %b, output = %b", *id, input, expectB, actualB)
		}
		if actualH != expectH {
			t.Errorf("[id = %d]戻り値が間違っています。 input = %b expect = %b, output = %b", *id, input, expectH, actualH)
		}
		*id++
	}

	id := 0
	testCase(&id, stob("0"), false, 0)

	id = 10
	testCase(&id, stob("00000010 01000100 01000100 01111111"),
		true, stob("00000010 01000100 01000100 01000000"))
	testCase(&id, stob("00000100 10010001 00010001 00111111"),
		true, stob("00000100 10010001 00010001 00000000"))
}

func Test_OrEqual(t *testing.T) {
	testCase := func(id *int, input Hand, hp p.HandPiece, num uint32, expect uint32) {
		h := Hand(input)
		h = h.OrEqual(num, hp)
		actual := h.NumOf(hp)
		if actual != expect {
			t.Errorf("[id = %d]期待値が間違っています。 input = %b, expect = %d, output = %d", *id, input, expect, actual)
		}
		*id++
	}

	id := 0
	testCase(&id, handOne(), p.HPawn, 0, 1)
	testCase(&id, handOne(), p.HPawn, 1, 1)
	testCase(&id, handOne(), p.HPawn, 2, 3)
	testCase(&id, handOne(), p.HPawn, 17, 17)

	id = 10
	testCase(&id, handOne(), p.HLance, 0, 1)
	testCase(&id, handOne(), p.HLance, 1, 1)
	testCase(&id, handOne(), p.HLance, 2, 3)
	testCase(&id, handOne(), p.HLance, 3, 3)

	id = 20
	testCase(&id, handOne(), p.HKnight, 0, 1)
	testCase(&id, handOne(), p.HKnight, 1, 1)
	testCase(&id, handOne(), p.HKnight, 2, 3)
	testCase(&id, handOne(), p.HKnight, 3, 3)

	id = 30
	testCase(&id, handOne(), p.HSilver, 0, 1)
	testCase(&id, handOne(), p.HSilver, 1, 1)
	testCase(&id, handOne(), p.HSilver, 2, 3)
	testCase(&id, handOne(), p.HSilver, 3, 3)

	id = 40
	testCase(&id, handOne(), p.HGold, 0, 1)
	testCase(&id, handOne(), p.HGold, 1, 1)
	testCase(&id, handOne(), p.HGold, 2, 3)
	testCase(&id, handOne(), p.HGold, 3, 3)

	id = 50
	testCase(&id, handOne(), p.HBishop, 0, 1)
	testCase(&id, handOne(), p.HBishop, 1, 1)

	id = 60
	testCase(&id, handOne(), p.HRook, 0, 1)
	testCase(&id, handOne(), p.HRook, 1, 1)

}

func Test_IsEqualOrSuperior(t *testing.T) {
	testCase := func(id *int, own Hand, to Hand, expect bool) {
		h := Hand(own)
		h2 := Hand(to)
		if h.IsEqualOrSuperior(h2) != expect {
			t.Errorf("[id = %d]期待値が間違っています。 input = %b, %b, expect = %s", *id, own, to, expect)
			t.Errorf("--- this - ref = %b", own-to)
		}
		*id++
	}

	id := 0
	testCase(&id, handOne(), handOne(), true)
	testCase(&id, handZero(), handOne(), false)

	id = 10
	testCase(&id, handOne(), handOne().PlusOne(p.HLance), false)
	testCase(&id, handOne(), handOne().PlusOne(p.HKnight), false)
	testCase(&id, handOne(), handOne().PlusOne(p.HSilver), false)
	testCase(&id, handOne(), handOne().PlusOne(p.HGold), false)
	testCase(&id, handOne(), handOne().PlusOne(p.HBishop), false)
	testCase(&id, handOne(), handOne().PlusOne(p.HRook), false)

	id = 20
	testCase(&id, handOne(), handOne().MinusOne(p.HLance), true)
	testCase(&id, handOne(), handOne().MinusOne(p.HKnight), true)
	testCase(&id, handOne(), handOne().MinusOne(p.HSilver), true)
	testCase(&id, handOne(), handOne().MinusOne(p.HGold), true)
	testCase(&id, handOne(), handOne().MinusOne(p.HBishop), true)
	testCase(&id, handOne(), handOne().MinusOne(p.HRook), true)

	id = 30
	testCase(&id, handOne(), handOne().PlusOne(p.HLance).MinusOne(p.HKnight), false)
	testCase(&id, handOne(), handOne().PlusOne(p.HKnight).MinusOne(p.HSilver), false)
	testCase(&id, handOne(), handOne().PlusOne(p.HSilver).MinusOne(p.HGold), false)
	testCase(&id, handOne(), handOne().PlusOne(p.HGold).MinusOne(p.HBishop), false)
	testCase(&id, handOne(), handOne().PlusOne(p.HBishop).MinusOne(p.HRook), false)
	testCase(&id, handOne(), handOne().PlusOne(p.HRook).MinusOne(p.HPawn), false)

	id = 40
	testCase(&id, handMax(), handMax().MinusOne(p.HPawn), true)
	testCase(&id, handMax(), handMax().MinusOne(p.HLance), true)
	testCase(&id, handMax(), handMax().MinusOne(p.HKnight), true)
	testCase(&id, handMax(), handMax().MinusOne(p.HSilver), true)
	testCase(&id, handMax(), handMax().MinusOne(p.HGold), true)
	testCase(&id, handMax(), handMax().MinusOne(p.HBishop), true)
	testCase(&id, handMax(), handMax().MinusOne(p.HRook), true)

	id = 50
	testCase(&id, handMax().MinusOne(p.HPawn), handMax(), false)
	testCase(&id, handMax().MinusOne(p.HLance), handMax(), false)
	testCase(&id, handMax().MinusOne(p.HKnight), handMax(), false)
	testCase(&id, handMax().MinusOne(p.HSilver), handMax(), false)
	testCase(&id, handMax().MinusOne(p.HGold), handMax(), false)
	testCase(&id, handMax().MinusOne(p.HBishop), handMax(), false)
	testCase(&id, handMax().MinusOne(p.HRook), handMax(), false)

}

func Test_PlusOne(t *testing.T) {

	testCase := func(id *int, input uint32, hp p.HandPiece, expect uint32) {
		h := Hand(input)
		h = h.PlusOne(hp)
		actual := h.NumOf(hp)
		if actual != expect {
			t.Errorf("[id = %d]期待値が間違っています。 input = %b, expect = %d, output = %d", *id, input, expect, h)
		}
		*id++
	}

	id := 0
	testCase(&id, 0, p.HPawn, 1)
	testCase(&id, 17, p.HPawn, 18)

	id = 10
	testCase(&id, 0, p.HLance, 1)
	testCase(&id, stob("00000000 00000000 00000000 11000000"), p.HLance, 4)

	id = 20
	testCase(&id, 0, p.HKnight, 1)
	testCase(&id, stob("00000000 00000000 00001100 00000000"), p.HKnight, 4)

	id = 30
	testCase(&id, 0, p.HSilver, 1)
	testCase(&id, stob("00000000 00000000 11000000 00000000"), p.HSilver, 4)

	id = 40
	testCase(&id, 0, p.HGold, 1)
	testCase(&id, stob("00000000 00001100 00000000 00000000"), p.HGold, 4)

	id = 50
	testCase(&id, 0, p.HBishop, 1)
	testCase(&id, stob("00000000 01000000 00000000 00000000"), p.HBishop, 2)

	id = 60
	testCase(&id, 0, p.HRook, 1)
	testCase(&id, stob("00000010 00000000 00000000 00000000"), p.HRook, 2)

}

func Test_MinusOne(t *testing.T) {

	testCase := func(id *int, input uint32, hp p.HandPiece, expect uint32) {
		h := Hand(input)
		h = h.MinusOne(hp)
		actual := h.NumOf(hp)
		if actual != expect {
			t.Errorf("[id = %d]期待値が間違っています。 input = %b, expect = %d, output = %d", *id, input, expect, h)
		}
		*id++
	}

	id := 0
	testCase(&id, 1, p.HPawn, 0)
	testCase(&id, 18, p.HPawn, 17)

	id = 10
	testCase(&id, stob("00000000 00000000 00000000 01000000"), p.HLance, 0)
	testCase(&id, stob("00000000 00000000 00000001 00000000"), p.HLance, 3)

	id = 20
	testCase(&id, stob("00000000 00000000 00000100 00000000"), p.HKnight, 0)
	testCase(&id, stob("00000000 00000000 00010000 00000000"), p.HKnight, 3)

	id = 30
	testCase(&id, stob("00000000 00000000 01000000 00000000"), p.HSilver, 0)
	testCase(&id, stob("00000000 00000001 00000000 00000000"), p.HSilver, 3)

	id = 40
	testCase(&id, stob("00000000 00000100 00000000 00000000"), p.HGold, 0)
	testCase(&id, stob("00000000 00010000 00000000 00000000"), p.HGold, 3)

	id = 50
	testCase(&id, stob("00000000 01000000 00000000 00000000"), p.HBishop, 0)
	testCase(&id, stob("00000000 10000000 00000000 00000000"), p.HBishop, 1)

	id = 60
	testCase(&id, stob("00000010 00000000 00000000 00000000"), p.HRook, 0)
	testCase(&id, stob("00000100 00000000 00000000 00000000"), p.HRook, 1)

}
