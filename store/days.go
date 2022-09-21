package store

func (d *Database) UpdateDay(
	ChatID, week, day string,
	subjects []Subject,
) error {
	// if err := d.Driver.Model(
	// 	&Day{}).
	// 	Where("groups.chat_id == ?", ChatID).
	// 	Updates(Day{Subjects: subjects}).
	// 	Error; err != nil {
	// 	return err
	// }
	return nil
}
