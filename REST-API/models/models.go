package models

type Employee struct {
	SSN          int    `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	DepartmentID int    `gorm:"not null"` // use 'DepartmentID' for clarity
	City         string
	Department   Department `gorm:"foreignKey:DepartmentID"` // Establishes foreign key relationship
}

type Department struct {
	Code   int     `gorm:"primaryKey"`
	Name   string  `gorm:"not null"`
	Budget float64 `gorm:"not null"`
}
