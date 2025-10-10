package models

type VolunteerRequest struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"required,gte=18"`
	Phone string `json:"phone" binding:"required,e164"`
	CNIC  string `json:"cnic" binding:"required,len=13"`

	Role         int    `json:"role" binding:"required,oneof=1 2 3 4"`
	City         string `json:"city" binding:"required"`
	Constituency string `json:"constituency" binding:"required"`
	Active       bool   `json:"active"`
}

type CreateVolunteerRequest struct {
	VolunteerRequest
}

type UpdateVolunteerRequest struct {
	VolunteerRequest
}
