package math

type math struct {
	a int
	b int
}

func NewMath(a, b int) *math {
	return &math{
		a: a,
		b: b,
	}
}

func (m *math) Sum() int {
	return m.a + m.b
}
