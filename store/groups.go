package store

func (d *Database) CreateGroup(ChatID string) error {
	return d.Driver.Create(&Group{ChatID: ChatID}).Error
}
