// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package politicos

type Candidature interface {
	GetYear() int
	New(record []string)
}
