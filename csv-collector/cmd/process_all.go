// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package cmd

import (
	"github.com/spf13/cobra"
)

var postProcessAllCmd = &cobra.Command{
	Use:   "process-all",
	Short: "Imports the data from CSV and post-processes it",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := importCSVFilesCmd.Execute(); err != nil {
			return err
		}
		if err := postProcessAll(); err != nil {
			return err
		}
		return nil
	},
}
