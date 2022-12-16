package harrypotter

import (
	"strings"
)

type FilterOption struct {
	name  string
	month string
	blood string
}

func (f FilterOption) whereName(s string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(f.name))
}

func (f FilterOption) whereMonth(s string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(f.month))
}

func (f FilterOption) whereBlood(s string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(f.blood))
}

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

func NewFilter(m map[string][]string) FilterOption {
	fm := filter(m)

	opt := FilterOption{
		name:  fm.get("name"),
		month: fm.get("birth_month"),
		blood: fm.get("blood_type"),
	}
	return opt
}
