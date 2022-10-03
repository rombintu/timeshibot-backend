package store

import (
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
	TTID    uint   `json:"tt_id"`
	Time    string `json:"time"`
	Title   string `json:"title"`
	Office  string `json:"office"`
	Teacher string `json:"teacher"`
	Comment string `json:"comment"`
}

// type JSONTime time.Time

// func (j *JSONTime) UnmarshalJSON(b []byte) error {
// 	s := strings.Trim(string(b), "\"")
// 	t, err := time.Parse("1900-01-01T09:00:00", s)
// 	if err != nil {
// 		return err
// 	}
// 	*j = JSONTime(t)
// 	return nil
// }

// func (j JSONTime) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(time.Time(j))
// }

// func (j JSONTime) Format(s string) string {
// 	t := time.Time(j)
// 	return t.Format(s)
// }
