package harrypotter

import (
	_ "embed"

	"golang.org/x/exp/maps"
)

//go:embed characters.json
var database string

type Character struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Blood    string `json:"blood"`
	Species  string `json:"species"`
	Patronus string `json:"patronus"`
	Born     string `json:"born"`
	Quote    string `json:"quote"`
	ImgURL   string `json:"imgUrl"`
}

func Mtos(cm map[int]Character) []Character {
	return maps.Values(cm)
}

func Stom(cs []Character) map[int]Character {
	cm := map[int]Character{}

	for _, c := range cs {
		cm[c.ID] = c
	}

	return cm
}
