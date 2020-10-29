// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/olhoneles/politicos-go/api"
	_ "github.com/olhoneles/politicos-go/config"
)

// @title Politicos API
// @version 0.1.0
// @description Politicos API data.

// @contact.name API Support
// @contact.url https://github.com/olhoneles/politicos-go

// @license.name AGPL-3.0+
// @license.url https://www.gnu.org/licenses/agpl-3.0.pt-br.html

// @host localhost:8888
// @BasePath /api/v1
func main() {
	server, err := api.NewServerFromDB()
	if err != nil {
		log.Error(err)
	}
	server.Start()
}
