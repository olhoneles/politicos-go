// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package cmd

import (
	"github.com/olhoneles/politicos-go/csv-collector/collector"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var postProcessPoliticalPartyCmd = &cobra.Command{
	Use:   "post-process-political-party",
	Short: "Post process political party data",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Processing political parties...")
		if err := collector.ProcessAllPoliticalParties(); err != nil {
			log.Fatalf("Couldn't process political parties! %v", err)
		}
	},
}
