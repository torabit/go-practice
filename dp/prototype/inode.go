package prototype

type Inode interface {
	print(string) string
	clone() Inode
}
