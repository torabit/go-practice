package factorymethod

type iHero interface {
	getName() string
	getRole() Role
	setName(name string)
	setRole(role Role)
}
