package composite

import "testing"

func TestComposite(t *testing.T) {
	file1 := &File{name: "vi", size: 100}
	file2 := &File{name: "lat", size: 121}
	file3 := &File{name: "so", size: 973}

	folder1 := &Folder{name: "memo"}
	folder2 := &Folder{name: "next"}

	folder1.add(file1)
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	if !(folder2.getSize() == 1194) {
		t.Fatalf("\nwant: %+v\ngot: %+v\n", 1194, folder2.getSize())
	}
}
