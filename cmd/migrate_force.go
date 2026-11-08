package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var migrateForceCommand = &cobra.Command{
	Use:   "force [version]",
	Short: "Force set migration version",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		version, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid migration version: %w", err)
		}

		if version < -1 {
			return fmt.Errorf("migration version must be greater than or equal to -1")
		}

		runtime, err := newMigrationRuntime(cmd)
		if err != nil {
			return err
		}
		defer runtime.Close()

		runtime.Logger.Warn(
			"Forcing migration version",
			zap.Int("version", version),
		)

		if err := runtime.Migration.Force(version); err != nil {
			return fmt.Errorf("failed to force migration version: %w", err)
		}

		runtime.Logger.Info(
			"Migration version updated successfully",
			zap.Int("version", version),
		)

		return nil
	},
}
