package visitor

type Shape int

const (
	SQUARE Shape = iota
	CIRCLE
	RECTANGLE
)

type IShape interface {
	getType() string
	accept(Visitor)
}
