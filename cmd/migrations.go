package cmd

import (
	"log"
	"marketplace-api/database"
	"marketplace-api/pkg/migrations"

	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	dbConn, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	migrations.Run(dbConn, args)
}

func Migrations() cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "migrations",
		Short: "Run migrations command",
		Run:   run,
	}

	return *rootCmd
}
