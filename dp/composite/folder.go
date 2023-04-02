package composite

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) getName() string {
	return f.name
}

func (f *Folder) getSize() int {
	var size int
	for _, composite := range f.components {
		size += composite.getSize()
	}
	return size
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}
