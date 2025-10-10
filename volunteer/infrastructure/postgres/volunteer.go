package postgres

import (
	"pvms/volunteer/domain"

	"gorm.io/gorm"
)

// VolunteerPostgres implements the ReadRepositories and WriteRepositories interfaces
type VolunteerPostgres struct {
	db *gorm.DB
}

// NewVolunteerPostgres creates new instance of VolunteerPostgres
func NewVolunteerPostgres(db *gorm.DB) *VolunteerPostgres {
	return &VolunteerPostgres{db: db}
}

// VolunteerRow represents the volunteer table in the database
type VolunteerRow struct {
	ID           string
	Name         string
	Age          int
	Phone        string
	CNIC         string
	City         string
	Constituency string
	Role         int
	Active       bool
}

// TableName specifies the Volunteer Table name for GORM
func (VolunteerRow) TableName() string {
	return "volunteers"
}

func (v *VolunteerRow) toDomain() *domain.Volunteer {
	return &domain.Volunteer{
		ID:           v.ID,
		Name:         v.Name,
		Age:          v.Age,
		Phone:        v.Phone,
		CNIC:         v.CNIC,
		City:         v.City,
		Constituency: v.Constituency,
		Role:         domain.Role(v.Role),
		IsActive:     v.Active,
	}
}

func (v *VolunteerRow) fromDomain(vd *domain.Volunteer) {
	v.ID = vd.ID
	v.Name = vd.Name
	v.Age = vd.Age
	v.Phone = vd.Phone
	v.CNIC = vd.CNIC
	v.City = vd.City
	v.Constituency = vd.Constituency
	v.Role = int(vd.Role)
	v.Active = vd.IsActive
}
