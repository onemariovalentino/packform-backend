package command

import (
	"context"
	"log"
	"os"
	"packform-backend/src/pkg/di"

	"github.com/spf13/cobra"
)

type (
	Command struct {
		comm *cobra.Command
	}
)

func New() *Command {
	services := di.NewDependency()

	importCmd := &cobra.Command{
		Use:   "import",
		Short: "Import data into a table from CSV file(s)",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			destination, _ := cmd.Flags().GetString("destination")
			files, _ := cmd.Flags().GetStringArray("files")

			err := services.OrderUsecase.FeedingDataFromCSV(ctx, destination, files)
			if err != nil {
				return err
			}

			return nil
		},
	}
	importCmd.Flags().StringP("destination", "d", "", "Destination for save to table")
	importCmd.Flags().StringArrayP("files", "f", []string{}, "Files to import")
	importCmd.MarkFlagRequired("destination")
	importCmd.MarkFlagRequired("files")

	rootCmd := &cobra.Command{Use: "cli"}
	rootCmd.AddCommand(importCmd)

	return &Command{comm: rootCmd}
}

func (c *Command) Run() {
	if err := c.comm.Execute(); err != nil {
		log.Fatalf("err:%s", err)
		os.Exit(2)
	}
}
