package iterator

// item
type Record struct {
	Id     int
	Title  string
	Artist string
}

// item aggregator
type RecordShelf struct {
	records  []*Record
	iterator *RecordShelfIterator
}

func NewRecordShelf() *RecordShelf {
	r := &RecordShelf{
		iterator: &RecordShelfIterator{},
	}
	r.iterator.recordShelf = r
	return r
}

func (r *RecordShelf) GetItemAt(index int) *Record {
	if r.GetSize() <= index {
		return nil
	}
	return r.records[index]
}

func (r *RecordShelf) HasNext() bool {
	return r.iterator.HasNext()
}

func (r *RecordShelf) Next() *Record {
	return r.iterator.Next()
}

func (r *RecordShelf) Append(record *Record) {
	r.records = append(r.records, record)
}

func (r *RecordShelf) GetSize() int {
	return int(len(r.records))
}

// iterator
type RecordShelfIterator struct {
	recordShelf *RecordShelf
	last        int
}

func (r *RecordShelfIterator) HasNext() bool {
	if r.last < r.recordShelf.GetSize() {
		return true
	}
	return false
}

func (r *RecordShelfIterator) Next() *Record {
	item := r.recordShelf.GetItemAt(r.last)
	r.last++
	return item
}
