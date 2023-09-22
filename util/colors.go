package util

type Color int

func (c Color) String() string {
	switch c {
	case Red:
		return "RED"
	case Green:
		return "GREEN"
	case Blue:
		return "BLUE"
	default:
		return "UNKNOWN"
	}
}

const (
	Red Color = iota
	Green
	Blue
)

func (c Color) GetColor(num int) string {
	switch num {
	case 1:
		return Red.String()
	case 2:
		return Green.String()
	case 3:
		return Blue.String()
	default:
		return "UNKNOWN"
	}
}
