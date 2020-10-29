// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package politicos

type CandidacyStatus struct {
	TseID string `json:"tseId" bson:"tseId"`
	Name  string `json:"name" bson:"name"`
	Slug  string `json:"slug" bson:"slug"`
}

func (c CandidacyStatus) GetCollectionName() string {
	return "candidacies-status"
}

func (c CandidacyStatus) Cast() Queryable {
	return &c
}
