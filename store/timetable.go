package store

func (s *Store) CreateOrUpdateTimeTable(ttFind Timetable, subjects []Subject) error {
	// TODO
	var tt Timetable
	if err := s.Driver.
		FirstOrCreate(&tt, ttFind).Error; err != nil {
		return err
	}
	if err := s.Driver.Unscoped().Delete(&Subject{}, Subject{TTID: tt.ID}).Error; err != nil {
		return err
	}
	tt.Subjects = subjects
	s.Driver.Save(&tt)

	return nil
}

func (s *Store) GetTimeTable(ttFind Timetable) ([]Subject, error) {
	if err := s.Open(); err != nil {
		return []Subject{}, err
	}

	var tt Timetable
	var subjects []Subject

	if err := s.Driver.First(&tt).
		Where(
			"chat_id == ? AND name == ? AND week == ?",
			ttFind.ChatID, ttFind.Name, ttFind.Week).Error; err != nil {
		return []Subject{}, err
	}

	if err := s.Driver.Find(&subjects).
		Order("time DESC"). // TODO
		Where("tt_id == ?", tt.ID).
		Error; err != nil {
		return []Subject{}, err
	}

	return subjects, nil
}

func (s *Store) GetTimeTableAll() ([]Timetable, error) {
	if err := s.Open(); err != nil {
		return []Timetable{}, err
	}

	var tts []Timetable
	err := s.Driver.Find(&tts).Error
	if err != nil {
		return []Timetable{}, err
	}
	return tts, nil
}
