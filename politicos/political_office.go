// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package politicos

type PoliticalOffice struct {
	Description string `json:"description" bson:"description"`
	Name        string `json:"name" bson:"name"`
	Slug        string `json:"slug" bson:"slug"`
	Term        int    `json:"term" bson:"term"`
	TseID       string `json:"tseId" bson:"tseId"`
	Wikipedia   string `json:"wikipedia" bson:"wikipedia"`
}

func (c PoliticalOffice) GetCollectionName() string {
	return "political-offices"
}

func (c PoliticalOffice) Cast() Queryable {
	return &c
}
