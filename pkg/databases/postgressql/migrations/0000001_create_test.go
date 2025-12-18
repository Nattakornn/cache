package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

var createTestTableMigration = &Migration{
	Number: 1,
	Name:   "create test table",
	Forwards: func(db *gorm.DB) error {

		const dropSql = `
			DROP TABLE IF EXISTS test CASCADE;
		`
		err := db.Exec(dropSql).Error
		if err != nil {
			return fmt.Errorf("unable to drop test table: %v", err)
		}

		const sql = `
			CREATE TABLE "test" (
				"id" bigserial,
				"name" text NOT NULL,
				"created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY ("id")
			);
		`

		err = db.Exec(sql).Error
		if err != nil {
			return fmt.Errorf("unable to create test table: %v", err)
		}

		return nil
	},
}

func init() {
	Migrations = append(Migrations, createTestTableMigration)
}
