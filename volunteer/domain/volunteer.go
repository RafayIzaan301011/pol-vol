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
	ID           string
	Name         string
	Age          int
	Phone        string
	CNIC         string
	City         string
	Constituency string
	Role         Role
	IsActive     bool
}

type VolunteerFilter struct {
	Age          *int
	Phone        *string
	Role         *int
	City         *string
	Constituency *string
	IsActive     *bool

	Limit  *int
	Offset *int
}

type VolunteerSummary struct {
	TotalCount    int
	ActiveCount   int
	InactiveCount int
}
type VolunteerStats struct {
	ByCity         map[string]int
	ByConstituency map[string]int
	ByRole         map[string]int
}

type VolunteerStatsResponse struct {
	TotalCount     int
	ActiveCount    int
	InactiveCount  int
	ByCity         map[string]int
	ByConstituency map[string]int
	ByRole         map[string]int
}

type ListVolunteers struct {
	Volunteers []Volunteer
	Summary    VolunteerSummary
	Stats      VolunteerStats
}
