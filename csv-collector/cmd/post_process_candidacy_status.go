// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package cmd

import (
	"github.com/olhoneles/politicos-go/csv-collector/collector"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var postProcessCandidacyStatusCmd = &cobra.Command{
	Use:   "post-process-candidacies-status",
	Short: "Post process candidacies status data",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Processing candidacies status...")
		if err := collector.ProcessAllCandidaciesStatus(); err != nil {
			log.Fatalf("Couldn't process candidacies status! %v", err)
		}
	},
}
