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
	var page int
	pageStr := c.Get("page")
	if pageStr == "" {
		page = 1
	} else {
		var err error
		page, err = strconv.Atoi(pageStr)
		return nil, err
	}
	opts := &GetAllOptions{Page: page}
	return opts, nil
}
