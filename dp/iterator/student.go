package iterator

type Sex int

const (
	Woman Sex = iota
	Man
)

type Student struct {
	name string
	sex  Sex // 0:woman 1:man
}
