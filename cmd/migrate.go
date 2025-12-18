package cmd

import (
	"github.com/Nattakornn/cache/config"
	"github.com/Nattakornn/cache/pkg/databases/postgressql/migrations"
	"github.com/Nattakornn/cache/pkg/logger"
	"github.com/spf13/cobra"
)

var forceMigrate bool = false

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Cache Database",
	Run: func(cmd *cobra.Command, args []string) {

		cfg := config.LoadConfig()

		logger.InitZapLogger(cfg.Utils().Log())
		defer logger.SyncLogger()

		migrations.Migrate(false, -1, forceMigrate, cfg.Db())

	},
}

func init() {
	rootCmd.AddCommand(MigrateCmd)
	MigrateCmd.PersistentFlags().BoolVar(&forceMigrate, "force", false, "force migrate (default is false)")
}
