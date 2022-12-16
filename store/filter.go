package store

import (
	"harry-potter/store/option"
)

func getValue(v map[string][]string, key string) string {
	if v == nil {
		return ""
	}
	vs := v[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

type filter map[string][]string

func (f filter) get(s string) string {
	return getValue(f, s)
}

func NewFilter(m map[string][]string) option.FilterOption {
	fm := filter(m)

	opt := option.FilterOption{
		Name:  fm.get("name"),
		Month: fm.get("birth_month"),
		Blood: fm.get("blood_type"),
	}
	return opt
}
