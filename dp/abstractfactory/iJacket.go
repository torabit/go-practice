package abstractfactory

type IJacket interface {
	setLogo(logo string)
	setSize(size int)
	setPrice(price int)
	getLogo() string
	getSize() int
	getPrice() int
}

type Jacket struct {
	logo  string
	size  int
	price int
}

func (s *Jacket) setLogo(logo string) {
	s.logo = logo
}

func (s *Jacket) setSize(size int) {
	s.size = size
}

func (s *Jacket) setPrice(price int) {
	s.price = price
}

func (s *Jacket) getLogo() string {
	return s.logo
}

func (s *Jacket) getSize() int {
	return s.size
}

func (s *Jacket) getPrice() int {
	return s.price
}
