package structure

type matrix struct {
	x int
	y int
}

func NewMatric(x, y int) *matrix {
	if x < 0 || y < 0 {
		return nil
	}

	return &matrix{x, y}
}
