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
	page, err := strconv.Atoi(c.Get("page"))
	if err != nil {
		return nil, err
	}
	opts := &GetAllOptions{Page: page}
	return opts, nil
}
