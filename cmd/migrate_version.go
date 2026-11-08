package cmd

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var migrateVersionCommand = &cobra.Command{
	Use:   "version",
	Short: "Show current migration version",
	RunE: func(cmd *cobra.Command, args []string) error {
		runtime, err := newMigrationRuntime(cmd)
		if err != nil {
			return err
		}
		defer runtime.Close()

		version, dirty, err := runtime.Migration.Version()
		if err != nil {
			if errors.Is(err, migrate.ErrNilVersion) {
				runtime.Logger.Info("No migration has been applied yet")
				return nil
			}

			return fmt.Errorf("failed to get migration version: %w", err)
		}

		runtime.Logger.Info(
			"Current migration version",
			zap.Uint("version", version),
			zap.Bool("dirty", dirty),
		)

		return nil
	},
}
