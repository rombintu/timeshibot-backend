package store

func (s *Store) CreateGroup(chatID string) error {
	tt := &Timetable{
		ChatID: chatID,
	}

	if err := s.Driver.Create(tt).Error; err != nil {
		return err
	}
	return nil
}

func (s *Store) GetGroup(chatID string) (Timetable, error) {
	if err := s.Open(); err != nil {
		return Timetable{}, err
	}
	var tt Timetable
	err := s.Driver.First(&tt).Where("chat_id == ?", chatID).Error
	if err != nil {
		return Timetable{}, err
	}

	return tt, nil
}

func (s *Store) GetGroupAll(chatID string) ([]Timetable, error) {
	if err := s.Open(); err != nil {
		return []Timetable{}, err
	}

	var groups []Timetable
	err := s.Driver.Find(&groups).Error
	if err != nil {
		return []Timetable{}, err
	}
	return groups, nil
}
