package composite

type File struct {
	name string
	size int
}

func (f *File) getName() string {
	return f.name
}

func (f *File) getSize() int {
	return f.size
}
