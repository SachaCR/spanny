package cmd

import (
	"context"
	"fmt"
	"os"

	conf "spanny/src/config"
	"spanny/src/ui"

	"cloud.google.com/go/spanner"
	"github.com/spf13/cobra"
)

var config *conf.Config

func init() {
	rootCmd.PersistentFlags().StringVarP(&env, "env", "e", "default", "Specify spanner environment")
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "./", "Indicate the configuration path")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Makes spanny more verbose")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(stateCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(createInstanceCmd)
	rootCmd.AddCommand(createDatabaseCmd)
	rootCmd.AddCommand(initMigrationCmd)
	rootCmd.AddCommand(listTablesCmd)
	rootCmd.AddCommand(listDatabasesCmd)
	rootCmd.AddCommand(queryCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(rollbackCmd)
	rootCmd.AddCommand(applyAllCmd)
	rootCmd.AddCommand(resetCmd)
}

var rootCmd = &cobra.Command{
	Use:   "spanny",
	Short: "Spanny database schema migration CLI tool for Spanner",
	Long:  `Spanny is a very CLI tool helping you to manage database schema migration with the Spanner emulator`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		loadedConfig, err := conf.LoadConfiguration(env, configPath)

		if err != nil {
			panic(fmt.Errorf("LOADING CONFIGURATION ERROR: %s", err.Error()))
		}

		config = &loadedConfig

		if verbose {
			println(fmt.Sprintf("\nSpanny environment: [%s]\n", env))
			println("Configuration loaded successfully")
			println(getDatabasePath())
			println()
		}

	},

	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		databasePath := getDatabasePath()

		client, err := spanner.NewClient(ctx, databasePath)

		if err != nil {
			println(err.Error())
			return
		}

		defer client.Close()

		stmt := spanner.Statement{SQL: `SELECT * FROM spanner_migrations ORDER BY applied_at ASC`}
		iter := client.Single().Query(ctx, stmt)
		defer iter.Stop()

		ui.RenderTableFromRowIterator(iter)

		stmtLock := spanner.Statement{SQL: `SELECT is_locked FROM spanner_migrations_lock`}
		iterLock := client.Single().Query(ctx, stmtLock)
		defer iterLock.Stop()

		ui.RenderTableFromRowIterator(iterLock)
	},
}

func Start() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
