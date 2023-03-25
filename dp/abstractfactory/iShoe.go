package abstractfactory

type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	setPrice(price int)
	getLogo() string
	getSize() int
	getPrice() int
}

type Shoe struct {
	logo  string
	size  int
	price int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) setPrice(price int) {
	s.price = price
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) getSize() int {
	return s.size
}

func (s *Shoe) getPrice() int {
	return s.price
}
