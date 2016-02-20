package color

import (
	"testing"
)

func Test_opposite(t *testing.T) {

	c := White
	o := c.Opposite()

	if o != Black {
		t.Errorf("opposite変換に失敗しました。入力 = %d, 出力 = %d ", c)
	}

	c = Black
	o = c.Opposite()
	if o != White {
		t.Errorf("opposite変換に失敗しました。入力 = %d, 出力 = %d ", c)
	}

}
