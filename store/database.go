package store

import (
	"fmt"

	postgres "gorm.io/driver/postgres"
	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/rombintu/timeshibot-backend/tools"
)

type Store struct {
	Driver *gorm.DB
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Open() error {
	if tools.GetEnvOrDefaultBool("PROD") {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			tools.GetEnvOrDefault("DB_HOST", tools.DefaultHost),
			tools.GetEnvOrDefault("DB_USER", tools.DefaultUser),
			tools.GetEnvOrDefault("DB_PASS", tools.DefaultPass),
			tools.GetEnvOrDefault("DB_NAME", tools.DefaultNameDatabase),
			tools.GetEnvOrDefault("DB_PORT", tools.DefaultPortDatabase),
			tools.GetEnvOrDefault("DB_MODE", tools.DefaultMode),
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
		s.Driver = db
	} else {
		db, err := gorm.Open(sqlite.Open(tools.DefaultDevDatabase), &gorm.Config{})
		if err != nil {
			return err
		}
		s.Driver = db
	}
	return nil
}
