package factorymethod

type IHero interface {
	getName() string
	getRole() Role
	setName(name string)
	setRole(role Role)
}
