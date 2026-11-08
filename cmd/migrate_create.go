package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var migrateCreateCommand = &cobra.Command{
	Use:   "create [migration_name]",
	Short: "Create a new migration",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := strings.ToLower(args[0])

		validName := regexp.MustCompile(`^[a-z][a-z0-9_]*$`)
		if !validName.MatchString(name) {
			return fmt.Errorf(
				"invalid migration name: use lowercase letters, numbers, and underscores",
			)
		}

		migrationDir := "./migrations"

		if err := os.MkdirAll(migrationDir, 0755); err != nil {
			return fmt.Errorf("failed to create migration directory: %w", err)
		}

		entries, err := os.ReadDir(migrationDir)
		if err != nil {
			return fmt.Errorf("failed to read migration directory: %w", err)
		}

		var maxVersion uint64

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			filename := entry.Name()
			parts := strings.SplitN(filename, "_", 2)

			if len(parts) < 2 {
				continue
			}

			version, err := strconv.ParseUint(parts[0], 10, 64)
			if err != nil {
				continue
			}

			if version > maxVersion {
				maxVersion = version
			}
		}

		nextVersion := maxVersion + 1
		prefix := fmt.Sprintf("%06d_%s", nextVersion, name)

		upPath := filepath.Join(
			migrationDir,
			prefix+".up.sql",
		)

		downPath := filepath.Join(
			migrationDir,
			prefix+".down.sql",
		)

		files := []string{upPath, downPath}

		created := make([]string, 0, len(files))

		for _, file := range files {
			f, err := os.OpenFile(
				file,
				os.O_CREATE|os.O_EXCL|os.O_WRONLY,
				0644,
			)
			if err != nil {
				for _, path := range created {
					_ = os.Remove(path)
				}

				return fmt.Errorf("failed to create migration file: %w", err)
			}

			if err := f.Close(); err != nil {
				for _, path := range created {
					_ = os.Remove(path)
				}

				_ = os.Remove(file)

				return fmt.Errorf("failed to close migration file: %w", err)
			}

			created = append(created, file)
		}

		cmd.Printf("Created migration:\n")
		cmd.Printf("  UP:   %s\n", upPath)
		cmd.Printf("  DOWN: %s\n", downPath)

		return nil
	},
}
