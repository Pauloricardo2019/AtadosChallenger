package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
)

var MigrationsToExec []*gormigrate.Migration

func GetMigrationsToExec() []*gormigrate.Migration {
	return MigrationsToExec
}
