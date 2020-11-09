// Copyright (c) 2020, Lucca Mendonça
// Licensed under the AGPL-3.0+ License

package cmd

import (
	"github.com/olhoneles/politicos-go/csv-collector/collector"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	postProcessCmd = &cobra.Command{
		Use: "post-process",
	}
	postProcessSteps = map[string]func() error{
		"political-party":  collector.ProcessAllPoliticalParties,
		"political-office": collector.ProcessAllPoliticalOffices,
		"candidacy-status": collector.ProcessAllCandidaciesStatus,
		"education":        collector.ProcessAllEducations,
	}
)

func init() {
	for runningStep, postProcessFunc := range postProcessSteps {
		postProcessCmd.AddCommand(&cobra.Command{
			Use: runningStep,
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Debugf("Post-processing '%s'...", cmd.Name())
				return postProcessFunc()
			},
		})
	}
	postProcessCmd.AddCommand(&cobra.Command{
		Use: "all",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Debug("Post-processing 'all'...")
			return postProcessAll()
		},
	})
}

func postProcessAll() error {
	var err error
	for runningStep, postProcessFunc := range postProcessSteps {
		log.Debugf("Running '%s'...	", runningStep)
		err = postProcessFunc()
		if err != nil {
			return err
		}
	}

	return nil
}
