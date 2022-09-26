package store

func (s *Store) CreateGroup(chatID string) error {
	group := &Group{
		ChatID: chatID,
	}

	if err := s.Driver.Create(group).Error; err != nil {
		return err
	}
	return nil
}

func (s *Store) GetGroup(chatID string) (Group, error) {
	if err := s.Open(); err != nil {
		return Group{}, err
	}
	var group Group
	err := s.Driver.First(&group).Where("chat_id == ?", chatID).Error
	if err != nil {
		return Group{}, err
	}

	return group, nil
}

func (s *Store) GetGroupAll(chatID string) ([]Group, error) {
	if err := s.Open(); err != nil {
		return []Group{}, err
	}

	var groups []Group
	err := s.Driver.Find(&groups).Error
	if err != nil {
		return []Group{}, err
	}
	return groups, nil
}
