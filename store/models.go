package store

import "gorm.io/gorm"

type REST struct {
	Message string `json:"message"`
	Error   int    `json:"error"`
}

type Timetable struct {
	gorm.Model
	ChatID   string
	Name     string
	Week     string
	Subjects []Subject `gorm:"foreignKey:ID"`
}

type Subject struct {
	gorm.Model
	Title   string `json:"title"`
	Office  string `json:"office"`
	Teacher string `json:"teacher"`
}
