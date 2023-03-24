package builder

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

type builderType int

const (
	normal builderType = iota
	igloo
)

func getBuilder(b builderType) IBuilder {
	if b == normal {
		return newNormalBuilder()
	}

	if b == igloo {
		return newIglooBuilder()
	}
	return nil
}
