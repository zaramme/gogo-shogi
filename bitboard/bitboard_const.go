package bitboard

import (
	c "github.com/zaramme/gogo-shogi/color"
	s "github.com/zaramme/gogo-shogi/square"
)

var maskBB = [s.SquareNum]BitBoard{
	BitBoard{1 << 0, 0}, //0,I9
	BitBoard{1 << 1, 0},
	BitBoard{1 << 2, 0},
	BitBoard{1 << 3, 0},
	BitBoard{1 << 4, 0},
	BitBoard{1 << 5, 0},
	BitBoard{1 << 6, 0},
	BitBoard{1 << 7, 0},
	BitBoard{1 << 8, 0},
	BitBoard{1 << 9, 0},
	BitBoard{1 << 10, 0},
	BitBoard{1 << 11, 0},
	BitBoard{1 << 12, 0},
	BitBoard{1 << 13, 0},
	BitBoard{1 << 14, 0},
	BitBoard{1 << 15, 0},
	BitBoard{1 << 16, 0},
	BitBoard{1 << 17, 0},
	BitBoard{1 << 18, 0},
	BitBoard{1 << 19, 0},
	BitBoard{1 << 20, 0},
	BitBoard{1 << 21, 0},
	BitBoard{1 << 22, 0},
	BitBoard{1 << 23, 0},
	BitBoard{1 << 24, 0},
	BitBoard{1 << 25, 0},
	BitBoard{1 << 26, 0},
	BitBoard{1 << 27, 0},
	BitBoard{1 << 28, 0},
	BitBoard{1 << 29, 0},
	BitBoard{1 << 30, 0},
	BitBoard{1 << 31, 0},
	BitBoard{1 << 32, 0},
	BitBoard{1 << 33, 0},
	BitBoard{1 << 34, 0},
	BitBoard{1 << 35, 0},
	BitBoard{1 << 36, 0},
	BitBoard{1 << 37, 0},
	BitBoard{1 << 38, 0},
	BitBoard{1 << 39, 0},
	BitBoard{1 << 40, 0},
	BitBoard{1 << 41, 0},
	BitBoard{1 << 42, 0},
	BitBoard{1 << 43, 0},
	BitBoard{1 << 44, 0},
	BitBoard{1 << 45, 0},
	BitBoard{1 << 46, 0},
	BitBoard{1 << 47, 0},
	BitBoard{1 << 48, 0},
	BitBoard{1 << 49, 0},
	BitBoard{1 << 50, 0},
	BitBoard{1 << 51, 0},
	BitBoard{1 << 52, 0},
	BitBoard{1 << 53, 0},
	BitBoard{1 << 54, 0},
	BitBoard{1 << 55, 0},
	BitBoard{1 << 56, 0},
	BitBoard{1 << 57, 0},
	BitBoard{1 << 58, 0},
	BitBoard{1 << 59, 0},
	BitBoard{1 << 60, 0},
	BitBoard{1 << 61, 0},
	BitBoard{1 << 62, 0}, // 62,C1
	BitBoard{0, 1 << 0},  // 63,B9
	BitBoard{0, 1 << 1},
	BitBoard{0, 1 << 2},
	BitBoard{0, 1 << 3},
	BitBoard{0, 1 << 4},
	BitBoard{0, 1 << 5},
	BitBoard{0, 1 << 6},
	BitBoard{0, 1 << 7},
	BitBoard{0, 1 << 8},
	BitBoard{0, 1 << 9},
	BitBoard{0, 1 << 10},
	BitBoard{0, 1 << 11},
	BitBoard{0, 1 << 12},
	BitBoard{0, 1 << 13},
	BitBoard{0, 1 << 14},
	BitBoard{0, 1 << 15},
	BitBoard{0, 1 << 16},
	BitBoard{0, 1 << 17}, // 80, A1
}

// 各マスのrookが利きを調べる必要があるマスの数
var rookBlockBits = [s.SquareNum]int{
	14, 13, 13, 13, 13, 13, 13, 13, 14,
	13, 12, 12, 12, 12, 12, 12, 12, 13,
	13, 12, 12, 12, 12, 12, 12, 12, 13,
	13, 12, 12, 12, 12, 12, 12, 12, 13,
	13, 12, 12, 12, 12, 12, 12, 12, 13,
	13, 12, 12, 12, 12, 12, 12, 12, 13,
	13, 12, 12, 12, 12, 12, 12, 12, 13,
	13, 12, 12, 12, 12, 12, 12, 12, 13,
	14, 13, 13, 13, 13, 13, 13, 13, 14,
}

