// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package main

import (
	"os"

	"github.com/olhoneles/politicos-go/csv-collector/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
