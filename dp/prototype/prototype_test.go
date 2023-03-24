package prototype

import "testing"

func TestPrototype(t *testing.T) {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}
	t.Run("case oridinal node", func(t *testing.T) {
		expectHierarchy := "  Folder2\n    Folder1\n        File1\n    File2\n    File3\n"
		hierarchy := folder2.print("  ")
		if hierarchy != expectHierarchy {
			t.Fatalf("The hierarchy is not same")
		}
	})

	t.Run("case clone node", func(t *testing.T) {
		expectHierarchy := "  Folder2_clone\n    Folder1_clone\n        File1_clone\n    File2_clone\n    File3_clone\n"
		clone := folder2.clone()
		hierarchy := clone.print("  ")
		if hierarchy != expectHierarchy {
			t.Fatal("The clone hierarchy is not same")
		}
	})
}
