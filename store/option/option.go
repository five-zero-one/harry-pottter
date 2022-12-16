package option

import "strings"

type FilterOption struct {
	Name  string
	Month string
	Blood string
}

func (f FilterOption) WhereName(s string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(f.Name))
}

func (f FilterOption) WhereMonth(s string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(f.Month))
}

func (f FilterOption) WhereBlood(s string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(f.Blood))
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
		Name:  fm.get("name"),
		Month: fm.get("birth_month"),
		Blood: fm.get("blood_type"),
	}
	return opt
}
