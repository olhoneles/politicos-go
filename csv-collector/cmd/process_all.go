// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package cmd

import (
	"github.com/spf13/cobra"
)

var postProcessAllCmd = &cobra.Command{
	Use:   "process-all",
	Short: "Process all data",
	Run: func(cmd *cobra.Command, args []string) {
		processCSVFilesCmd.Execute()
		postProcessPoliticalPartyCmd.Execute()
		postProcessPoliticalOfficeCmd.Execute()
		postProcessCandidacyStatusCmd.Execute()
		postProcessEducationCmd.Execute()
	},
}
