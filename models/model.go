package models

import "time"

// Model contains default database elements
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id" validate:"omitempty"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
