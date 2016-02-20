package common

import (
	"strconv"
	"strings"
	"testing"
)

func stob(s string) uint64 {
	// split -> join で半角スペースを削除
	sa := strings.Split(s, " ")
	sb := strings.Join(sa, "")
	v, _ := strconv.ParseInt(sb, 2, 64)
	return uint64(v)
}

func Test_Count1s(t *testing.T) {

	testCase := func(x uint64, expect int) {
		v := Count1s(x)
		if v != expect {
			t.Errorf("error input=%x, actual=%d, expect = %d", x, expect, v)
		}

	}

	var x uint64
	x = 0x0000000000000001
	testCase(x, 1)
	x = 0x0000000000000101
	testCase(x, 2)
	x = 0x000000000000000f
	testCase(x, 4)
	x = 0x000000000003030f
	testCase(x, 8)
	x = 0x8888888888888888
	testCase(x, 16)
	x = 0x3333333333333333
	testCase(x, 32)
	x = 0xFFFFFFFFFFFFFFFF
	testCase(x, 64)

}
