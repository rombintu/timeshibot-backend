package store

import (
	"time"

	"gorm.io/gorm"
)

type REST struct {
	Message interface{} `json:"message"`
	Error   int         `json:"error"`
}

type Timetable struct {
	gorm.Model
	ChatID   string
	Name     string
	Week     string
	Subjects []Subject `gorm:"foreignKey:TTID;"`
	Comment  string    `json:"comment"`
}

type Subject struct {
	gorm.Model
	TTID    uint      `json:"tt_id"`
	Time    time.Time `json:"time"`
	Title   string    `json:"title"`
	Office  string    `json:"office"`
	Teacher string    `json:"teacher"`
	Comment string    `json:"comment"`
}
