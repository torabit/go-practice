package iterator

import (
	"reflect"
	"testing"
)

func TestStudentCollection(t *testing.T) {
	student1 := &Student{name: "横山 貴大", sex: Man}
	student2 := &Student{name: "相沢 萌", sex: Woman}

	t.Run("case: hasNext method, has student", func(t *testing.T) {
		studentCollection := &StudentCollection{
			students: []*Student{student1},
		}
		iterator := studentCollection.createIterator()

		if iterator.hasNext() != true {
			t.Fatalf("\nwant: %+v\ngot: %+v\n", true, iterator.hasNext())
		}

	})

	t.Run("case: hasNext method, has no student", func(t *testing.T) {
		studentCollection := &StudentCollection{}
		iterator := studentCollection.createIterator()

		if iterator.hasNext() != false {
			t.Fatalf("\nwant: %+v\ngot: %+v\n", false, iterator.hasNext())
		}
	})

	t.Run("case: Next method, does received student", func(t *testing.T) {
		studentCollection := &StudentCollection{
			students: []*Student{student1},
		}
		iterator := studentCollection.createIterator()
		student := iterator.next()

		if !reflect.DeepEqual(student1, student) {
			t.Fatalf("\nwant: %+v\ngot: %+v\n", student1, student)
		}
	})

	t.Run("case: Next method, does not received student", func(t *testing.T) {
		studentCollection := &StudentCollection{}
		iterator := studentCollection.createIterator()
		student := iterator.next()

		if student != nil {
			t.Fatalf("\nwant: %+v\ngot: %+v\n", nil, student)
		}
	})

	t.Run("case: for check out put", func(t *testing.T) {
		expectedNames := []string{"横山 貴大", "相沢 萌"}

		studentCollection := &StudentCollection{
			students: []*Student{student1, student2},
		}
		iterator := studentCollection.createIterator()
		var studentNames []string
		for iterator.hasNext() {
			name := iterator.next().name
			studentNames = append(studentNames, name)
		}

		if !reflect.DeepEqual(studentNames, expectedNames) {
			t.Fatalf("\nwant: %+v\ngot: %+v", studentNames, expectedNames)
		}
	})
}
