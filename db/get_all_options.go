// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package db

import (
	"net/url"
	"strconv"
)

type GetAllOptions struct {
	Page int
}

func NewGetAllOptionsBuilder(c url.Values) (*GetAllOptions, error) {
	pageStr := c.Get("page")
	if pageStr == "" {
		return &GetAllOptions{Page: 1}, nil
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return &GetAllOptions{Page: 1}, err
	}
	return &GetAllOptions{Page: page}, nil
}
