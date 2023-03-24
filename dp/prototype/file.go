package prototype

type File struct {
	name string
}

func (f *File) print(indentation string) string {
	return indentation + f.name + "\n"
}

func (f *File) clone() Inode {
	return &File{name: f.name + "_clone"}
}
