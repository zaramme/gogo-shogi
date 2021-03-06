package common

import (
//	"fmt"
)

const (
	CacheLineSize = 64
)

type Key uint64

var bitTable = [64]int{
	63, 30, 3, 32, 25, 41, 22, 33, 15, 50, 42, 13, 11, 53, 19, 34, 61, 29, 2,
	51, 21, 43, 45, 10, 18, 47, 1, 54, 9, 57, 0, 35, 62, 31, 40, 4, 49, 5, 52,
	26, 60, 6, 23, 44, 46, 27, 56, 16, 7, 39, 48, 24, 59, 14, 12, 55, 38, 28,
	58, 20, 37, 17, 36, 8}

func FirstOneFromLSB(b uint64) int {
	tmp := b ^ b - 1
	old := (tmp & 0xffffffff) ^ (tmp >> 32)
	return bitTable[(old * 0x783a9b23 >> 26)]
}

// 超絶遅いコードなので後で書き換えること。
func FirstOneFromMSB(b uint64) int {
	for i := uint64(63); 0 <= i; i-- {
		b >>= 1
		if b&1 == 1 {
			return int(63 - i)
		}
	}
	return 0
}

// 任意の値に含まれる1のビットの数を数える。
func Count1s(x uint64) int {
	// ２ビットずつのグループに分け、その２ビットの含まれる１の数の値に変換
	// 00 -> 00 , 01 or 10 -> 01 11 -> 10
	x = x - ((x >> 1) & 0x5555555555555555)
	// ４ビットずつのグループに分け、前半２ビット＋後半２ビットで加算
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	// ８ビットずつのグループに分け、前半４ビット＋後半４ビットで加算
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	// １６ビットずつのグループに分け、前半８ビット＋後半８ビットで加算（以下ビット数を２倍にして繰り返し）
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)

	//末尾７ビットの数値でマスク(最大値が６４のため)
	return (int(x) & 0x0000007f)
}
