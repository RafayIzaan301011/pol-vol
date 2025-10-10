package models

type VolunteerResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
	CNIC  string `json:"cnic"`

	City         string `json:"city"`
	Constituency string `json:"constituency"`
	Role         string `json:"role"`
	Active       bool   `json:"active"`
}
