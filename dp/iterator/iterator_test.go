package iterator

import "testing"

func TestRecordShelf(t *testing.T) {

	sampleRecord := &Record{Id: 1, Title: "On & On", Artist: "Erykha Badu"}

	t.Run("case Append method", func(t *testing.T) {
		r := NewRecordShelf()
		r.Append(sampleRecord)
		if r.Next() != sampleRecord {
			t.Fatal("Record could not append")
		}
	})

	t.Run("case GetSize method", func(t *testing.T) {
		r := NewRecordShelf()
		r.Append(sampleRecord)
		r.Append(sampleRecord)

		if r.GetSize() != 2 {
			t.Fatal("There is a difference in the number of records")
		}
	})

	t.Run("case GetItemAt method, has no items", func(t *testing.T) {
		r := NewRecordShelf()
		if r.GetItemAt(0) != nil {
			t.Fatal("Record Shelf has items")
		}
	})

	t.Run("case GetItemAt method, has items", func(t *testing.T) {
		r := NewRecordShelf()
		r.Append(sampleRecord)
		if r.GetItemAt(0) == nil {
			t.Fatal("Record Shelf has no items")
		}
	})

}

func TestRecordShelfIterator(t *testing.T) {

	sampleRecord := &Record{Id: 1, Title: "On & On", Artist: "Erykha Badu"}
	sampleRecord2 := &Record{Id: 2, Title: "FINAL DISTANCE", Artist: "Hikaru Utada"}

	t.Run("case HasNext method, has next", func(t *testing.T) {
		r := NewRecordShelf()
		r.Append(sampleRecord)
		if r.HasNext() != true {
			t.Fatal("Record has no next")
		}
	})

	t.Run("case HasNext method, has no next", func(t *testing.T) {
		r := NewRecordShelf()
		if r.HasNext() != false {
			t.Fatal("Record has next", r.Next())
		}
	})

	t.Run("case Next method", func(t *testing.T) {
		r := NewRecordShelf()
		r.Append(sampleRecord)
		r.Append(sampleRecord2)
		if r.Next() != sampleRecord {
			t.Fatal("The order of the records is different")
		}
	})
}
