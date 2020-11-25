// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package politicos

type Queryable interface {
	Cast() Queryable
	GetCollectionName() string
	GetID() string
	SetSlug(slug string)
}
