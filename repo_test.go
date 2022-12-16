package harrypotter

import (
	"testing"

	"github.com/hyphengolang/prelude/testing/is"
)

func TestRepo(t *testing.T) {
	is := is.New(t)

	r := newRepo()

	t.Run("return all members of the list", func(t *testing.T) {
		var q FilterOption
		cs, err := r.Filter(q)
		is.NoErr(err)         // filter characters list
		is.Equal(len(cs), 25) // 25 characters in the list
	})

	t.Run("return all half-blood characters", func(t *testing.T) {
		var q = FilterOption{
			blood: "half-blood",
		}
		cs, err := r.Filter(q)
		is.NoErr(err)         // filter by blood type
		is.Equal(len(cs), 11) // X characters in the list
	})

	t.Run("return all characters that born in February", func(t *testing.T) {
		var q = FilterOption{
			month: "february",
		}
		cs, err := r.Filter(q)
		is.NoErr(err)        // filter by blood type
		is.Equal(len(cs), 2) // X characters in the list
	})

	t.Run("return all characters that are `half-blood` and born in `September`", func(t *testing.T) {
		var q = FilterOption{
			blood: "half-blood",
			month: "Sep",
		}
		cs, err := r.Filter(q)
		is.NoErr(err)        // filter by blood type
		is.Equal(len(cs), 4) // X characters in the list
	})
}
