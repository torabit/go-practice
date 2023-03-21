package factorymethod

type Role int

const (
	Tank Role = iota
	Damage
	Support
)

type Hero struct {
	name string
	role Role
}

func (h *Hero) setName(name string) {
	h.name = name
}

func (h *Hero) setRole(role Role) {
	h.role = role
}

func (h *Hero) getName() string {
	return h.name
}

func (h *Hero) getRole() Role {
	return h.role
}
