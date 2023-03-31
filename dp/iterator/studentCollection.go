package iterator

type StudentCollection struct {
	students []*Student
}

func (t *StudentCollection) createIterator() *StudentIterator {
	return &StudentIterator{
		students: t.students,
	}
}
