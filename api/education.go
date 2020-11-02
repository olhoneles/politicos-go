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

// getAllEducations godoc
// @Summary Get all educations
// @Description get all educations data
// @Accept json
// @Produce json
// @Failure 200 {array} politicos.Education
// @Failure 404 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /educations [get]
func (s *server) getAllEducations(c echo.Context) error {
	log.Debug("[API] Retrieving all educations")

	opts, err := db.NewGetAllOptionsBuilder(c.QueryParams())
	if err != nil {
		errMsg := fmt.Sprintf("Error on convert page value to int: %v", err)
		return logAndReturnError(c, errMsg)
	}

	result, err := s.db.GetAll(&politicos.Education{}, opts)
	if err != nil {
		errMsg := fmt.Sprintf("Error on retrieve educations: %v", err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}
