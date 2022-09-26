package store

import "gorm.io/gorm"

type REST struct {
	Message string `json:"message"`
	Error   int    `json:"error"`
}

type Group struct {
	gorm.Model
	ChatID string `json:"chat_id" gorm:"unique"`
	Weeks  []Week `gorm:"foreignKey:ID"`
}

type Week struct {
	gorm.Model
	Days []Day `gorm:"foreignKey:ID"`
}

type Day struct {
	gorm.Model
	Subjects []Subject `gorm:"foreignKey:ID"`
}

type Subject struct {
	gorm.Model
	Title   string `json:"title"`
	Office  string `json:"office"`
	Teacher string `json:"teacher"`
}
