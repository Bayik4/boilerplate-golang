package cmd

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var migrateUpCommand = &cobra.Command{
	Use:   "up",
	Short: "Run all pending migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		runtime, err := newMigrationRuntime(cmd)
		if err != nil {
			return err
		}
		defer runtime.Close()

		runtime.Logger.Info("Running database migrations")

		if err := runtime.Migration.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				runtime.Logger.Info("Database is already up to date")
				return nil
			}

			return fmt.Errorf("failed to run migrations: %w", err)
		}

		version, dirty, err := runtime.Migration.Version()
		if err != nil {
			return fmt.Errorf("failed to get migration version: %w", err)
		}

		runtime.Logger.Info(
			"Database migrations completed",
			zap.Uint("version", version),
			zap.Bool("dirty", dirty),
		)

		return nil
	},
}
