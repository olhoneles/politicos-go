// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package politicos

type PoliticalParty struct {
	FoundedDate string `json:"foundedDate" bson:"foundedDate"`
	Logo        string `json:"logo" bson:"logo"`
	Name        string `json:"name" bson:"name"`
	Siglum      string `json:"siglum" bson:"siglum"`
	TseNumber   string `json:"tseNumber" bson:"tseNumber"`
	Website     string `json:"website" bson:"website"`
	Wikipedia   string `json:"wikipedia" bson:"wikipedia"`
}

func (c PoliticalParty) GetCollectionName() string {
	return "political-parties"
}

func (c PoliticalParty) Cast() Queryable {
	return &c
}
