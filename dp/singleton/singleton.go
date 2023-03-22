package singleton

type Singleton struct {
	input int
}

func (s *Singleton) getInstance(input int) *Singleton {
	s.input = input
	return s
}

type SingletonClass struct {
	Singleton
}
