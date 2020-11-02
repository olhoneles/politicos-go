// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package cmd

import (
	"github.com/olhoneles/politicos-go/csv-collector/collector"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var postProcessEducationCmd = &cobra.Command{
	Use:   "post-process-education",
	Short: "Post process education data",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Processing education...")
		if err := collector.ProcessAllEducations(); err != nil {
			log.Fatalf("Couldn't process education! %v", err)
		}
	},
}
