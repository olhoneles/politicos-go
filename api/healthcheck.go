// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func (s *server) healthcheck(c echo.Context) error {
	log.Debug("[API] Ping")
	if err := s.db.Ping(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "WORKING")
}
