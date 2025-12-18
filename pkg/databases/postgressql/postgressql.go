package postgresql

import (
	"os"

	"github.com/Nattakornn/cache/config"
	"github.com/Nattakornn/cache/pkg/logger"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func ConnectDb(cfg config.IDbConfig) *sqlx.DB {
	db, err := sqlx.Connect("pgx", cfg.Url())
	if err != nil {
		logger.Logger.Errorf("connect to postgresql db failed: %v\n", err)
		os.Exit(1)
	}
	db.DB.SetMaxOpenConns(cfg.MaxOpenConns())
	return db
}
