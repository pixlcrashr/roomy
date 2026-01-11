package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/pixlcrashr/roomy/pkg/db/migrations"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  `Run database migrations to create or update the database schema.`,
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply all pending migrations",
	Long:  `Apply all pending database migrations to bring the schema up to date.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := mustConnectDB()
		defer db.Close()

		fmt.Println("Running database migrations...")
		if err := migrations.Run(db); err != nil {
			fmt.Fprintf(os.Stderr, "Migration failed: %v\n", err)
			os.Exit(1)
		}

		printVersion(db)
		fmt.Println("Migrations completed successfully")
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Rollback the last migration",
	Long:  `Rollback the most recently applied database migration.`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		db := mustConnectDB()
		defer db.Close()

		fmt.Println("Rolling back migrations...")
		var err error
		if all {
			err = migrations.RollbackAll(db)
		} else {
			err = migrations.Rollback(db)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Rollback failed: %v\n", err)
			os.Exit(1)
		}

		printVersion(db)
		fmt.Println("Rollback completed successfully")
	},
}

var migrateVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current migration version",
	Long:  `Display the current database migration version and dirty state.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := mustConnectDB()
		defer db.Close()

		printVersion(db)
	},
}

var migrateForceCmd = &cobra.Command{
	Use:   "force [version]",
	Short: "Force set migration version",
	Long: `Force set the migration version without running migrations.
This is useful for fixing a dirty database state after a failed migration.
Use version -1 to clear the version.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var version int
		if _, err := fmt.Sscanf(args[0], "%d", &version); err != nil {
			fmt.Fprintf(os.Stderr, "Invalid version: %v\n", err)
			os.Exit(1)
		}

		db := mustConnectDB()
		defer db.Close()

		if err := migrations.Force(db, version); err != nil {
			fmt.Fprintf(os.Stderr, "Force version failed: %v\n", err)
			os.Exit(1)
		}

		printVersion(db)
		fmt.Println("Version forced successfully")
	},
}

var migrateStepsCmd = &cobra.Command{
	Use:   "steps [n]",
	Short: "Apply n migrations",
	Long: `Apply n migrations. Positive n applies up migrations, negative n applies down migrations.
Example: migrate steps 2   (apply 2 up migrations)
Example: migrate steps -1  (rollback 1 migration)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var n int
		if _, err := fmt.Sscanf(args[0], "%d", &n); err != nil {
			fmt.Fprintf(os.Stderr, "Invalid number of steps: %v\n", err)
			os.Exit(1)
		}

		db := mustConnectDB()
		defer db.Close()

		fmt.Printf("Applying %d migration steps...\n", n)
		if err := migrations.Steps(db, n); err != nil {
			fmt.Fprintf(os.Stderr, "Steps migration failed: %v\n", err)
			os.Exit(1)
		}

		printVersion(db)
		fmt.Println("Steps migration completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	migrateCmd.AddCommand(migrateVersionCmd)
	migrateCmd.AddCommand(migrateForceCmd)
	migrateCmd.AddCommand(migrateStepsCmd)

	migrateDownCmd.Flags().Bool("all", false, "Rollback all migrations")
}

func mustConnectDB() *sql.DB {
	db, err := sql.Open("postgres", config.Database.URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to ping database: %v\n", err)
		os.Exit(1)
	}

	return db
}

func printVersion(db *sql.DB) {
	version, dirty, err := migrations.Version(db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get version: %v\n", err)
		return
	}

	dirtyStr := ""
	if dirty {
		dirtyStr = " (dirty)"
	}
	fmt.Printf("Current version: %d%s\n", version, dirtyStr)
}
