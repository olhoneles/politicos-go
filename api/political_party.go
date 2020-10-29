// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/olhoneles/politicos-go/politicos"
	log "github.com/sirupsen/logrus"
)

// getAllPoliticalParties godoc
// @Summary Get all political parties
// @Description get all political parties data
// @Accept json
// @Produce json
// @Failure 200 {array} politicos.PoliticalParty
// @Failure 404 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /political-parties [get]
func (s *server) getAllPoliticalParties(c echo.Context) error {
	log.Debug("[API] Retrieving all political parties")

	// FIXME
	var page int
	pageStr := c.QueryParam("page")
	if pageStr == "" {
		page = 1
	} else {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			errMsg := fmt.Sprintf("Error on convert page value to int: %v", err)
			return logAndReturnError(c, errMsg)
		}
	}

	result, err := s.db.GetAll(&politicos.PoliticalParty{}, page)
	if err != nil {
		errMsg := fmt.Sprintf("Error on retrieve candidatures: %v", err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}
