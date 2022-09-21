package store

import (
	"fmt"

	postgres "gorm.io/driver/postgres"
	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/rombintu/timeshibot-backend/tools"
)

type Database struct {
	Driver *gorm.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Open() error {
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
		d.Driver = db
	} else {
		db, err := gorm.Open(sqlite.Open(tools.DefaultDevDatabase), &gorm.Config{})
		if err != nil {
			return err
		}
		d.Driver = db
	}
	return nil
}

// func (d *Database) Close() error {
// 	return d.Driver.Close()
// }
