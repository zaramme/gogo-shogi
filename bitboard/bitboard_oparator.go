package bitboard

func (b BitBoard) Inverse() BitBoard {
	b[0] = -b[0] - 1
	b[1] = -b[1] - 1
	return b
}

// &=
func (b1 BitBoard) AndAssign(b2 BitBoard) BitBoard {
	//高速化処理
	// #if defined (HAVE_SSE2) || defined (HAVE_SSE4)
	// 		_mm_store_si128(&this->m_, _mm_and_si128(this->m_, rhs.m_));
	// #else
	b1[0] &= b2[0]
	b1[1] &= b2[1]
	return b1
}

// |=
func (b1 BitBoard) OrAssign(b2 BitBoard) BitBoard {
	b1[0] |= b2[0]
	b1[1] |= b2[1]
	return b1
}

// ^=
func (b1 BitBoard) XorAssign(b2 BitBoard) BitBoard {
	b1[0] ^= b2[0]
	b1[1] ^= b2[1]
	return b1
}

// <<
func (b BitBoard) LeftShift(num uint64) BitBoard {
	b[0] <<= num
	b[1] <<= num
	return b
}

// >>
func (b BitBoard) RightShift(num uint64) BitBoard {
	b[0] >>= num
	b[1] >>= num
	return b
}

// & | ^ ==, !=記号はすべて代入演算子で代用とする。
// func and(b1, b2 BitBoard) BitBoard {
// 	return b1.andAssign(b2)
// }

// func or(b1, b2 BitBoard) BitBoard {
// 	return b1.orAssign(b2)
// }

// func xor(b1, b2 BitBoard) BitBoard {
// 	return b1.xorAssign(b2)
// }

// func equal(b1, b2 BitBoard) bool {
// 	return b1[0] == b2[0] && b1[1] == b2[1]
// }
