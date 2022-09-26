package store

func (s *Store) UpdateDay(
	chatID string,
	weekIndex, dayIndex int,
	subjects []Subject,
) error {
	if err := s.Open(); err != nil {
		return err
	}

	group, err := s.GetGroup(chatID)
	if err != nil {
		return err
	}
	var week []Week
	var days []Day

	s.Driver.Model(&group).Association("Week").Find(&week)
	s.Driver.Model(&week).Association("Days").Find(&days)

	days[dayIndex].Subjects = subjects
	if err := s.Driver.Save(&days); err != nil {
		return err.Error
	}
	return nil
}
