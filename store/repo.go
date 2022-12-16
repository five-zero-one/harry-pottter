package store

import (
	"encoding/json"
	"fmt"
	harrypotter "harry-potter"
	"harry-potter/pkg/mapconv"
	"harry-potter/store/option"
	"strings"

	_ "embed"
)

//go:embed characters.json
var database string

type Repo interface {
	Filter(f option.FilterOption) ([]harrypotter.Character, error)
	Search(key any) (*harrypotter.Character, error)
}

type repo struct {
	db map[int]harrypotter.Character
}

func New() Repo {
	var cs []harrypotter.Character
	if err := json.NewDecoder(strings.NewReader(database)).Decode(&cs); err != nil {
		panic(err)
	}

	r := &repo{db: mapconv.Stom(cs)}
	return r
}

func (r *repo) Filter(f option.FilterOption) ([]harrypotter.Character, error) {
	cs := make([]harrypotter.Character, 0)
	for _, c := range r.db {
		if f.WhereBlood(c.Blood) && f.WhereMonth(c.Born) && f.WhereName(c.Name) {
			cs = append(cs, c)
		}
	}

	return cs, nil
}

func (r *repo) Search(key any) (*harrypotter.Character, error) {
	switch key := key.(type) {
	case int:
		if c, ok := r.db[key]; ok {
			return &c, nil
		} else {
			return nil, fmt.Errorf("repo: character not found")
		}
	case string:
		for _, c := range r.db {
			// if c.Blood =
			_ = c
		}
	default:
		// panic or something
	}

	return nil, nil
}
