package color

// Color is 先手・後手を表現する型
type Color int

// 先手・後手を表現する
const (
	Black    Color = 0
	White    Color = 1
	ColorNum Color = 3 // as sentinel
)

func (c Color) Opposite() Color {
	return Color(c ^ 1)
}
