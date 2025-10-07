package domain

import "encoding/json"

type Role int

const (
	RoleTier1 Role = iota
	RoleTier2
	RoleLeader
	RoleAdmin
)

func (r Role) String() string {
	roles := [...]string{
		"Tier 1",
		"Tier 2",
		"Leader",
		"Admin",
	}

	if r < RoleTier1 || r > RoleAdmin {
		return "Unknown"
	}
	return roles[r]
}

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

type Volunteer struct {
	Name         string
	Age          int
	Phone        string
	CNIC         string
	City         string
	Constituency string
	Role         Role
	IsActive     bool
}
