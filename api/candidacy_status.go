// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/olhoneles/politicos-go/db"
	"github.com/olhoneles/politicos-go/politicos"
	log "github.com/sirupsen/logrus"
)

// getAllCandidaciesStatus godoc
// @Summary Get all candidacies status
// @Description get all candidacies status data
// @Accept json
// @Produce json
// @Failure 200 {array} politicos.CandidacyStatus
// @Failure 404 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /candidacies-status [get]
func (s *server) getAllCandidaciesStatus(c echo.Context) error {
	log.Debug("[API] Retrieving all candidacies status")

	opts, err := db.NewGetAllOptionsBuilder(c.QueryParams())
	if err != nil {
		errMsg := fmt.Sprintf("Error on convert page value to int: %v", err)
		return logAndReturnError(c, errMsg)
	}

	result, err := s.db.GetAll(&politicos.CandidacyStatus{}, opts)
	if err != nil {
		errMsg := fmt.Sprintf("Error on retrieve candidacies status: %v", err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}
