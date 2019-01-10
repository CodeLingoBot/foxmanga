package migration

import (
	"Project/foxmanga/pkg/category"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/xorm/migrate"
)

func GetCategoryMigration() []*migrate.Migration {
	migrations := []*migrate.Migration{
		{
			ID: "create-category-table",
			Migrate: func(engine *xorm.Engine) error {
				return engine.Sync2(&category.Category{})
			},
			Rollback: func(engine *xorm.Engine) error {
				return engine.DropTables(&category.Category{})
			},
		},
	}
	return migrations
}
