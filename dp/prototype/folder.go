package prototype

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indentation string) string {
	var hierarchy string
	hierarchy += indentation + f.name + "\n"
	for _, i := range f.children {
		hierarchy += i.print(indentation + indentation)
	}
	return hierarchy
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
