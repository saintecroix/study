package model

type Genre struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Synonym     string `json:"synonym"`
	Description string `json:"description"`
}
