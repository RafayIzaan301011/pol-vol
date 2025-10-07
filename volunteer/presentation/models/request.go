package models

type Volunteer struct {
	Name         string `json:"name" binding:"required"`
	Age          int    `json:"age" binding:"required,gte=18"`
	Phone        string `json:"phone" binding:"required,e164"`
	CNIC         string `json:"cnic" binding:"required,len=13"`
	Role         string `json:"role" binding:"required,oneof=admin user guest"`
	City         string `json:"city" binding:"required"`
	Constituency string `json:"constituency" binding:"required"`
	Active       bool   `json:"active"`
}

type CreateVolunteerRequest struct {
	Volunteer
}

type UpdateVolunteerRequest struct {
	Volunteer
}
