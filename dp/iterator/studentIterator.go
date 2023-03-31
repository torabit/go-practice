package iterator

type StudentIterator struct {
	students []*Student
	index    int
}

func (s *StudentIterator) hasNext() bool {
	if s.index < len(s.students) {
		return true
	}
	return false
}

func (s *StudentIterator) next() *Student {
	if s.hasNext() {
		student := s.students[s.index]
		s.index++
		return student
	}
	return nil
}
