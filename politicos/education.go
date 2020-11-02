// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package politicos

type Education struct {
	Name  string `json:"name" bson:"name"`
	Slug  string `json:"slug" bson:"slug"`
	TseID string `json:"tseId" bson:"tseId"`
}

func (c Education) GetCollectionName() string {
	return "educations"
}

func (c Education) Cast() Queryable {
	return &c
}
