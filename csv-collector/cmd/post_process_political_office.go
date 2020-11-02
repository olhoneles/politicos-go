// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package cmd

import (
	"github.com/olhoneles/politicos-go/csv-collector/collector"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var postProcessPoliticalOfficeCmd = &cobra.Command{
	Use:   "post-process-political-office",
	Short: "Post process political office data",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Processing political offices...")
		if err := collector.ProcessAllPoliticalOffices(); err != nil {
			log.Fatalf("Couldn't process political offices! %v", err)
		}
	},
}
