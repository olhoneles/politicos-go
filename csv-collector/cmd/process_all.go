// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package cmd

import (
	"github.com/spf13/cobra"
)

var postProcessAllCmd = &cobra.Command{
	Use:   "process-all",
	Short: "Process all data",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := processCSVFilesCmd.Execute(); err != nil {
			return err
		}
		if err := postProcessPoliticalPartyCmd.Execute(); err != nil {
			return err
		}
		if err := postProcessPoliticalOfficeCmd.Execute(); err != nil {
			return err
		}

		if err := postProcessCandidacyStatusCmd.Execute(); err != nil {
			return err
		}
		if err := postProcessEducationCmd.Execute(); err != nil {
			return err
		}
		return nil
	},
}
