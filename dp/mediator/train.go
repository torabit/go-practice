package mediator

type ITrain interface {
	arrive()
	depart()
	permitArrival()
}
