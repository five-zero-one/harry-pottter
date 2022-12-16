package mapconv

import (
	harrypotter "harry-potter"

	"golang.org/x/exp/maps"
)

func Mtos(cm map[int]harrypotter.Character) []harrypotter.Character {
	return maps.Values(cm)
}

func Stom(cs []harrypotter.Character) map[int]harrypotter.Character {
	cm := map[int]harrypotter.Character{}

	for _, c := range cs {
		cm[c.ID] = c
	}

	return cm
}