// 各マスのbishopが利きを調べる必要があるマスの数
var bishopBlockBits = [s.SquareNum]int{
	7, 6, 6, 6, 6, 6, 6, 6, 7,
	6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 8, 8, 8, 8, 8, 6, 6,
	6, 6, 8, 10, 10, 10, 8, 6, 6,
	6, 6, 8, 10, 12, 10, 8, 6, 6,
	6, 6, 8, 10, 10, 10, 8, 6, 6,
	6, 6, 8, 8, 8, 8, 8, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6,
	7, 6, 6, 6, 6, 6, 6, 6, 7,
}

// Magic BitBoard で利きを求める際のシフト量
var rookShiftBits = [s.SquareNum]int{
	50, 51, 51, 51, 51, 51, 51, 51, 50,
	51, 52, 52, 52, 52, 52, 52, 52, 50, // [17]: 51 -> 50
	51, 52, 52, 52, 52, 52, 52, 52, 51,
	51, 52, 52, 52, 52, 52, 52, 52, 51,
	51, 52, 52, 52, 52, 52, 52, 52, 51,
	51, 52, 52, 52, 52, 52, 52, 52, 50, // [53]: 51 -> 50
	51, 52, 52, 52, 52, 52, 52, 52, 51,
	51, 52, 52, 52, 52, 52, 52, 52, 51,
	50, 51, 51, 51, 51, 51, 51, 51, 50,
}

// Magic BitBoard で利きを求める際のシフト量
var bishopShiftBits = [s.SquareNum]int{
	57, 58, 58, 58, 58, 58, 58, 58, 57,
	58, 58, 58, 58, 58, 58, 58, 58, 58,
	58, 58, 56, 56, 56, 56, 56, 58, 58,
	58, 58, 56, 54, 54, 54, 56, 58, 58,
	58, 58, 56, 54, 52, 54, 56, 58, 58,
	58, 58, 56, 54, 54, 54, 56, 58, 58,
	58, 58, 56, 56, 56, 56, 56, 58, 58,
	58, 58, 58, 58, 58, 58, 58, 58, 58,
	57, 58, 58, 58, 58, 58, 58, 58, 57,
}

var RookMagic = [s.SquareNum]uint64{
	0x140000400809300, 0x1320000902000240, 0x8001910c008180,
	0x40020004401040, 0x40010000d01120, 0x80048020084050,
	0x40004000080228, 0x400440000a2a0a, 0x40003101010102,
	0x80c4200012108100, 0x4010c00204000c01, 0x220400103250002,
	0x2600200004001, 0x40200052400020, 0xc00100020020008,
	0x9080201000200004, 0x2200201000080004, 0x80804c00202001,
	0x45383000009100, 0x30002800020040, 0x40104000988084,
	0x108001000800415, 0x14005000400009, 0xd21001001c00045,
	0xc0003000200024, 0x40003000280004, 0x40021000091102,
	0x2008a20408000d00, 0x2000100084010040, 0x144080008008001,
	0x50102400100026a2, 0x1040020008001010, 0x1200200028005010,
	0x4280030030020898, 0x480081410011004, 0x34000040800110a,
	0x101000010c0021, 0x9210800080082, 0x6100002000400a7,
	0xa2240800900800c0, 0x9220082001000801, 0x1040008001140030,
	0x40002220040008, 0x28000124008010c, 0x40008404940002,
	0x40040800010200, 0x90000809002100, 0x2800080001000201,
	0x1400020001000201, 0x180081014018004, 0x1100008000400201,
	0x80004000200201, 0x420800010000201, 0x2841c00080200209,
	0x120002401040001, 0x14510000101000b, 0x40080000808001,
	0x834000188048001, 0x4001210000800205, 0x4889a8007400201,
	0x2080044080200062, 0x80004002861002, 0xc00842049024,
	0x8040000202020011, 0x400404002c0100, 0x2080028202000102,
	0x8100040800590224, 0x2040009004800010, 0x40045000400408,
	0x2200240020802008, 0x4080042002200204, 0x4000b0000a00a2,
	0xa600000810100, 0x1410000d001180, 0x2200101001080,
	0x100020014104e120, 0x2407200100004810, 0x80144000a0845050,
	0x1000200060030c18, 0x4004200020010102, 0x140600021010302,
}

