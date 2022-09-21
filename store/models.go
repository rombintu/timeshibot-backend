package store

import "gorm.io/gorm"

type REST struct {
	Message string `json:"message"`
	Error   int    `json:"error"`
}

type Group struct {
	gorm.Model
	ChatID    string    `json:"chat_id" gorm:"unique"`
	TimeTable TimeTable `gorm:"foreignKey:ID"`
}

type TimeTable struct {
	gorm.Model
	Weeks []Week `gorm:"foreignKey:ID"`
}

type Week struct {
	gorm.Model
	Days []Day `gorm:"foreignKey:ID"`
}

type Day struct {
	gorm.Model
	WeekName string    `json:"chat_id" gorm:"unique"`
	Subjects []Subject `gorm:"foreignKey:ID"`
}

type Subject struct {
	gorm.Model
	Name    string
	Office  string
	Teacher string
}
