package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202306290926_create_voluntary_table",
		Migrate: func(tx *gorm.DB) error {

			sql := `CREATE TABLE IF NOT EXISTS voluntaries (
					id SERIAL PRIMARY KEY,
					first_name VARCHAR(255) NOT NULL,
    				last_name VARCHAR(255) NOT NULL,
    				neighborhood VARCHAR(255) NOT NULL,
    				city VARCHAR(255) NOT NULL,
					created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
				)`
			if err := tx.Exec(sql).Error; err != nil {
				return err
			}

			return nil

		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("voluntaries")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
