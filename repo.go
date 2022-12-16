package harrypotter

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Repo struct {
	db map[int]Character
}

func newRepo() *Repo {
	var cs []Character
	if err := json.NewDecoder(strings.NewReader(database)).Decode(&cs); err != nil {
		panic(err)
	}

	r := &Repo{db: Stom(cs)}
	return r
}

func (r *Repo) Filter(f FilterOption) ([]Character, error) {
	cs := make([]Character, 0)
	for _, c := range r.db {
		if f.whereBlood(c.Blood) && f.whereMonth(c.Born) && f.whereName(c.Name) {
			cs = append(cs, c)
		}
	}

	return cs, nil
}

func (r *Repo) Search(key any) (*Character, error) {
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
