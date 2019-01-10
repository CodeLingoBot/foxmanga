package migration

import (
	"github.com/go-xorm/xorm"
	migrate2 "github.com/go-xorm/xorm/migrate"
)

var (
	migration MigrationPipeline
)

type MigrationPipeline interface {
	Migrate() error
}

type globalMigration struct {
	Db *xorm.Engine
}

func GetMigration(db *xorm.Engine) MigrationPipeline {
	if migration == nil {
		migration = &globalMigration{Db: db}
	}

	return migration
}
func (m *globalMigration) Migrate() error {
	var totalMigrations []*migrate2.Migration
	totalMigrations = append(totalMigrations, GetCategoryMigration()...)
	//_m := migrate2.New(m.Db, migrate2.DefaultOptions, totalMigrations)
	_m := migrate2.New(m.Db, migrate2.DefaultOptions, totalMigrations)
	return _m.Migrate()
}
