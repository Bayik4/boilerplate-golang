package cmd

import (
	"fmt"
	"net"
	"net/url"

	"github.com/bayik4/boilerplate-golang/boot"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const migrationPath = "file://migrations"

type migrationRuntime struct {
	Migration *migrate.Migrate
	Logger    *zap.Logger
}

func newMigrationRuntime(cmd *cobra.Command) (*migrationRuntime, error) {
	ctx := cmd.Context()
	vpr := viper.New()

	logConfig := boot.LoadLoggerConfig(vpr)

	logger, err := boot.LoggerInit(logConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	cfg := boot.LoadConfig(ctx, vpr, logger)
	cfg.Logger = logConfig

	db := cfg.Database.Pgsql.DBInit

	dsn := url.URL{
		Scheme: "postgres",
		User: url.UserPassword(
			db.Username,
			db.Password,
		),
		Host: net.JoinHostPort(
			db.Host,
			db.Port,
		),
		Path: "/" + db.Dbname,
		RawQuery: url.Values{
			"sslmode": {"disable"},
		}.Encode(),
	}

	m, err := migrate.New(
		migrationPath,
		dsn.String(),
	)
	if err != nil {
		_ = logger.Sync()
		return nil, fmt.Errorf("failed to initialize migration: %w", err)
	}

	return &migrationRuntime{
		Migration: m,
		Logger:    logger,
	}, nil
}

func (r *migrationRuntime) Close() {
	sourceErr, databaseErr := r.Migration.Close()

	if sourceErr != nil {
		r.Logger.Error(
			"failed to close migration source",
			zap.Error(sourceErr),
		)
	}

	if databaseErr != nil {
		r.Logger.Error(
			"failed to close migration database",
			zap.Error(databaseErr),
		)
	}

	_ = r.Logger.Sync()
}

var migrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Manage database migrations",
	Long:  "Manage database schema migrations.",
}

func init() {
	migrateCommand.AddCommand(
		migrateUpCommand,
		migrateDownCommand,
		migrateVersionCommand,
		migrateForceCommand,
		migrateCreateCommand,
	)
}
