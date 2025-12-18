package migrations

import (
	"errors"
	"fmt"

	"sort"

	"github.com/Nattakornn/cache/config"
	"github.com/Nattakornn/cache/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Migration struct {
	Number uint `gorm:"primary_key"`
	Name   string

	Forwards func(db *gorm.DB) error `gorm:"-"`
}

var Migrations []*Migration

func Migrate(dryRun bool, number int, forceMigrate bool, cfg config.IDbConfig) error {

	// check for duplicate migration Number
	migrationIDs := make(map[uint]struct{})
	for _, migration := range Migrations {
		if _, ok := migrationIDs[migration.Number]; ok {
			err := fmt.Errorf("duplicate migration number found: %d", migration.Number)
			logger.Logger.Errorf("unable to apply migrations, err: %+v", err)
			return err
		}

		migrationIDs[migration.Number] = struct{}{}
	}

	sort.Slice(Migrations, func(i, j int) bool {
		return Migrations[i].Number < Migrations[j].Number
	})

	db, err := gorm.Open(postgres.Open(cfg.Url()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logger.Logger.Errorf("unable to connect db: %+v", err)
		return fmt.Errorf("unable to connect db: %v", err)
	}

	// Force Migrate Zone
	if forceMigrate {
		logger.Logger.Infof("=== FORCE MIGRATE ===")
		if err := db.Migrator().DropTable(&Migration{}); err != nil {
			return fmt.Errorf("unable to drop migrations table: %v", err)
		}
	}

	// Make sure Migration table is there
	logger.Logger.Debugf("ensuring migrations table is present")
	if err := db.AutoMigrate(&Migration{}); err != nil {
		return fmt.Errorf("unable to automatically migrate migrations table: %v", err)
	}

	var latest Migration
	if err := db.Order("number desc").First(&latest).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("unable to find latest migration: %v", err)
	}

	noMigrationsApplied := latest.Number == 0

	if noMigrationsApplied && len(Migrations) == 0 {
		logger.Logger.Infof("no migrations to apply")
		return nil
	}

	if latest.Number >= Migrations[len(Migrations)-1].Number {
		logger.Logger.Infof("no migrations to apply")
		return nil
	}

	if number == -1 {
		number = int(Migrations[len(Migrations)-1].Number)
	}

	if uint(number) <= latest.Number && latest.Number > 0 {
		logger.Logger.Infof("no migrations to apply, specified number is less than or equal to latest migration; backwards migrations are not supported")
		return nil
	}

	for _, migration := range Migrations {
		if migration.Number > uint(number) {
			break
		}

		if migration.Number <= latest.Number {
			continue
		}

		if latest.Number > 0 {
			logger.Logger.Infof("continuing migration starting from %d", migration.Number)
		}

		migrationLogger := logger.Logger.With(
			"migration_number", migration.Number,
		)

		migrationLogger.Infof("applying migration %q", migration.Name)

		if dryRun {
			continue
		}

		tx := db.Begin()

		if err := migration.Forwards(tx); err != nil {
			logger.Logger.Errorf("unable to apply migration, rolling back. err: %+v", err)
			if err := tx.Rollback().Error; err != nil {
				logger.Logger.Errorf("unable to rollback... err: %+v", err)
			}
			break
		}

		if err := tx.Commit().Error; err != nil {
			logger.Logger.Errorf("unable to commit transaction... err: %+v", err)
			break
		}

		// Create migration record
		if err := db.Create(migration).Error; err != nil {
			logger.Logger.Errorf("unable to create migration record. err: %+v", err)
			break
		}
	}

	return nil
}