var BishopMagic = [s.SquareNum]uint64{
	0x20101042c8200428, 0x840240380102, 0x800800c018108251,
	0x82428010301000, 0x481008201000040, 0x8081020420880800,
	0x804222110000, 0xe28301400850, 0x2010221420800810,
	0x2600010028801824, 0x8048102102002, 0x4000248100240402,
	0x49200200428a2108, 0x460904020844, 0x2001401020830200,
	0x1009008120, 0x4804064008208004, 0x4406000240300ca0,
	0x222001400803220, 0x226068400182094, 0x95208402010d0104,
	0x4000807500108102, 0xc000200080500500, 0x5211000304038020,
	0x1108100180400820, 0x10001280a8a21040, 0x100004809408a210,
	0x202300002041112, 0x4040a8000460408, 0x204020021040201,
	0x8120013180404, 0xa28400800d020104, 0x200c201000604080,
	0x1082004000109408, 0x100021c00c410408, 0x880820905004c801,
	0x1054064080004120, 0x30c0a0224001030, 0x300060100040821,
	0x51200801020c006, 0x2100040042802801, 0x481000820401002,
	0x40408a0450000801, 0x810104200000a2, 0x281102102108408,
	0x804020040280021, 0x2420401200220040, 0x80010144080c402,
	0x80104400800002, 0x1009048080400081, 0x100082000201008c,
	0x10001008080009, 0x2a5006b80080004, 0xc6288018200c2884,
	0x108100104200a000, 0x141002030814048, 0x200204080010808,
	0x200004013922002, 0x2200000020050815, 0x2011010400040800,
	0x1020040004220200, 0x944020104840081, 0x6080a080801c044a,
	0x2088400811008020, 0xc40aa04208070, 0x4100800440900220,
	0x48112050, 0x818200d062012a10, 0x402008404508302,
	0x100020101002, 0x20040420504912, 0x2004008118814,
	0x1000810650084024, 0x1002a03002408804, 0x2104294801181420,
	0x841080240500812, 0x4406009000004884, 0x80082004012412,
	0x80090880808183, 0x300120020400410, 0x21a090100822002,
}

// 以下は初期化処理の際に値を算出、格納する
var RookAttack [495616]BitBoard

var RookAttackIndex [s.SquareNum]BitBoard

var RookBlockMask [s.SquareNum]BitBoard
var BishopAttack [20224]BitBoard
var BishopAttackIndex [s.SquareNum]uint64
var BishopBlockMask [s.SquareNum]BitBoard
var LanceAttack [c.ColorNum][s.SquareNum][128]BitBoard

var KingAttack [s.SquareNum]BitBoard
var GoldAttack [c.ColorNum][s.SquareNum]BitBoard
var SilverAttack [c.ColorNum][s.SquareNum]BitBoard
var KnightAttack [c.ColorNum][s.SquareNum]BitBoard
var PawnAttack [c.ColorNum][s.SquareNum]BitBoard

var BetweenBB [s.SquareNum][s.SquareNum]BitBoard

var RookAttackToEdge [s.SquareNum]BitBoard
var BishopAttackToEdge [s.SquareNum]BitBoard
var LanceAttackToEdge [c.ColorNum][s.SquareNum]BitBoard

var GoldCheckTable [c.ColorNum][s.SquareNum]BitBoard
var SilverCheckTable [c.ColorNum][s.SquareNum]BitBoard
var KnightCheckTable [c.ColorNum][s.SquareNum]BitBoard
var LanceCheckTable [c.ColorNum][s.SquareNum]BitBoard

// indexを求めるために必要なビットシフト量
var Slide = [s.SquareNum]uint{
	1, 1, 1, 1, 1, 1, 1, 1, 1,
	10, 10, 10, 10, 10, 10, 10, 10, 10,
	19, 19, 19, 19, 19, 19, 19, 19, 19,
	28, 28, 28, 28, 28, 28, 28, 28, 28,
	37, 37, 37, 37, 37, 37, 37, 37, 37,
	46, 46, 46, 46, 46, 46, 46, 46, 46,
	55, 55, 55, 55, 55, 55, 55, 55, 55,
	1, 1, 1, 1, 1, 1, 1, 1, 1,
	10, 10, 10, 10, 10, 10, 10, 10, 10,
}

