package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202306290928_create_action_table",
		Migrate: func(tx *gorm.DB) error {

			sql := `CREATE TABLE IF NOT EXISTS actions (
					id SERIAL PRIMARY KEY,
					name VARCHAR(255) NOT NULL,
    				institution VARCHAR(255) NOT NULL,
    				city VARCHAR(255) NOT NULL,
					neighborhood VARCHAR(255) NOT NULL,
    				address VARCHAR(255) NOT NULL,
					created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
				)`
			if err := tx.Exec(sql).Error; err != nil {
				return err
			}

			return nil

		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("actions")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
