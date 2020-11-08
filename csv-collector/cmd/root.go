// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

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

func postInitCommands(commands []*cobra.Command) {
	for _, cmd := range commands {
		presetRequiredFlags(cmd)
		if cmd.HasSubCommands() {
			postInitCommands(cmd.Commands())
		}
	}
}

func presetRequiredFlags(cmd *cobra.Command) {
	viper.BindPFlags(cmd.Flags()) // #nosec
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if viper.IsSet(f.Name) && viper.GetString(f.Name) != "" {
			cmd.Flags().Set(f.Name, viper.GetString(f.Name))
		}
	})
}

func init() {
	cobra.OnInitialize(func() {
		postInitCommands(rootCmd.Commands())
	})

	rootCmd.AddCommand(fetchCSVFilesCmd)
	rootCmd.AddCommand(postProcessAllCmd)
	rootCmd.AddCommand(processCSVFilesCmd)
	rootCmd.AddCommand(postProcessPoliticalPartyCmd)
	rootCmd.AddCommand(postProcessPoliticalOfficeCmd)
	rootCmd.AddCommand(postProcessCandidacyStatusCmd)
	rootCmd.AddCommand(postProcessEducationCmd)
}