var FileIMask BitBoard = BitBoard{0x1ff << (9 * 0), 0}
var FileHMask BitBoard = BitBoard{0x1ff << (9 * 1), 0}
var FileGMask BitBoard = BitBoard{0x1ff << (9 * 2), 0}
var FileFMask BitBoard = BitBoard{0x1ff << (9 * 3), 0}
var FileEMask BitBoard = BitBoard{0x1ff << (9 * 4), 0}
var FileDMask BitBoard = BitBoard{0x1ff << (9 * 5), 0}
var FileCMask BitBoard = BitBoard{0x1ff << (9 * 6), 0}
var FileBMask BitBoard = BitBoard{0, 0x1ff << (9 * 0)}
var FileAMask BitBoard = BitBoard{0, 0x1ff << (9 * 1)}

var FileMask = [s.FileNum]BitBoard{
	FileIMask, FileHMask, FileGMask, FileFMask, FileEMask, FileDMask,
	FileCMask, FileBMask, FileAMask,
}

var Rank9Mask BitBoard = BitBoard{0x40201008040201 << 0, 0x201 << 0}
var Rank8Mask BitBoard = BitBoard{0x40201008040201 << 1, 0x201 << 1}
var Rank7Mask BitBoard = BitBoard{0x40201008040201 << 2, 0x201 << 2}
var Rank6Mask BitBoard = BitBoard{0x40201008040201 << 3, 0x201 << 3}
var Rank5Mask BitBoard = BitBoard{0x40201008040201 << 4, 0x201 << 4}
var Rank4Mask BitBoard = BitBoard{0x40201008040201 << 5, 0x201 << 5}
var Rank3Mask BitBoard = BitBoard{0x40201008040201 << 6, 0x201 << 6}
var Rank2Mask BitBoard = BitBoard{0x40201008040201 << 7, 0x201 << 7}
var Rank1Mask BitBoard = BitBoard{0x40201008040201 << 8, 0x201 << 8}

var RankMusk = [s.RankNum]BitBoard{
	Rank9Mask, Rank8Mask, Rank7Mask, Rank6Mask, Rank5Mask,
	Rank4Mask, Rank3Mask, Rank2Mask, Rank1Mask,
}

var InFrontOfRank9Black BitBoard = NewBitBoardAllZero()
var InFrontOfRank8Black BitBoard = Rank9Mask
var InFrontOfRank7Black BitBoard = InFrontOfRank8Black.OrAssign(Rank8Mask)
var InFrontOfRank6Black BitBoard = InFrontOfRank7Black.OrAssign(Rank7Mask)
var InFrontOfRank5Black BitBoard = InFrontOfRank6Black.OrAssign(Rank6Mask)
var InFrontOfRank4Black BitBoard = InFrontOfRank5Black.OrAssign(Rank5Mask)
var InFrontOfRank3Black BitBoard = InFrontOfRank4Black.OrAssign(Rank6Mask)
var InFrontOfRank2Black BitBoard = InFrontOfRank3Black.OrAssign(Rank7Mask)
var InFrontOfRank1Black BitBoard = InFrontOfRank2Black.OrAssign(Rank8Mask)

var InFrontOfRank1White BitBoard = NewBitBoardAllZero()
var InFrontOfRank2White BitBoard = Rank1Mask
var InFrontOfRank3White BitBoard = InFrontOfRank2White.OrAssign(Rank2Mask)
var InFrontOfRank4White BitBoard = InFrontOfRank3White.OrAssign(Rank3Mask)
var InFrontOfRank5White BitBoard = InFrontOfRank4White.OrAssign(Rank4Mask)
var InFrontOfRank6White BitBoard = InFrontOfRank5White.OrAssign(Rank5Mask)
var InFrontOfRank7White BitBoard = InFrontOfRank6White.OrAssign(Rank6Mask)
var InFrontOfRank8White BitBoard = InFrontOfRank7White.OrAssign(Rank7Mask)
var InFrontOfRank9White BitBoard = InFrontOfRank8White.OrAssign(Rank8Mask)

var inFrontMask = [c.ColorNum][s.RankNum]BitBoard{
	{InFrontOfRank9Black, InFrontOfRank8Black, InFrontOfRank7Black, InFrontOfRank6Black, InFrontOfRank5Black, InFrontOfRank4Black, InFrontOfRank3Black, InFrontOfRank2Black, InFrontOfRank1Black},
	{InFrontOfRank9White, InFrontOfRank8White, InFrontOfRank7White, InFrontOfRank6White, InFrontOfRank5White, InFrontOfRank4White, InFrontOfRank3White, InFrontOfRank2White, InFrontOfRank1White},
}
