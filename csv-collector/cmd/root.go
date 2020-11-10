// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package cmd

import (
	"github.com/spf13/cobra"

	_ "github.com/olhoneles/politicos-go/config"
)

var (
	rootCmd = &cobra.Command{
		Use:   "collector",
		Short: "Politicos API management tool",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(fetchCSVFilesCmd)
	rootCmd.AddCommand(postProcessAllCmd)
	rootCmd.AddCommand(processCSVFilesCmd)
	rootCmd.AddCommand(postProcessPoliticalPartyCmd)
	rootCmd.AddCommand(postProcessPoliticalOfficeCmd)
	rootCmd.AddCommand(postProcessCandidacyStatusCmd)
	rootCmd.AddCommand(postProcessEducationCmd)
}
