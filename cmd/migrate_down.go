package cmd

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var migrateDownCommand = &cobra.Command{
	Use:   "down",
	Short: "Rollback the last migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		runtime, err := newMigrationRuntime(cmd)
		if err != nil {
			return err
		}
		defer runtime.Close()

		runtime.Logger.Info("Rolling back last migration")

		if err := runtime.Migration.Steps(-1); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				runtime.Logger.Info("No migration available to rollback")
				return nil
			}

			return fmt.Errorf("failed to rollback migration: %w", err)
		}

		version, dirty, err := runtime.Migration.Version()
		if err != nil {
			if errors.Is(err, migrate.ErrNilVersion) {
				runtime.Logger.Info("All migrations have been rolled back")
				return nil
			}

			return fmt.Errorf("failed to get migration version: %w", err)
		}

		runtime.Logger.Info(
			"Migration rollback completed",
			zap.Uint("version", version),
			zap.Bool("dirty", dirty),
		)

		return nil
	},
}
