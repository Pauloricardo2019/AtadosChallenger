package main

import (
	"atados/challenger/cmd/migration/migrations"
	"atados/challenger/internal/config"
	"github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/go-gormigrate/gormigrate/v2"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var envPath = "./dev.env"

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	cfg := config.NewConfig(logger).GetConfig(envPath)

	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migrationsToExec := migrations.GetMigrationsToExec()
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrationsToExec)
	if err := m.Migrate(); err != nil {
		log.Fatalln("Could not migrate: ", err)
	}

	log.Println("Migration did run successfully")

}
