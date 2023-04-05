package flyweight

type TDress struct {
	color string
}

func (t *TDress) getColor() string {
	return t.color
}

func newTDress() *TDress {
	return &TDress{color: "red"}
}
